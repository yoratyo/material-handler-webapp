package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	transactionNFCDTO "github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
)

func (s *service) GetRegisterNFC(ctx *gin.Context, request transactionNFCDTO.PatchRegisterNFCRequestDTO) (dao.TransactionNFC, error) {
	if request.ID != nil {
		return s.repository.GetByID(ctx, *request.ID)
	}

	return s.repository.GetRegisterNFC(ctx, request)
}
