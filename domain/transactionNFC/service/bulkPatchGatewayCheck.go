package service

import (
	"github.com/gin-gonic/gin"
	transactionNFCDTO "github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
	helper "github.com/yoratyo/material-handler-webapp/shared"
)

func (s *service) BulkPatchGatewayCheck(ctx *gin.Context, request transactionNFCDTO.PatchGatewayCheckRequestDTO) error {
	currDate, currTime := helper.GetCurrentTime()

	var updates = map[string]interface{}{
		"is_gateway_check":   true,
		"date_gateway_check": currDate,
		"time_gateway_check": currTime,
	}

	_, err := s.repository.PatchGatewayCheck(ctx, request.ToUpdateGatewayCheck(), updates)
	if err != nil {
		return err
	}

	//log.Printf("%d Row affected\n", rowsAffected)

	return nil
}
