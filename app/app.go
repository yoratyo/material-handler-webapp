package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/handlers"
	"os"
)

type App struct {
	Engine *gin.Engine
	Handler handlers.Handler
}

func (a *App) Start() error {
	engine := NewRouter(a)

	engine.LoadHTMLGlob("templates/*.html")

	engine.Static("/page/css", "assets/css")
	engine.Static("/page/fonts", "assets/fonts")
	engine.Static("/page/img", "assets/img")
	engine.Static("/page/js", "assets/js")
	engine.Static("/page/scss", "assets/scss")
	engine.Static("/page/vendor", "assets/vendor")

	return engine.Run(":" + os.Getenv("APP_PORT"))
}

func NewApp(handler handlers.Handler) *App {
	return &App{
		Engine: gin.Default(),
		Handler: handler,
	}
}
