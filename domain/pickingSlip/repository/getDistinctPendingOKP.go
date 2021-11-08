package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetDistinctPendingOKP(ctx *gin.Context) ([]string, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return []string{}, err
	}

	var list []string

	query := orm.
		WithContext(ctx).
		Model(&dao.MasterPickingSlip{}).
		Where("is_ready_for_pick = ?", true).
		Where("is_complete_pick = ?", false).
		Where("is_cancel = ?", false)

	if err := query.Distinct().Pluck("batch_no", &list).Error; err != nil {
		return []string{}, err
	}

	return list, nil
}
