package service

import (
	"github.com/yoratyo/material-handler-webapp/domain/pickingSlip"
	"github.com/yoratyo/material-handler-webapp/shared"
)

type service struct {
	repository pickingSlip.Repository
	resource   shared.Resource
}

func New(repository pickingSlip.Repository, resource shared.Resource) (pickingSlip.Service, error) {
	return &service{
		repository: repository,
		resource:   resource,
	}, nil
}
