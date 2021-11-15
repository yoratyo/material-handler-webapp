package service

import (
	"github.com/gin-gonic/gin"
	helper "github.com/yoratyo/material-handler-webapp/shared"
	"gorm.io/gorm"
)

func (s *service) PatchTotalQty(ctx *gin.Context, ID string) error {
	currDate, _ := helper.GetCurrentTime()

	var updates = map[string]interface{}{
		"total_qty":     gorm.Expr("total_qty + weight_pack"),
		"modified_date": currDate,
	}

	return s.repository.PatchByID(ctx, ID, updates)
}
