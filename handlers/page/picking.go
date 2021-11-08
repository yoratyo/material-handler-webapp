package page

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	pickingSlipDTO "github.com/yoratyo/material-handler-webapp/model/dto/pickingSlip"
	userDTO "github.com/yoratyo/material-handler-webapp/model/dto/userLogin"
	"github.com/yoratyo/material-handler-webapp/shared"
	"net/http"
)

func (h handler) Picking(c *gin.Context) {
	session := sessions.Default(c)
	userSession := session.Get(shared.USERKEY)
	userID, ok := userSession.(uint64)
	if !ok {
		c.JSON(http.StatusNotFound, userDTO.LoginResponseDTO{
			Message: "Failed get session",
		})
		return
	}

	var (
		err       error
		activeOKP string = "0"
	)

	response := gin.H{
		"isPickingMenu": true,
	}

	// auth me
	user, err := h.domain.User.GetByID(c, userID)
	if err != nil {
		response["errMessage"] = "Failed get user data"
		return
	}

	// get form request param
	var params pickingSlipDTO.GetPendingPickingRequestDTO
	err = c.Bind(&params)
	if err != nil {
		response["errMessage"] = "Failed to get param"
		return
	}

	if params.BatchNo != nil {
		activeOKP = *params.BatchNo
		if *params.BatchNo == "0" {
			params.BatchNo = nil
			activeOKP = "0"
		}
	}

	if params.SupplierLotNo != nil {
		if *params.SupplierLotNo == "" {
			params.SupplierLotNo = nil
		}
	}

	// get list pending OKP
	listOKP, err := h.domain.PickingSlip.GetDistinctPendingOKP(c)
	if err != nil {
		response["errMessage"] = "Failed get list pending data"
		return
	}

	// get list pending picking slip
	count, items, err := h.domain.PickingSlip.GetPending(c, params)
	if err != nil {
		response["errMessage"] = "Failed get list pending data"
		return
	}

	response["username"] = user.Username
	response["items"] = items
	response["okp"] = activeOKP
	response["listBatchNo"] = listOKP
	response["totalPicking"] = count

	errMessage, err := c.Cookie("error")
	if err == nil {
		response["errMessage"] = errMessage
	}

	successMessage, err := c.Cookie("success")
	if err == nil {
		response["successMessage"] = successMessage
	}

	fmt.Println(response)

	c.HTML(http.StatusOK, "picking.html", response)
}
