package app

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/shared"
	"net/http"
)

func NewRouter(a *App) *gin.Engine {
	router := a.Engine

	router.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	// Test routes
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// Page routes
	page := router.Group("/page")
	// Required auth
	page.GET("/home", shared.AuthRequired, a.Handler.Page.Home)
	page.GET("/picking", shared.AuthRequired, a.Handler.Page.Picking)
	page.GET("/nfc", shared.AuthRequired, a.Handler.Page.NFC)
	page.GET("/gateway", shared.AuthRequired, a.Handler.Page.Gateway)
	// Public route
	page.GET("/login", shared.Public, a.Handler.Page.SignIn)

	// Page action routes
	action := router.Group("/action")
	// Auth action
	auth := action.Group("/auth")
	auth.POST("/login", a.Handler.PageAction.DoLogin)
	auth.GET("/logout", a.Handler.PageAction.DoLogout)
	// PickingSlip action
	pickingSlip := action.Group("/pickingSlip")
	pickingSlip.POST("/submit", a.Handler.PageAction.SubmitPickingSlip)
	// Transaction NFC action
	transactionNFC := action.Group("/transaction-NFC")
	transactionNFC.POST("/submit", a.Handler.PageAction.SubmitRegisterNFC)

	// Websocket action routes
	websocket := router.Group("/ws")
	websocket.GET("/statusNFC", a.Handler.WebSocket.StatusNFC)

	// API action routes
	api := router.Group("/api")
	api.PATCH("/gateway-checked-nfc", a.Handler.Api.GatewayCheckNFC)

	return router
}
