package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *service) PatchTotalQty(ctx *gin.Context, ID string) error {
	var updates = map[string]interface{}{
		"total_qty":     gorm.Expr("total_qty + weight_pack"),
		"modified_date": time.Now(),
	}

	return s.repository.PatchByID(ctx, ID, updates)
}
