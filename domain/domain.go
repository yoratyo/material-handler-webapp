package domain

import (
	"github.com/yoratyo/material-handler-webapp/domain/pickingSlip"
	"github.com/yoratyo/material-handler-webapp/domain/transactionNFC"
	"github.com/yoratyo/material-handler-webapp/domain/user"
)

type Domain struct {
	User           user.Service
	PickingSlip    pickingSlip.Service
	TransactionNFC transactionNFC.Service
}

func New(user user.Service, pickingSlip pickingSlip.Service, transactionNFC transactionNFC.Service) (Domain, error) {
	return Domain{
		User:           user,
		PickingSlip:    pickingSlip,
		TransactionNFC: transactionNFC,
	}, nil
}
