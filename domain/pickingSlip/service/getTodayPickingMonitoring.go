package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"math"
)

func (s *service) GetTodayPickingMonitoring(ctx *gin.Context) (dao.PickingMonitoring, error) {
	picking, err := s.repository.GetPickingMonitoringToday(ctx)
	if err != nil {
		return dao.PickingMonitoring{}, err
	}

	if picking.TotalCount == 0 {
		picking.Percentage = 0
	} else {
		percentage := (float64(picking.ProgressCount) / float64(picking.TotalCount)) * 100
		picking.Percentage = math.Ceil(percentage*100) / 100
	}

	return picking, nil
}
