package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
)

func (r *repository) GetByID(ctx *gin.Context, userID uint64) (dao.UserLogin, error) {
	var response dao.UserLogin

	if err := r.resource.DB.
		WithContext(ctx).
		Model(&dao.UserLogin{}).
		Where("id_user = ?", userID).
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
