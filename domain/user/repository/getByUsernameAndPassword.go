package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	userDTO "github.com/yoratyo/material-handler-webapp/model/dto/userLogin"
)

func (r *repository) GetByUsernameAndPassword(ctx *gin.Context, req userDTO.LoginRequestDTO) (dao.UserLogin, error) {
	var response dao.UserLogin

	if err := r.resource.DB.
		WithContext(ctx).
		Model(&dao.UserLogin{}).
		Where("username = ?", req.Username).
		Where("password = ?", req.Password).
		Take(&response).Error; err != nil {
		//entry := log.WithFields(log.Fields{
		//	"method": ctx.Request.Method,
		//	"path":   ctx.Request.RequestURI,
		//	"status": ctx.Writer.Status(),
		//})
		//entry.Error(err.Error())
		return dao.UserLogin{}, err
	}

	return response, nil
}
