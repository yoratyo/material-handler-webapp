package page

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	transactionNFCDTO "github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
	userDTO "github.com/yoratyo/material-handler-webapp/model/dto/userLogin"
	"github.com/yoratyo/material-handler-webapp/shared"
	"net/http"
)

func (h handler) NFC(c *gin.Context) {
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
		activeOKP string = "0"
		err error
	)

	response := gin.H{
		"isNFCMenu": true,
	}

	// auth me
	user, err := h.domain.User.GetByID(c, userID)
	if err != nil {
		response["errMessage"] = "Failed get user data"
		return
	}

	// get form request param
	var params transactionNFCDTO.GetRegisterNFCRequestDTO
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
	listOKP, err := h.domain.TransactionNFC.GetPendingOKP(c, shared.NFC_STATUS_PENDING_REGISTER)
	if err != nil {
		response["errMessage"] = "Failed get list pending data"
		return
	}

	// get list to be register NFC
	count, items, err := h.domain.TransactionNFC.GetListRegisterNFC(c, params)
	if err != nil {
		response["errMessage"] = "Failed get list pending data"
		return
	}

	response["username"] = user.Username
	response["items"] = items
	response["barcode"] = params.SupplierLotNo
	response["totalRegister"] = count
	response["listBatchNo"] = listOKP
	response["okp"] = activeOKP

	errMessage, err := c.Cookie("error")
	if err == nil {
		response["errMessage"] = errMessage
	}

	successMessage, err := c.Cookie("success")
	if err == nil {
		response["successMessage"] = successMessage
	}

	c.HTML(http.StatusOK, "nfc.html", response)
}
