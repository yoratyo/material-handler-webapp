package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) PatchByID(ctx *gin.Context, ID uint64, data map[string]interface{}) error {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return err
	}

	if err := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Where("id = ?", ID).
		Updates(data).
		Error; err != nil {
		return err
	}

	return nil
}
