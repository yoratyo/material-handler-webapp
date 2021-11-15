package handlers

import (
	"github.com/yoratyo/material-handler-webapp/handlers/api"
	"github.com/yoratyo/material-handler-webapp/handlers/page"
	"github.com/yoratyo/material-handler-webapp/handlers/pageAction"
	"github.com/yoratyo/material-handler-webapp/handlers/websocket"
)

type Handler struct {
	Api        api.Handler
	Page       page.Handler
	PageAction pageAction.Handler
	WebSocket  websocket.Handler
}

func New(api api.Handler, page page.Handler, pageAction pageAction.Handler, webSocket websocket.Handler) (Handler, error) {
	return Handler{
		Api:        api,
		Page:       page,
		PageAction: pageAction,
		WebSocket:  webSocket,
	}, nil
}
