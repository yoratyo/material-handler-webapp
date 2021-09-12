package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	userDTO "github.com/yoratyo/material-handler-webapp/model/dto/userLogin"
)

func (s *service) GetUserLogin(ctx *gin.Context, req userDTO.LoginRequestDTO) (dao.UserLogin, error) {
	return s.repository.GetByUsernameAndPassword(ctx, req)
}
