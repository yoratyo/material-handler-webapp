package service

import (
	"github.com/yoratyo/material-handler-webapp/domain/user"
	"github.com/yoratyo/material-handler-webapp/shared"
)

type service struct {
	repository user.Repository
	resource   shared.Resource
}

func New(repository user.Repository, resource shared.Resource) (user.Service, error) {
	return &service{
		repository: repository,
		resource:   resource,
	}, nil
}
