package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetCountPendingToRegister(ctx *gin.Context) (uint64, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return 0, err
	}

	var total int64

	query := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Where("is_register = ?", false)

	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	return uint64(total), nil
}
