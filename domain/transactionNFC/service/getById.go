package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
)

func (s *service) GetByID(ctx *gin.Context, ID uint64) (dao.TransactionNFC, error) {
	return s.repository.GetByID(ctx, ID)
}
