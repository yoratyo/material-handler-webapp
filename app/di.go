package app

import (
	"github.com/yoratyo/material-handler-webapp/domain"
	userRepository "github.com/yoratyo/material-handler-webapp/domain/user/repository"
	userService "github.com/yoratyo/material-handler-webapp/domain/user/service"
	"github.com/yoratyo/material-handler-webapp/handlers"
	"github.com/yoratyo/material-handler-webapp/handlers/api"
	"github.com/yoratyo/material-handler-webapp/handlers/page"
	"github.com/yoratyo/material-handler-webapp/handlers/pageAction"
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
		domain.New,
		// Handlers
		api.NewHandler,
		page.NewHandler,
		pageAction.NewHandler,
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
