package service

import (
	"github.com/gin-gonic/gin"
	transactionNFCDTO "github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
	helper "github.com/yoratyo/material-handler-webapp/shared"
)

func (s *service) PatchCompleteRegister(ctx *gin.Context, ID uint64, data transactionNFCDTO.PatchRegisterNFCRequestDTO) error {
	currDate, currTime := helper.GetCurrentTime()

	var updates = map[string]interface{}{
		"nfc_tag_no":        data.DataNFC,
		"is_register":       true,
		"date_register":     currDate,
		"time_register":     currTime,
		"operator_register": data.OperatorRegister,
	}

	return s.repository.PatchByID(ctx, ID, updates)
}
