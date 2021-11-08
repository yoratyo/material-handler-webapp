package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) BulkCreate(ctx *gin.Context, req []dao.TransactionNFC) ([]dao.TransactionNFC, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return []dao.TransactionNFC{}, err
	}

	if err := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Create(&req).
		Error; err != nil {
		return []dao.TransactionNFC{}, err
	}

	return req, nil
}
