package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
)

func (s *service) GetListUnreadNFC(ctx *gin.Context) ([]dao.NfcReader, error) {
	items, err := s.repository.GetListUnreadNFC(ctx)
	if err != nil {
		return []dao.NfcReader{}, err
	}

	return items, nil
}
