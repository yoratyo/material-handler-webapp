package shared

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(USERKEY)
	if user == nil {
		// Abort the request with the appropriate error code
		//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})

		c.Redirect(http.StatusFound, "/page/login")
		c.Abort()
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

func Public(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(USERKEY)
	if user != nil {
		c.Redirect(http.StatusFound, "/page/home")
		c.Abort()
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}
