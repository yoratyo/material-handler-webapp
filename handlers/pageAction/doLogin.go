package pageAction

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	userDTO "github.com/yoratyo/material-handler-webapp/model/dto/userLogin"
	"github.com/yoratyo/material-handler-webapp/shared"
	"net/http"
)

func (h handler) DoLogin(c *gin.Context) {
	session := sessions.Default(c)

	// validate form input
	var req userDTO.LoginRequestDTO
	errs := c.ShouldBind(&req)
	if errs != nil {
		shared.SetErrorCookie(c, "username & password required")
		c.Redirect(http.StatusFound, "/page/login")
		c.Abort()
		return
	}

	// auth login
	res, err := h.domain.User.GetUserLogin(c, req)
	if err != nil {
		shared.SetErrorCookie(c, "username & password not match")
		c.Redirect(http.StatusFound, "/page/login")
		c.Abort()
		return
	}

	// Save the user data in the session
	session.Set(shared.USERKEY, res.IDUser)
	if err := session.Save(); err != nil {
		shared.SetErrorCookie(c, "failed to save session user")
		c.Redirect(http.StatusFound, "/page/login")
		c.Abort()
		return
	}

	c.Redirect(http.StatusFound, "/page/home")
	c.Abort()
}
