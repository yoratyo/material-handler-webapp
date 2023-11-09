package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetDistinctItemOKP(ctx *gin.Context, batchNo string) ([]dao.ItemPickingSlip, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return []dao.ItemPickingSlip{}, err
	}

	var list []dao.ItemPickingSlip

	query := orm.
		WithContext(ctx).
		Model(&dao.MasterPickingSlip{}).
		Where("batch_no = ?", batchNo)

	if err := query.Distinct("item_code", "item_description").Find(&list).Error; err != nil {
		return []dao.ItemPickingSlip{}, err
	}

	return list, nil
}
