package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetByID(ctx *gin.Context, ID uint64) (dao.TransactionNFC, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return dao.TransactionNFC{}, err
	}

	var item dao.TransactionNFC

	query := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Where("id = ?", ID)

	if err := query.Take(&item).Error; err != nil {
		return dao.TransactionNFC{}, err
	}

	return item, nil
}
