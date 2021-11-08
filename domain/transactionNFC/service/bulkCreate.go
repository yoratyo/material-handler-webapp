package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
)

func (s *service) BulkCreate(ctx *gin.Context, pickingSlip dao.MasterPickingSlip, actualBag uint64) ([]dao.TransactionNFC, error) {
	exist, err := s.repository.CheckExistByIDPickingSlip(ctx, pickingSlip.ID)
	if !exist && err == nil {
		var items []dao.TransactionNFC

		for no := 1; no <= int(actualBag); no++ {
			trx := dao.TransactionNFC{
				IDMasterPickingSlip: pickingSlip.ID,
				PickSlipNo:          pickingSlip.PickSlipNo,
				BatchNo:             pickingSlip.BatchNo,
				ItemCode:            pickingSlip.ItemCode,
				ItemDescription:     pickingSlip.ItemDescription,
				StockLocator:        pickingSlip.StockLocator,
				LotNo:               pickingSlip.LotNo,
				SupplierLotNo:       pickingSlip.SupplierLotNo,
				BagNo:               uint64(no),
				TotalBag:            actualBag,
				WeightPack:          *pickingSlip.WeightPack,
				RemainingWeight:     *pickingSlip.WeightPack,
				IsRegister:          false,
				IsGatewayCheck:      false,
				IsCompleteBag:       false,
				IsSplitMaterial:     false,
			}
			items = append(items, trx)
		}

		return s.repository.BulkCreate(ctx, items)
	}

	return []dao.TransactionNFC{}, errors.New("Not valid to insert to transaction NFC")
}
