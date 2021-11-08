package pickingSlip

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	pickingSlipDTO "github.com/yoratyo/material-handler-webapp/model/dto/pickingSlip"
)

type Repository interface {
	GetPendingToPick(ctx *gin.Context, req pickingSlipDTO.GetPendingPickingRequestDTO) ([]dao.MasterPickingSlip, error)
	GetCountPendingToPick(ctx *gin.Context) (uint64, error)
	GetDistinctPendingOKP(ctx *gin.Context) ([]string, error)
	PatchByID(ctx *gin.Context, ID uint64, data map[string]interface{}) error
	GetByID(ctx *gin.Context, ID uint64) (dao.MasterPickingSlip, error)
}
