package page

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	userDTO "github.com/yoratyo/material-handler-webapp/model/dto/userLogin"
	"github.com/yoratyo/material-handler-webapp/shared"
	"net/http"
)

func (h handler) Gateway(c *gin.Context) {
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

	c.HTML(http.StatusOK, "gateway.html", gin.H{
		"isGatewayMenu": true,
		"username":      user.Username,
	})
}
