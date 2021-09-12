package page

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) SignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
