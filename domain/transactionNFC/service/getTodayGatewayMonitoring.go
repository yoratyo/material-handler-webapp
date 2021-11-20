package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"math"
)

func (s *service) GetTodayGatewayMonitoring(ctx *gin.Context) (dao.GatewayMonitoring, error) {
	gateway, err := s.repository.GetCountGatewayToday(ctx)
	if err != nil {
		return dao.GatewayMonitoring{}, err
	}

	var percentage float64 = 0

	if gateway.TotalRegistered != 0 {
		percentage = (float64(gateway.TotalChecked) / float64(gateway.TotalRegistered)) * 100
	}

	gatewayRes := dao.GatewayMonitoring{
		ProgressCount: gateway.TotalChecked,
		TotalCount:    gateway.TotalRegistered,
		Percentage:    math.Ceil(percentage*100) / 100,
	}

	return gatewayRes, nil
}
