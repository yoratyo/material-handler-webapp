package domain

import (
	"github.com/yoratyo/material-handler-webapp/domain/user"
)

type Domain struct {
	User user.Service
}

func New(user user.Service) (Domain, error) {
	return Domain{
		User: user,
	}, nil
}
