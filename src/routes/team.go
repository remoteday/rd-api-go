package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/remoteday/rd-api-go/src/platform"
	"github.com/remoteday/rd-api-go/src/team"
)

// NewTeamHTTPHandler -
func NewTeamHTTPHandler(r *gin.Engine, app platform.App) {
	handler := &Handler{
		App: app,
	}
	r.GET("/teams/:id", handler.get)
	r.GET("/teams", handler.list)
	r.POST("/teams", handler.create)
	r.PUT("/teams/:id", handler.replace)
	r.DELETE("/teams/:id", handler.delete)
}

// TODO: use dependency injection
func (h *Handler) get(c *gin.Context) {
	ctx := context.Background()
	usecase := h.App.Usecases.Team

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		fmt.Print(err)
	}

	response, err := usecase.FindByID(ctx, id)

	if err != nil {
		fmt.Print(err)
	}

	c.JSON(http.StatusOK, team.ToTeamDTO(response))
}

func (h *Handler) list(c *gin.Context) {
	ctx := context.Background()
	usecase := h.App.Usecases.Team

	response, err := usecase.FindAll(ctx)

	if err != nil {
		fmt.Print(err)
	}

	c.JSON(http.StatusOK, team.ToTeamDTOs(response))
}

func (h *Handler) create(c *gin.Context) {
	ctx := context.Background()
	var teamDto team.DTO

	err := c.BindJSON(&teamDto)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	response, err := h.App.Usecases.Team.Create(ctx, team.ToTeam(teamDto))

	c.JSON(http.StatusCreated, team.ToTeamDTO(response))
}

func (h *Handler) replace(c *gin.Context) {
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

	fmt.Println("responseresponseresponse", response)

	c.JSON(http.StatusOK, team.ToTeamDTO(response))
}

func (h *Handler) delete(c *gin.Context) {
	ctx := context.Background()
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		fmt.Print(err)
	}

	err = h.App.Usecases.Team.Delete(ctx, id)

	if err != nil {
		fmt.Print(err)
	}

	c.JSON(http.StatusNoContent, nil)
}
