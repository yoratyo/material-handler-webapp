package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetDefaultMonitoringNFC(ctx *gin.Context) (dao.MonitoringNFC, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return dao.MonitoringNFC{}, err
	}

	var item dao.MonitoringNFC

	query := orm.
		WithContext(ctx).
		Model(&dao.MonitoringNFC{}).
		Where("is_default = ?", true)

	if err := query.Take(&item).Error; err != nil {
		return dao.MonitoringNFC{}, err
	}

	return item, nil
}
