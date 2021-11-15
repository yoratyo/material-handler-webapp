package repository

import (
	"github.com/yoratyo/material-handler-webapp/domain/masterMaterial"
	"github.com/yoratyo/material-handler-webapp/shared"
)

type repository struct {
	resource shared.Resource
}

func New(resource shared.Resource) (masterMaterial.Repository, error) {
	return &repository{
		resource: resource,
	}, nil
}
