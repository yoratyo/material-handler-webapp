package service

import (
	"github.com/yoratyo/material-handler-webapp/domain/transactionNFC"
	"github.com/yoratyo/material-handler-webapp/shared"
)

type service struct {
	repository transactionNFC.Repository
	resource   shared.Resource
}

func New(repository transactionNFC.Repository, resource shared.Resource) (transactionNFC.Service, error) {
	return &service{
		repository: repository,
		resource:   resource,
	}, nil
}
