package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yoratyo/material-handler-webapp/model/dao"
	transactionNFCDTO "github.com/yoratyo/material-handler-webapp/model/dto/transactionNFC"
	"time"
)

func (s *service) GetMonitoringGateway(ctx *gin.Context, conn *websocket.Conn) error {
	var (
		response  transactionNFCDTO.GetMonitoringGatewayResponseDTO
		header    dao.CountGatewayTransactionNFC
		okpDetail []dao.CountGatewayPerOKP
		err       error
	)

	ticker := time.NewTicker(3 * time.Second)
	defer func() {
		ticker.Stop()
		conn.Close()
	}()
	// we want to kick off a for loop that runs for the
	// duration of our websockets connection
	for {
		// every time our ticker ticks
		for t := range ticker.C {
			// print out that we are updating the stats
			fmt.Printf("[Gateway Monitoring] Updating Websocket connection %s: %+v\n", ctx.ClientIP(), t)

			header, err = s.repository.GetCountGatewayToday(ctx)
			if err != nil {
				return err
			}

			okpDetail, err = s.repository.GetCountGatewayTodayPerOKP(ctx)
			if err != nil {
				return err
			}

			currentTime := time.Now()
			today := currentTime.Format("Monday, 02 January 2006")

			response.Header.TodayDate = today
			response.Header.TotalChecked = header.TotalChecked
			response.Header.TotalRegistered = header.TotalRegistered

			var listOkp []transactionNFCDTO.ListOKPMonitoringGateway
			for _, okp := range okpDetail {
				o := transactionNFCDTO.ListOKPMonitoringGateway{
					BatchNo:            okp.BatchNo,
					FormulaDescription: okp.FormulaDescription,
					TotalChecked:       okp.TotalChecked,
					TotalRegistered:    okp.TotalRegistered,
				}
				listOkp = append(listOkp, o)
			}
			response.ListOKP = listOkp

			// next we marshal our response into a JSON string
			jsonString, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
				return err
			}

			// and finally we write this JSON string to our WebSocket
			// connection and record any errors if there has been any
			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
}
