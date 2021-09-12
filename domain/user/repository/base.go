package repository

import (
	"github.com/yoratyo/material-handler-webapp/domain/user"
	"github.com/yoratyo/material-handler-webapp/shared"
)

type repository struct {
	resource shared.Resource
}

func New(resource shared.Resource) (user.Repository, error) {
	return &repository{
		resource: resource,
	}, nil
}
