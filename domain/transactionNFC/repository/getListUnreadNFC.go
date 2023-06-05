package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
)

func (r *repository) GetListUnreadNFC(ctx *gin.Context) ([]dao.NfcReader, error) {
	var items []dao.NfcReader

	query := r.resource.DB.
		WithContext(ctx).
		Model(&dao.NfcReader{}).
		Where("bitActive = ?", 0)

	if err := query.Find(&items).Error; err != nil {
		return []dao.NfcReader{}, err
	}

	return items, nil
}
