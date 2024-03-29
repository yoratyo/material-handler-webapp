package pageAction

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/yoratyo/material-handler-webapp/domain"
	"github.com/yoratyo/material-handler-webapp/shared"
)

type (
	Handler interface {
		DoLogin(ctx *gin.Context)
		DoLogout(ctx *gin.Context)
		SubmitPickingSlip(ctx *gin.Context)
		SubmitRegisterNFC(ctx *gin.Context)
	}
	handler struct {
		domain   domain.Domain
		resource shared.Resource
	}
)

func NewHandler(domain domain.Domain, resource shared.Resource) (Handler, error) {
	return &handler{
		domain:   domain,
		resource: resource,
	}, nil
}

func addError(c *gin.Context, msg string, path string, params url.Values) {
	shared.SetErrorCookie(c, msg)
	location := url.URL{Path: path}
	if len(params) != 0 {
		location.RawQuery = params.Encode()
	}
	c.Redirect(http.StatusFound, location.RequestURI())
}
