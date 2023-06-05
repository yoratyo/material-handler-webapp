package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
	transactionNFCDTO "github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
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

func (h handler) SchedulerGatewayCheckNFC(time time.Time) {
	c, _ := context.Background().(*gin.Context)
	list, err := h.domain.TransactionNFC.GetListUnreadNFC(c)
	if err != nil {
		log.Println(err)
	}

	tags := make([]transactionNFCDTO.DetailPatchGatewayCheck, len(list))
	tagStr := []string{}
	for i, n := range list {
		tags[i] = transactionNFCDTO.DetailPatchGatewayCheck{
			TagNFC: n.NFCTagNo,
		}
		tagStr = append(tagStr, n.NFCTagNo)
	}

	err = h.domain.TransactionNFC.BulkPatchGatewayCheck(c, transactionNFCDTO.PatchGatewayCheckRequestDTO{
		Items: tags,
	})
	if err != nil {
		log.Println(err)
	}

	err = h.domain.TransactionNFC.BulkPatchReadNFC(c, tagStr)
	if err != nil {
		log.Println(err)
	}
}
