package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/shared"
)

func (h handler) StatusNFC(c *gin.Context) {
	ws, err := shared.NewWebsocketConnection(c.Writer, c.Request)
	if err != nil {
		fmt.Fprintf(c.Writer, "%+v\n", err)
	}

	go h.domain.TransactionNFC.SetTickerMonitoringNFC(c, ws)
}
