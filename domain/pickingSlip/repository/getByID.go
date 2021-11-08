package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetByID(ctx *gin.Context, ID uint64) (dao.MasterPickingSlip, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return dao.MasterPickingSlip{}, err
	}

	var item dao.MasterPickingSlip

	query := orm.
		WithContext(ctx).
		Model(&dao.MasterPickingSlip{}).
		Where("id = ?", ID)

	if err := query.Take(&item).Error; err != nil {
		return dao.MasterPickingSlip{}, err
	}

	return item, nil
}
