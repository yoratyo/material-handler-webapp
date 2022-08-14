package service

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	"gorm.io/gorm"
)

func (s *service) SetTickerMonitoringNFC(ctx *gin.Context, conn *websocket.Conn) error {
	var (
		nfc dao.MonitoringNFC
		err error
	)

	ticker := time.NewTicker(2 * time.Second)
	defer func() {
		ticker.Stop()
		conn.Close()
	}()
	// we want to kick off a for loop that runs for the
	// duration of our websockets connection
	for {
		// every time our ticker ticks
		for range ticker.C {
			// print out that we are updating the stats
			// fmt.Printf("[NFC status] Updating Websocket connection %s: %+v\n", ctx.ClientIP(), t)

			nfc, err = s.repository.GetMonitoringNFCByIP(ctx, ctx.ClientIP())
			if err != nil {
				// fmt.Println(err)
				if errors.Is(err, gorm.ErrRecordNotFound) {
					nfc, err = s.repository.GetDefaultMonitoringNFC(ctx)
					if err != nil {
						// fmt.Println(err)
						return err
					}
				} else {
					return err
				}
			}
			// next we marshal our response into a JSON string
			jsonString, err := json.Marshal(nfc)
			if err != nil {
				// fmt.Println(err)
				return err
			}

			// and finally we write this JSON string to our WebSocket
			// connection and record any errors if there has been any
			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
				// fmt.Println(err)
				return err
			}
		}
	}
}
