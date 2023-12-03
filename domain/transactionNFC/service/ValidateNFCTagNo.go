package service

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func (s *service) ValidateNFCTagNo(ctx *gin.Context, nfcTagNo string) error {
	// Validate nfc already registered
	trxItems, err := s.repository.GetByNFCTagNo(ctx, nfcTagNo)
	if err != nil {
		return err
	}
	if len(trxItems) > 0 {
		return errors.New("NFC already registered")
	}

	return nil
}
