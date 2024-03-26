package pageAction

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	pickingSlipDTO "github.com/yoratyo/material-handler-webapp/model/dto/pickingSlip"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (h handler) SubmitPickingSlip(c *gin.Context) {
	var (
		err    error
		path   = "/page/picking"
		params = url.Values{}
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
		addError(c, err.Error(), path, params)
		return
	}

	if req.BatchNo != nil {
		if *req.BatchNo != "0" {
			params.Add("batchNo", *req.BatchNo)
		}
	}

	if req.ItemCode != nil {
		if *req.ItemCode != "0" {
			params.Add("itemCode", *req.ItemCode)
		}
	}

	req.OperatorPick = user.Username

	// start transaction
	tx := h.resource.DB.Begin()
	c.Set(shared.MySQLTransaction, tx)

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			addError(c, "Panic error", path, params)
			return
		}
	}()

	// get picking slip
	pickingSlip, err := h.domain.PickingSlip.GetByID(c, req.ID)
	if err != nil {
		tx.Rollback()
		addError(c, err.Error(), path, params)

		return
	}

	// validate picking slip
	if !pickingSlip.IsReadyForPick || pickingSlip.IsCancel || pickingSlip.IsCompletePick {
		tx.Rollback()
		addError(c, "Data picking slip not valid to pick", path, params)

		return
	}
	if pickingSlip.WeightPack == nil {
		tx.Rollback()
		addError(c, "Weight pack is empty", path, params)

		return
	}

	// Patch complete pickingSlip
	err = h.domain.PickingSlip.PatchCompletePick(c, req.ID, req)
	if err != nil {
		tx.Rollback()
		addError(c, err.Error(), path, params)

		return
	}

	if req.ActualBag != 0 {
		// Bulk create transaction NFC
		_, err = h.domain.TransactionNFC.BulkCreate(c, pickingSlip, req.ActualBag)
		if err != nil {
			tx.Rollback()
			addError(c, err.Error(), path, params)

			return
		}
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		addError(c, err.Error(), path, params)

		return
	}

	shared.SetSuccessCookie(c, "Success to submit picking slip")
	location := url.URL{Path: path}
	if len(params) != 0 {
		location.RawQuery = params.Encode()
	}
	c.Redirect(http.StatusFound, location.RequestURI())
	c.Abort()
}
