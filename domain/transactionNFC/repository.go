package transactionNFC

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
)

type Repository interface {
	BulkCreate(ctx *gin.Context, req []dao.TransactionNFC) ([]dao.TransactionNFC, error)
	GetByID(ctx *gin.Context, ID uint64) (dao.TransactionNFC, error)
	CheckExistByIDPickingSlip(ctx *gin.Context, IDPickingSlip uint64) (bool, error)
}
