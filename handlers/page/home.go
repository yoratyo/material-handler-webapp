package page

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	userDTO "github.com/yoratyo/material-handler-webapp/model/dto/userLogin"
	"github.com/yoratyo/material-handler-webapp/shared"
	"net/http"
	"time"
)

func (h handler) Home(c *gin.Context) {
	session := sessions.Default(c)
	userSession := session.Get(shared.USERKEY)
	userID, ok := userSession.(uint64)
	if !ok {
		c.JSON(http.StatusNotFound, userDTO.LoginResponseDTO{
			Message: "Failed get session",
		})
		return
	}

	// auth me
	user, err := h.domain.User.GetByID(c, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, userDTO.LoginResponseDTO{
			Message: "Failed get user data",
		})
		return
	}

	response := gin.H{
		"isHomeMenu": true,
		"username":   user.Username,
	}

	// get monitoring picking slip
	picking, err := h.domain.PickingSlip.GetTodayPickingMonitoring(c)
	if err != nil {
		response["errMessage"] = "Failed get monitoring picking"
		return
	}

	// get monitoring register NFC
	register, err := h.domain.TransactionNFC.GetTodayRegisterMonitoring(c)
	if err != nil {
		response["errMessage"] = "Failed get monitoring register"
		return
	}

	// get monitoring Gateway
	gateway, err := h.domain.TransactionNFC.GetTodayGatewayMonitoring(c)
	if err != nil {
		response["errMessage"] = "Failed get monitoring gateway"
		return
	}

	currentTime := time.Now()
	today := currentTime.Format("Monday, 02 January 2006")

	response["today"] = today
	response["picking"] = picking
	response["register"] = register
	response["gateway"] = gateway

	errMessage, err := c.Cookie("error")
	if err == nil {
		response["errMessage"] = errMessage
		shared.RemoveErrorCookie(c)
	}

	c.HTML(http.StatusOK, "index.html", response)
}
