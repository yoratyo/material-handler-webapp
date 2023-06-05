package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/domain"
	"github.com/yoratyo/material-handler-webapp/shared"
)

type (
	Handler interface {
		GatewayCheckNFC(c *gin.Context)
		SchedulerGatewayCheckNFC(time time.Time)
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
