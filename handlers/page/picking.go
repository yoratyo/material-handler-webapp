package page

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	pickingSlipDTO "github.com/yoratyo/material-handler-webapp/model/dto/pickingSlip"
	userDTO "github.com/yoratyo/material-handler-webapp/model/dto/userLogin"
	"github.com/yoratyo/material-handler-webapp/shared"
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
		err                                error
		activeOKP                          string = "0"
		milisNotifSuccess, milisNotifError int64  = 0, 0
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

	// Validate user access
	if user.UserLevel == shared.ROLE_LEVEL_OPERATOR {
		if !user.PickingMenu {
			shared.SetErrorCookie(c, "You don't have access to Picking Material")
			c.Redirect(http.StatusFound, "/page/home")
			c.Abort()
		}
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
	response["role"] = user.UserLevel
	response["items"] = items
	response["okp"] = activeOKP
	response["listBatchNo"] = listOKP
	response["totalPicking"] = count

	errMessage, err := c.Cookie("error")
	if err == nil {
		timeError, err := c.Cookie("errorTime")
		if err == nil {
			milisNotifError, err = strconv.ParseInt(timeError, 10, 64)
			if err != nil {
				fmt.Println("Failed convert Error notif time")
			}
			shared.RemoveErrorCookie(c)
		}
	}

	successMessage, err := c.Cookie("success")
	if err == nil {
		timeSuccess, err := c.Cookie("successTime")
		if err == nil {
			milisNotifSuccess, err = strconv.ParseInt(timeSuccess, 10, 64)
			if err != nil {
				fmt.Println("Failed convert Success notif time")
			}
			shared.RemoveSuccessCookie(c)
		}
	}

	if milisNotifSuccess != 0 && milisNotifError != 0 {
		if milisNotifSuccess > milisNotifError {
			response["successMessage"] = successMessage
		} else {
			response["errMessage"] = errMessage
		}
	} else {
		if milisNotifSuccess != 0 {
			response["successMessage"] = successMessage
		}
		if milisNotifError != 0 {
			response["errMessage"] = errMessage
		}
	}

	c.HTML(http.StatusOK, "picking.html", response)
}
