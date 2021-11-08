package service

import "github.com/gin-gonic/gin"

func (s *service) GetDistinctPendingOKP(ctx *gin.Context) ([]string, error) {
	return s.repository.GetDistinctPendingOKP(ctx)
}
