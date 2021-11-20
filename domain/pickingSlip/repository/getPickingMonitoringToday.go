package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetPickingMonitoringToday(ctx *gin.Context) (dao.PickingMonitoring, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return dao.PickingMonitoring{}, err
	}

	var res dao.PickingMonitoring

	query := orm.
		WithContext(ctx).
		Model(&dao.MasterPickingSlip{}).
		Select("sum(case when is_complete_pick = true then 1 else 0 end) AS ProgressCount", "count(*) AS TotalCount").
		Where("is_ready_for_pick = ?", true).
		Where("DATE(date_submit) = CURDATE()")

	if err := query.Find(&res).Error; err != nil {
		return dao.PickingMonitoring{}, err
	}

	return res, nil
}
