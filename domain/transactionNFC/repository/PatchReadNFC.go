package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
)

func (r *repository) PatchReadNFR(ctx *gin.Context, tags []string, data map[string]interface{}) (int64, error) {
	result := r.resource.DB.
		WithContext(ctx).
		Model(&dao.NfcReader{}).
		Where("txtTid IN ?", tags).
		Where("bitActive = ?", 0).
		Updates(data)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
