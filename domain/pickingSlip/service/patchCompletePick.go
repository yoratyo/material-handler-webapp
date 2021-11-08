package service

import (
	"github.com/gin-gonic/gin"
	pickingSlipDTO "github.com/yoratyo/material-handler-webapp/model/dto/pickingSlip"
	helper "github.com/yoratyo/material-handler-webapp/shared"
)

func (s *service) PatchCompletePick(ctx *gin.Context, ID uint64, data pickingSlipDTO.PatchCompletePickRequestDTO) error {
	currDate, currTime := helper.GetCurrentTime()

	var updates = map[string]interface{}{
		"actual_bag":       data.ActualBag,
		"is_complete_pick": true,
		"date_pick":        currDate,
		"time_pick":        currTime,
		"operator_pick":    data.OperatorPick,
	}

	return s.repository.PatchByID(ctx, ID, updates)
}
