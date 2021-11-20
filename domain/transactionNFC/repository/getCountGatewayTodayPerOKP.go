package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetCountGatewayTodayPerOKP(ctx *gin.Context) ([]dao.CountGatewayPerOKP, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return []dao.CountGatewayPerOKP{}, err
	}

	var res []dao.CountGatewayPerOKP

	query := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Joins("JOIN master_picking_slip ON master_picking_slip.id = transaction_nfc.id_master_picking_slip").
		Select("sum(case when transaction_nfc.is_gateway_check = true then 1 else 0 end) AS TotalChecked",
			"count(*) AS TotalRegistered",
			"master_picking_slip.batch_no AS BatchNo",
			"master_picking_slip.formula_description AS FormulaDescription").
		Where("transaction_nfc.is_register = ?", true).
		Where("DATE(master_picking_slip.date_submit) = CURDATE()").
		Group("master_picking_slip.batch_no").
		Group("master_picking_slip.formula_description")

	if err := query.Find(&res).Error; err != nil {
		return []dao.CountGatewayPerOKP{}, err
	}

	return res, nil
}
