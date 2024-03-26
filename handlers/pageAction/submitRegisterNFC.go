package pageAction

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	pickingSlipDTO "github.com/yoratyo/material-handler-webapp/model/dto/pickingSlip"
	transactionNFCDTO "github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (h handler) SubmitRegisterNFC(c *gin.Context) {
	var (
		err    error
		path   = "/page/nfc"
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
	var req transactionNFCDTO.PatchRegisterNFCRequestDTO
	err = c.ShouldBind(&req)
	if err != nil {
		addError(c, err.Error(), path, params)
		return
	}

	if req.OKPNo != nil {
		if *req.OKPNo != "0" {
			params.Add("batchNo", *req.OKPNo)
		}
	}

	if req.ItemCode != nil {
		if *req.ItemCode != "0" {
			params.Add("itemCode", *req.ItemCode)
		}
	}

	// Set OKP number to null if 0 (not selected)
	if req.BatchNo != nil {
		if *req.BatchNo == "0" {
			req.BatchNo = nil
		}
	}

	// Directly to page NFC, just filter by OKP number
	if req.BatchNo != nil && req.SupplierLotNo == nil {
		urlPath := fmt.Sprintf("%s?batchNo=%s", path, *req.BatchNo)
		c.Redirect(http.StatusFound, urlPath)
		c.Abort()

		return
	}

	// Check get by supplier lot number
	if req.BatchNo == nil && req.SupplierLotNo != nil {
		items, err := h.domain.TransactionNFC.GetBySupplierLotNo(c, *req.SupplierLotNo)
		if err != nil {
			addError(c, err.Error(), path, params)

			return
		}

		if len(items) != 1 {
			if len(items) > 1 {
				shared.SetSuccessCookie(c, fmt.Sprintf("There are %d OKP with this supplier lot, please choose", len(items)))
			}
			urlPath := fmt.Sprintf("%s?supplierLotNo=%s", path, *req.SupplierLotNo)
			c.Redirect(http.StatusFound, urlPath)
			c.Abort()

			return
		}
	}

	req.OperatorRegister = user.Username

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

	// filtering to get transaction NFC
	trxNFC, err := h.domain.TransactionNFC.GetRegisterNFC(c, req)
	if err != nil {
		tx.Rollback()
		addError(c, err.Error(), path, params)

		return
	}

	// validate transaction NFC
	if trxNFC.NFCTagNo != nil || trxNFC.DateRegister != nil || trxNFC.TimeRegister != nil || trxNFC.OperatorRegister != nil {
		tx.Rollback()
		addError(c, "Data transaction not valid to register", path, params)

		return
	}

	if req.DataNFC == "" {
		// Get NFC data with validation
		nfc, err := h.domain.TransactionNFC.GetMonitoringNFC(c)
		if err != nil {
			tx.Rollback()
			addError(c, err.Error(), path, params)

			return
		}

		req.DataNFC = *nfc.NfcData
	} else {
		if len(req.DataNFC) < 16 {
			tx.Rollback()
			addError(c, "NFC data less than 16 characters.", path, params)

			return
		}

		req.DataNFC = req.DataNFC[len(req.DataNFC)-16:]

		err = h.domain.TransactionNFC.ValidateNFCTagNo(c, req.DataNFC)
		if err != nil {
			tx.Rollback()
			addError(c, err.Error(), path, params)

			return
		}
	}

	// Patch complete register NFC
	err = h.domain.TransactionNFC.PatchCompleteRegister(c, trxNFC.ID, req)
	if err != nil {
		tx.Rollback()
		addError(c, err.Error(), path, params)

		return
	}

	// Update total qty on master material
	err = h.domain.MasterMaterial.PatchTotalQty(c, trxNFC.ItemCode)
	if err != nil {
		tx.Rollback()
		addError(c, err.Error(), path, params)

		return
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		addError(c, err.Error(), path, params)

		return
	}

	shared.SetSuccessCookie(c, "Success to register NFC")
	location := url.URL{Path: path}
	if len(params) != 0 {
		location.RawQuery = params.Encode()
	}
	c.Redirect(http.StatusFound, location.RequestURI())
	c.Abort()
}
