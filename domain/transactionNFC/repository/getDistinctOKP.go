package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetDistinctOKP(ctx *gin.Context, status string) ([]string, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return []string{}, err
	}

	var list []string

	query := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{})

	switch status {
	case shared.NFC_STATUS_PENDING_REGISTER:
		query.Where("is_register = ?", false)
	case shared.NFC_STATUS_PENDING_GATEWAY_CHECK:
		query.
			Where("is_register = ?", true).
			Where("is_gateway_check = ?", false)
	default:
		//
	}

	if err := query.Distinct().Pluck("batch_no", &list).Error; err != nil {
		return []string{}, err
	}

	return list, nil
}
