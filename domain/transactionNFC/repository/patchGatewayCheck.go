package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
)

func (r *repository) PatchGatewayCheck(ctx *gin.Context, Ids []string, data map[string]interface{}) (int64, error) {
	result := r.resource.DB.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Where("nfc_tag_no IN ?", Ids).
		Where("is_register = ?", true).
		Where("is_gateway_check = ?", false).
		Updates(data)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
