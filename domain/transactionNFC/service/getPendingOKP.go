package service

import "github.com/gin-gonic/gin"

func (s *service) GetPendingOKP(ctx *gin.Context, status string) ([]string, error) {
	return s.repository.GetDistinctOKP(ctx, status)
}
