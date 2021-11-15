package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetBySupplierLotNoGroupByBatchNo(ctx *gin.Context, supplierLotNo string) ([]dao.TransactionNFC, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return []dao.TransactionNFC{}, err
	}

	var items []dao.TransactionNFC

	query := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Select("batch_no").
		Where("is_register = ?", false).
		Where("supplier_lot_no LIKE ?", supplierLotNo+"%").
		Group("batch_no")

	if err := query.Find(&items).Error; err != nil {
		return []dao.TransactionNFC{}, err
	}

	return items, nil
}
