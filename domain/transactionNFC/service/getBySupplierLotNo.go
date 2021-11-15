package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
)

func (s *service) GetBySupplierLotNo(ctx *gin.Context, supplierLotNo string) ([]dao.TransactionNFC, error) {
	return s.repository.GetBySupplierLotNoGroupByBatchNo(ctx, supplierLotNo)
}
