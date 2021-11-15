package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	transactionNFCDTO "github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetRegisterNFC(ctx *gin.Context, req transactionNFCDTO.PatchRegisterNFCRequestDTO) (dao.TransactionNFC, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return dao.TransactionNFC{}, err
	}

	var item dao.TransactionNFC

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

	if err := query.Order("bag_no").Take(&item).Error; err != nil {
		return dao.TransactionNFC{}, err
	}

	return item, nil
}
