package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/remoteday/rd-api-go/src/platform"
)

// NewWebAdapter -
func NewWebAdapter(app platform.App) http.Handler {

	r := gin.Default()

	NewHealthcheckHTTPHandler(r, app)

	NewTeamHTTPHandler(r, app)

	NewRoomHTTPHandler(r, app)

	return r
}
