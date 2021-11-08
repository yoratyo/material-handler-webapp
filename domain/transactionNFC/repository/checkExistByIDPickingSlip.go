package repository

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
	"gorm.io/gorm"
)

func (r *repository) CheckExistByIDPickingSlip(ctx *gin.Context, IDPickingSlip uint64) (bool, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return false, err
	}

	var item dao.TransactionNFC

	query := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Where("id_master_picking_slip = ?", IDPickingSlip)

	if err := query.First(&item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
