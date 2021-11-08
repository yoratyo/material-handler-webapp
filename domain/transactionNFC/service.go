package transactionNFC

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
)

type Service interface {
	BulkCreate(ctx *gin.Context, pickingSlip dao.MasterPickingSlip, actualBag uint64) ([]dao.TransactionNFC, error)
	GetByID(ctx *gin.Context, ID uint64) (dao.TransactionNFC, error)
}
