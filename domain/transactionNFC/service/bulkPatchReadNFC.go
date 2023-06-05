package service

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (s *service) BulkPatchReadNFC(ctx *gin.Context, tags []string) error {
	var updates = map[string]interface{}{
		"bitActive":          1,
		"dtmTransactionDate": time.Now(),
	}

	_, err := s.repository.PatchReadNFR(ctx, tags, updates)
	if err != nil {
		return err
	}

	//log.Printf("%d Row affected\n", rowsAffected)

	return nil
}
