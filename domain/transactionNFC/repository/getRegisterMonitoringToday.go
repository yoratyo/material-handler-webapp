package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetRegisterMonitoringToday(ctx *gin.Context) (dao.RegisterNFCMonitoring, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return dao.RegisterNFCMonitoring{}, err
	}

	var res dao.RegisterNFCMonitoring

	query := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Joins("JOIN master_picking_slip ON master_picking_slip.id_transact = transaction_nfc.id_master_picking_slip").
		Select("sum(case when transaction_nfc.is_register = true then 1 else 0 end) AS ProgressCount", "count(*) AS TotalCount").
		Where("DATE(master_picking_slip.date_submit) = CURDATE()")

	if err := query.Find(&res).Error; err != nil {
		return dao.RegisterNFCMonitoring{}, err
	}

	return res, nil
}
