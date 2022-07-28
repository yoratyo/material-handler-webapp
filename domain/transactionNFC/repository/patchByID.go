package repository

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) PatchByID(ctx *gin.Context, ID uint64, data map[string]interface{}) error {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return err
	}

	result := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Where("id_transact = ?", ID).
		Updates(data)
	if result.Error != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("transaction nfc %d not found", ID)
	}

	return nil
}
