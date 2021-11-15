package masterMaterial

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	PatchTotalQty(ctx *gin.Context, ID string) error
}
