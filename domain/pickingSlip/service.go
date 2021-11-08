package pickingSlip

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	pickingSlipDTO "github.com/yoratyo/material-handler-webapp/model/dto/pickingSlip"
)

type Service interface {
	GetPending(ctx *gin.Context, req pickingSlipDTO.GetPendingPickingRequestDTO) (uint64, []dao.MasterPickingSlip, error)
	GetDistinctPendingOKP(ctx *gin.Context) ([]string, error)
	PatchCompletePick(ctx *gin.Context, ID uint64, data pickingSlipDTO.PatchCompletePickRequestDTO) error
	GetByID(ctx *gin.Context, ID uint64) (dao.MasterPickingSlip, error)
}
