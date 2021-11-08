package repository

import (
	"github.com/yoratyo/material-handler-webapp/domain/transactionNFC"
	"github.com/yoratyo/material-handler-webapp/shared"
)

type repository struct {
	resource shared.Resource
}

func New(resource shared.Resource) (transactionNFC.Repository, error) {
	return &repository{
		resource: resource,
	}, nil
}
