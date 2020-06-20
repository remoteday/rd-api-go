package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/remoteday/rd-api-go/src/platform"
)

// HandlerHealthcheck -
type HandlerHealthcheck struct {
	App platform.App
}

// NewHealthcheckHTTPHandler -
func NewHealthcheckHTTPHandler(r *gin.Engine, app platform.App) {
	handler := &HandlerHealthcheck{
		App: app,
	}
	r.GET("/_health", handler.Healthcheck)
	r.GET("/", handler.Healthcheck)
}

// Healthcheck -
func (a *HandlerHealthcheck) Healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
