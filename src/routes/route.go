package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/remoteday/rd-api-go/src/platform"
)

// Handler represent the httphandler for healthcheck
type Handler struct {
	App platform.App
}

// NewWebAdapter -
func NewWebAdapter(app platform.App) http.Handler {

	r := gin.Default()

	NewHealthcheckHTTPHandler(r, app)

	NewTeamHTTPHandler(r, app)

	return r
}
