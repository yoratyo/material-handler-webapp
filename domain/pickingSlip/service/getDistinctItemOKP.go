package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
)

func (s *service) GetDistinctItemOKP(ctx *gin.Context, batchNo string) ([]dao.ItemPickingSlip, error) {
	return s.repository.GetDistinctItemOKP(ctx, batchNo)
}
