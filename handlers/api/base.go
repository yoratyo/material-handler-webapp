package api

import (
	"github.com/yoratyo/material-handler-webapp/domain"
	"github.com/yoratyo/material-handler-webapp/shared"
)

type (
	Handler interface {

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
