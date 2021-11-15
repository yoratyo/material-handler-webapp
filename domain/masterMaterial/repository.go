package masterMaterial

import "github.com/gin-gonic/gin"

type Repository interface {
	PatchByID(ctx *gin.Context, ID string, data map[string]interface{}) error
}
