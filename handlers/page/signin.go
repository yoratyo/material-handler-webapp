package page

import (
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/shared"
	"net/http"
)

func (h handler) SignIn(c *gin.Context) {
	response := gin.H{}

	errMessage, err := c.Cookie("error")
	if err == nil {
		response["errMessage"] = errMessage
		shared.RemoveErrorCookie(c)
	}

	c.HTML(http.StatusOK, "login.html", response)
}
