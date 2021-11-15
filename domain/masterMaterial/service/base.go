package service

import (
	"github.com/yoratyo/material-handler-webapp/domain/masterMaterial"
	"github.com/yoratyo/material-handler-webapp/shared"
)

type service struct {
	repository masterMaterial.Repository
	resource   shared.Resource
}

func New(repository masterMaterial.Repository, resource shared.Resource) (masterMaterial.Service, error) {
	return &service{
		repository: repository,
		resource:   resource,
	}, nil
}
