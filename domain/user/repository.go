package user

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	userDTO "github.com/yoratyo/material-handler-webapp/model/dto/userLogin"
)

type Repository interface {
	GetByUsernameAndPassword(ctx *gin.Context, req userDTO.LoginRequestDTO) (dao.UserLogin, error)
	GetByID(ctx *gin.Context, userID uint64) (dao.UserLogin, error)
}
