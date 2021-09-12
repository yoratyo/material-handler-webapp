package pageAction

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/shared"
	"net/http"
)

func (h handler) DoLogout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(shared.USERKEY)
	//fmt.Println("Logout>>>>>>>>>>>>>>>>>")
	//
	//entry := log.WithFields(log.Fields{
	//	"method":     c.Request.Method,
	//	"path":       c.Request.RequestURI,
	//	"status":     c.Writer.Status(),
	//})
	//entry.Info("Logout>>>>>>>")

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(shared.USERKEY)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.Redirect(http.StatusFound, "/page/login")
	c.Abort()
}
