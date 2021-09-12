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
		c.JSON(http.StatusBadRequest, userDTO.LoginResponseDTO{
			Message: "Payload failed",
		})
		return
	}

	// auth login
	res, err := h.domain.User.GetUserLogin(c, req)
	if err != nil {
		c.JSON(http.StatusNotFound, userDTO.LoginResponseDTO{
			Message: "Failed Login",
		})
		return
	}

	// Save the user data in the session
	session.Set(shared.USERKEY, res.IDUser)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.Redirect(http.StatusFound, "/page/home")
	c.Abort()
}
