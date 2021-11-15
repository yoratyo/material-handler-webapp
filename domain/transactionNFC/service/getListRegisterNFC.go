package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	transactionNFCDTO "github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
)

func (s *service) GetListRegisterNFC(ctx *gin.Context, req transactionNFCDTO.GetRegisterNFCRequestDTO) (uint64, []dao.TransactionNFC, error) {
	count, err := s.repository.GetCountPendingToRegister(ctx)
	if err != nil {
		return 0, []dao.TransactionNFC{}, err
	}

	items, err := s.repository.GetListRegisterNFC(ctx, req)
	if err != nil {
		return 0, []dao.TransactionNFC{}, err
	}

	return count, items, nil
}
