package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetCountGatewayToday(ctx *gin.Context) (dao.CountGatewayTransactionNFC, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return dao.CountGatewayTransactionNFC{}, err
	}

	var res dao.CountGatewayTransactionNFC

	query := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Joins("JOIN master_picking_slip ON master_picking_slip.id = transaction_nfc.id_master_picking_slip").
		Select("sum(case when transaction_nfc.is_gateway_check = true then 1 else 0 end) AS TotalChecked", "count(*) AS TotalRegistered").
		Where("transaction_nfc.is_register = ?", true).
		Where("DATE(master_picking_slip.date_submit) = CURDATE()")

	if err := query.Find(&res).Error; err != nil {
		return dao.CountGatewayTransactionNFC{}, err
	}

	return res, nil
}
