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
		Model(&dao.MasterPickingSlip{}).
		Where("id_transact = ?", ID).
		Updates(data)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Printf("no updated rows for master picking slip %d \n", ID)
		return nil
	}

	return nil
}
