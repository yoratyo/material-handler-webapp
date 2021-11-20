package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"math"
)

func (s *service) GetTodayRegisterMonitoring(ctx *gin.Context) (dao.RegisterNFCMonitoring, error) {
	register, err := s.repository.GetRegisterMonitoringToday(ctx)
	if err != nil {
		return dao.RegisterNFCMonitoring{}, err
	}

	if register.TotalCount == 0 {
		register.Percentage = 0
	} else {
		percentage := (float64(register.ProgressCount) / float64(register.TotalCount)) * 100
		register.Percentage = math.Ceil(percentage*100) / 100
	}

	return register, nil
}
