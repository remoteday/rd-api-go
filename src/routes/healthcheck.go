package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/remoteday/rd-api-go/src/platform"
)

// NewHealthcheckHTTPHandler -
func NewHealthcheckHTTPHandler(r *gin.Engine, app platform.App) {
	handler := &Handler{
		App: app,
	}
	r.GET("/_health", handler.Healthcheck)
	r.GET("/", handler.Healthcheck)
}

// Healthcheck -
func (a *Handler) Healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
