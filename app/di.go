package app

import (
	"github.com/yoratyo/material-handler-webapp/domain"
	materialRepository "github.com/yoratyo/material-handler-webapp/domain/masterMaterial/repository"
	materialService "github.com/yoratyo/material-handler-webapp/domain/masterMaterial/service"
	pickSlipRepository "github.com/yoratyo/material-handler-webapp/domain/pickingSlip/repository"
	pickSlipService "github.com/yoratyo/material-handler-webapp/domain/pickingSlip/service"
	transactionNFCRepository "github.com/yoratyo/material-handler-webapp/domain/transactionNFC/repository"
	transactionNFCService "github.com/yoratyo/material-handler-webapp/domain/transactionNFC/service"
	userRepository "github.com/yoratyo/material-handler-webapp/domain/user/repository"
	userService "github.com/yoratyo/material-handler-webapp/domain/user/service"
	"github.com/yoratyo/material-handler-webapp/handlers"
	"github.com/yoratyo/material-handler-webapp/handlers/api"
	"github.com/yoratyo/material-handler-webapp/handlers/page"
	"github.com/yoratyo/material-handler-webapp/handlers/pageAction"
	"github.com/yoratyo/material-handler-webapp/handlers/websocket"
	"github.com/yoratyo/material-handler-webapp/shared"
	"go.uber.org/dig"
)

func BuildRuntime() (*App, error) {
	c := dig.New()
	servicesConstructors := []interface{}{
		// Shared
		shared.NewDB,
		shared.NewResource,
		// Domain, Repository & Service
		userRepository.New,
		userService.New,
		pickSlipRepository.New,
		pickSlipService.New,
		transactionNFCRepository.New,
		transactionNFCService.New,
		materialRepository.New,
		materialService.New,
		domain.New,
		// Handlers
		api.NewHandler,
		page.NewHandler,
		pageAction.NewHandler,
		websocket.NewHandler,
		handlers.New,
		// App
		NewApp,
	}

	for _, service := range servicesConstructors {
		if err := c.Provide(service); err != nil {
			return nil, err
		}
	}

	var result *App
	err := c.Invoke(func(a *App) {
		result = a
	})
	return result, err
}
