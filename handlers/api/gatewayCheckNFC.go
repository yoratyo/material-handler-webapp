package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
	"net/http"
)

func (h handler) GatewayCheckNFC(c *gin.Context) {
	// validate payload
	var req transactionNFC.PatchGatewayCheckRequestDTO
	errs := c.ShouldBindJSON(&req)
	if errs != nil {
		c.JSON(http.StatusBadRequest, transactionNFC.PatchGatewayCheckResponseDTO{
			Message: "payload failed",
		})
		return
	}

	go h.domain.TransactionNFC.BulkPatchGatewayCheck(c, req)

	c.JSON(http.StatusOK, transactionNFC.PatchGatewayCheckResponseDTO{
		Message: "success",
	})
}
