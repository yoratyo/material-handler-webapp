package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	pickingSlipDTO "github.com/yoratyo/material-handler-webapp/model/dto/pickingSlip"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetPendingToPick(ctx *gin.Context, req pickingSlipDTO.GetPendingPickingRequestDTO) ([]dao.MasterPickingSlip, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return []dao.MasterPickingSlip{}, err
	}

	var items []dao.MasterPickingSlip

	query := orm.
		WithContext(ctx).
		Model(&dao.MasterPickingSlip{}).
		Where("is_ready_for_pick = ?", true).
		Where("is_complete_pick = ?", false).
		Where("is_cancel = ?", false)

	if req.BatchNo != nil {
		query = query.Where("batch_no = ?", *req.BatchNo)
	}

	if req.SupplierLotNo != nil {
		query = query.Where("supplier_lot_no LIKE ?", *req.SupplierLotNo+"%")
	}

	if err := query.Find(&items).Error; err != nil {
		return []dao.MasterPickingSlip{}, err
	}

	return items, nil
}
