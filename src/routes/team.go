package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/remoteday/rd-api-go/src/platform"
	"github.com/remoteday/rd-api-go/src/team"
	log "github.com/sirupsen/logrus"
)

// HandlerTeam
type HandlerTeam struct {
	App platform.App
}

// NewTeamHTTPHandler -
func NewTeamHTTPHandler(r *gin.Engine, app platform.App) {
	handler := &HandlerTeam{
		App: app,
	}
	r.GET("/teams/:id", handler.get)
	r.GET("/teams", handler.list)
	r.POST("/teams", handler.create)
	r.PUT("/teams/:id", handler.replace)
	r.DELETE("/teams/:id", handler.delete)
}

func (h *HandlerTeam) get(c *gin.Context) {
	ctx := context.Background()

	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": http.StatusBadRequest})
		return
	}

	response, err := h.App.Usecases.Team.FindByID(ctx, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found", "status": http.StatusNotFound})
		return
	}

	c.JSON(http.StatusOK, team.ToTeamDTO(response))
}

func (h *HandlerTeam) list(c *gin.Context) {
	ctx := context.Background()
	usecase := h.App.Usecases.Team

	response, err := usecase.FindAll(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, team.ToTeamDTOs(response))
}

func (h *HandlerTeam) create(c *gin.Context) {
	ctx := context.Background()
	var teamDto team.DTO

	err := c.BindJSON(&teamDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "status": http.StatusInternalServerError})
	}

	response, err := h.App.Usecases.Team.Create(ctx, team.ToTeam(teamDto))

	c.JSON(http.StatusCreated, team.ToTeamDTO(response))
}

func (h *HandlerTeam) replace(c *gin.Context) {
	ctx := context.Background()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		fmt.Print(err)
	}

	var teamDto team.DTO

	err = c.BindJSON(&teamDto)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	response, err := h.App.Usecases.Team.Update(ctx, id, team.ToTeam(teamDto))

	c.JSON(http.StatusOK, team.ToTeamDTO(response))
}

func (h *HandlerTeam) delete(c *gin.Context) {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		fmt.Print(err)
	}

	err = h.App.Usecases.Team.Delete(ctx, id)

	if err != nil {
		log.Error(err)
	}

	c.JSON(http.StatusNoContent, nil)
}
