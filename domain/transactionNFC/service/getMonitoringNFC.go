package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"gorm.io/gorm"
)

func (s *service) GetMonitoringNFC(ctx *gin.Context) (dao.MonitoringNFC, error) {
	var (
		nfc dao.MonitoringNFC
		err error
	)
	nfc, err = s.repository.GetMonitoringNFCByIP(ctx, ctx.ClientIP())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			nfc, err = s.repository.GetDefaultMonitoringNFC(ctx)
			if err != nil {
				return dao.MonitoringNFC{}, err
			}
		} else {
			return dao.MonitoringNFC{}, err
		}
	}

	// Validate NFC not in reader
	if nfc.NfcData == nil {
		return dao.MonitoringNFC{}, errors.New("There is no NFC data on the reader")
	}

	// Validate nfc already registered
	err = s.ValidateNFCTagNo(ctx, *nfc.NfcData)
	if err != nil {
		return dao.MonitoringNFC{}, err
	}

	return nfc, nil
}
