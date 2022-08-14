package repository

import (
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
		return result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Printf("Master material %s have no weight_pack \n", ID)
		// return errors.New(fmt.Sprintf("Master material %s not found", ID))
		return nil
	}

	return nil
}
