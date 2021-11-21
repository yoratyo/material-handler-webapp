package page

import "C"
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

	//c.Writer.WriteHeader(200)
	//indexHtml, _ := asset.Asset("template/login.html")
	////_, _ = c.Writer.Write(indexHtml)
	////c.Writer.Header().Add("Accept", "text/html")
	////c.Writer.Flush()
	//
	//c.Header("Content-Type", "text/html;  charset=utf-8")
	//c.String(200, string(indexHtml))

	c.HTML(http.StatusOK, "login.html", response)
}
