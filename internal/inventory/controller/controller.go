package icontroller

import (
	"awesomeProject/Project/WMS/internal/domain"
	"awesomeProject/Project/WMS/internal/inventory/service"
	"github.com/gin-gonic/gin"
	oresponse "github.com/omniful/go_commons/response"
	"net/http"
	"strconv"
)

type InventoryController struct {
	service service.InventoryService
}

func NewController(s service.InventoryService) *InventoryController {
	return &InventoryController{
		service: s,
	}
}

func (c *InventoryController) GetInventory(ctx *gin.Context) {
	hubID, _ := strconv.Atoi(ctx.Query("hub_id"))
	skuID, _ := strconv.Atoi(ctx.Query("sku_id"))

	inventories, err := c.service.GetInventory(ctx, hubID, skuID)
	if err.Exists() {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusInternalServerError)
		return
	}
	oresponse.NewSuccessResponse(ctx, inventories)
}

func (c *InventoryController) UpdateInventory(ctx *gin.Context) {
	var inventory domain.Inventory
	if err := ctx.ShouldBindJSON(&inventory); err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusBadRequest)
		return
	}

	err := c.service.UpdateInventory(ctx, inventory)
	if err.Exists() {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusInternalServerError)
		return
	}
	oresponse.NewSuccessResponse(ctx, inventory)
}

func (c *InventoryController) ValidateInventory(ctx *gin.Context) {
	skuID, _ := strconv.Atoi(ctx.Query("sku_id"))
	hubID, _ := strconv.Atoi(ctx.Query("hub_id"))
	quantity, _ := strconv.Atoi(ctx.Query("quantity"))

	if skuID == 0 || hubID == 0 || quantity <= 0 {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusBadRequest)
		return
	}

	isAvailable, err := c.service.ValidateInventory(ctx, skuID, hubID, quantity)
	if err.Exists() {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusInternalServerError)
		return
	}
	oresponse.NewSuccessResponse(ctx, isAvailable)
}
