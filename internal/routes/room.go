package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/remoteday/rd-api-go/internal/platform"
	"github.com/remoteday/rd-api-go/internal/room"
)

// HandlerRoom -
type HandlerRoom struct {
	App platform.App
}

// NewRoomHTTPHandler -
func NewRoomHTTPHandler(r *gin.Engine, app platform.App) {
	handler := &HandlerRoom{
		App: app,
	}
	r.GET("/rooms/:id", handler.get)
	r.GET("/rooms", handler.list)
}

func (h *HandlerRoom) get(c *gin.Context) {
	ctx := context.Background()

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": http.StatusBadRequest})
		return
	}

	response, err := h.App.Usecases.Room.FindByID(ctx, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found", "status": http.StatusNotFound})
		return
	}

	c.JSON(http.StatusOK, room.ToRoomDTO(response))
}

func (h *HandlerRoom) list(c *gin.Context) {
	ctx := context.Background()

	response, err := h.App.Usecases.Room.FindAll(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, room.ToRoomDTOs(response))
}
