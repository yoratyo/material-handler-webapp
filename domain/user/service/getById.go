package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
)

func (s *service) GetByID(ctx *gin.Context, userID uint64) (dao.UserLogin, error) {
	return s.repository.GetByID(ctx, userID)
}
