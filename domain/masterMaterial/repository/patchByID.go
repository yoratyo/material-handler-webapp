package repository

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) PatchByID(ctx *gin.Context, ID string, data map[string]interface{}) error {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return err
	}

	result := orm.
		WithContext(ctx).
		Model(&dao.MasterMaterial{}).
		Where("item_code = ?", ID).
		Updates(data)
	if result.Error != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("Master material %s not found", ID))
	}

	return nil
}
