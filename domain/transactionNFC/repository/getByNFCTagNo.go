package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (r *repository) GetByNFCTagNo(ctx *gin.Context, nfcTagNo string) ([]dao.TransactionNFC, error) {
	orm, err := shared.GetORMTransaction(ctx, r.resource)
	if err != nil {
		return []dao.TransactionNFC{}, err
	}

	var items []dao.TransactionNFC

	query := orm.
		WithContext(ctx).
		Model(&dao.TransactionNFC{}).
		Where("nfc_tag_no = ?", nfcTagNo)

	if err := query.Find(&items).Error; err != nil {
		return []dao.TransactionNFC{}, err
	}

	return items, nil
}
