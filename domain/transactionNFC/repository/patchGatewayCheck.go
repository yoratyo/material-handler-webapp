package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) PatchGatewayCheck(ctx *gin.Context, Ids []string, data map[string]interface{}) (int64, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return 0, err
	}

	result := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Where("nfc_tag_no IN ?", Ids).
		Where("is_register = ?", true).
		Where("is_gateway_check = ?", false).
		Updates(data)
	if result.Error != nil {
		return 0, err
	}

	return result.RowsAffected, nil
}
