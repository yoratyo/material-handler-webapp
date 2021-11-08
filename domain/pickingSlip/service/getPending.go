package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	pickingSlipDTO "github.com/yoratyo/material-handler-webapp/model/dto/pickingSlip"
)

func (s *service) GetPending(ctx *gin.Context, req pickingSlipDTO.GetPendingPickingRequestDTO) (uint64, []dao.MasterPickingSlip, error) {
	count, err := s.repository.GetCountPendingToPick(ctx)
	if err != nil {
		return 0, []dao.MasterPickingSlip{}, err
	}

	items, err := s.repository.GetPendingToPick(ctx, req)
	if err != nil {
		return 0, []dao.MasterPickingSlip{}, err
	}

	return count, items, nil
}
