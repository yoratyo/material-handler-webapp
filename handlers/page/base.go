package page

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/domain"
	"github.com/yoratyo/material-handler-webapp/shared"
)

type (
	Handler interface {
		Home(ctx *gin.Context)
		SignIn(ctx *gin.Context)
		Picking(ctx *gin.Context)
		NFC(ctx *gin.Context)
		Gateway(ctx *gin.Context)
	}
	handler struct {
		domain   domain.Domain
		resource shared.Resource
	}
)

func NewHandler(domain domain.Domain, resource shared.Resource) (Handler, error) {
	return &handler{
		domain:   domain,
		resource: resource,
	}, nil
}
