package pageAction

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	pickingSlipDTO "github.com/yoratyo/material-handler-webapp/model/dto/pickingSlip"
	"github.com/yoratyo/material-handler-webapp/shared"
	"net/http"
	"net/url"
)

func (h handler) SubmitPickingSlip(c *gin.Context) {
	var (
		err error
	)

	session := sessions.Default(c)
	userSession := session.Get(shared.USERKEY)
	userID, ok := userSession.(uint64)
	if !ok {
		c.JSON(http.StatusNotFound, pickingSlipDTO.PatchCompletePickResponseDTO{
			Message: "Failed get session",
		})
		return
	}

	// auth me
	user, err := h.domain.User.GetByID(c, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, pickingSlipDTO.PatchCompletePickResponseDTO{
			Message: "Failed get user data",
		})
		return
	}

	// validate form input
	var req pickingSlipDTO.PatchCompletePickRequestDTO
	err = c.ShouldBind(&req)
	if err != nil {
		addError(c, err.Error())
		return
	}

	req.OperatorPick = user.Username

	// start transaction
	tx := h.resource.DB.Begin()
	c.Set(shared.MySQLTransaction, tx)

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			addError(c, "Panic error")
			return
		}
	}()

	// get picking slip
	pickingSlip, err := h.domain.PickingSlip.GetByID(c, req.ID)
	if err != nil {
		tx.Rollback()
		addError(c, err.Error())

		return
	}

	// validate picking slip
	if !pickingSlip.IsReadyForPick || pickingSlip.IsCancel || pickingSlip.IsCompletePick {
		tx.Rollback()
		addError(c, "Data picking slip not valid to pick")

		return
	}
	if pickingSlip.WeightPack == nil {
		tx.Rollback()
		addError(c, "Weight pack is empty")

		return
	}

	// Patch complete pickingSlip
	err = h.domain.PickingSlip.PatchCompletePick(c, req.ID, req)
	if err != nil {
		tx.Rollback()
		addError(c, err.Error())

		return
	}

	// Bulk create transaction NFC
	_, err = h.domain.TransactionNFC.BulkCreate(c, pickingSlip, req.ActualBag)
	if err != nil {
		tx.Rollback()
		addError(c, err.Error())

		return
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		addError(c, err.Error())

		return
	}

	c.SetCookie("success", "Success to submit picking slip", 10, "/", c.Request.URL.Hostname(), false, true)
	location := url.URL{Path: "/page/picking"}
	c.Redirect(http.StatusFound, location.RequestURI())
	c.Abort()
}

func addError(c *gin.Context, msg string) {
	c.SetCookie("error", msg, 10, "/", c.Request.URL.Hostname(), false, true)
	location := url.URL{Path: "/page/picking"}
	c.Redirect(http.StatusFound, location.RequestURI())
}
