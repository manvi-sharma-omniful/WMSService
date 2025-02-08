package hcontroller

import (
	"awesomeProject/Project/WMS/internal/domain"
	hubService "awesomeProject/Project/WMS/internal/hubs/service"
	"strconv"

	"github.com/omniful/go_commons/http"
	oresponse "github.com/omniful/go_commons/response"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service hubService.HubService
}

func NewController(s hubService.HubService) *Controller {
	return &Controller{
		service: s,
	}
}

func (c *Controller) GetHubByID(ctx *gin.Context) {
	hubID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusBadRequest)
		return
	}

	hub, err := c.service.GetHubByID(ctx, uint(hubID))
	if err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusNotFound)
		return
	}
	oresponse.NewSuccessResponse(ctx, hub)
}

func (c *Controller) CreateHub(ctx *gin.Context) {
	var hub domain.Hub
	if err := ctx.ShouldBindJSON(&hub); err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusBadRequest)
		return
	}
	newHub, err := c.service.CreateHub(ctx, hub)
	if err.Exists() {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusInternalServerError)
		return
	}

	oresponse.NewSuccessResponse(ctx, newHub)
}

func (c *Controller) DeleteHub(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusBadRequest)
		return
	}

	err = c.service.DeleteHub(ctx, id)
	if err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusInternalServerError)
		return
	}

	oresponse.NewSuccessResponse(ctx, gin.H{"message": "Hub deleted successfully"})
}
