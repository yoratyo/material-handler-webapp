package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	transactionNFCDTO "github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetListRegisterNFC(ctx *gin.Context, req transactionNFCDTO.GetRegisterNFCRequestDTO) ([]dao.TransactionNFC, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return []dao.TransactionNFC{}, err
	}

	var items []dao.TransactionNFC

	query := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Where("is_register = ?", false)

	if req.BatchNo != nil {
		query = query.Where("batch_no = ?", *req.BatchNo)
	}

	if req.SupplierLotNo != nil {
		query = query.Where("supplier_lot_no LIKE ?", *req.SupplierLotNo+"%")
	}

	if err := query.Find(&items).Error; err != nil {
		return []dao.TransactionNFC{}, err
	}

	return items, nil
}
