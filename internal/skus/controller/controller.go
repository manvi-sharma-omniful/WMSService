package controller

import (
	"awesomeProject/Project/WMS/internal/domain"
	"awesomeProject/Project/WMS/internal/skus/service"
	//"context"
	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/http"
	oresponse "github.com/omniful/go_commons/response"
	"strconv"
)

type Controller struct {
	service skuservice.SKUService
}

func NewController(s skuservice.SKUService) *Controller {
	return &Controller{
		service: s,
	}
}

func (c *Controller) GetSKUs(ctx *gin.Context) gin.HandlerFunc {
	skus := c.service.GetSKUs(ctx)
	oresponse.NewSuccessResponse(ctx, skus)
}

func (c *Controller) GetSKUByID(ctx *gin.Context) gin.HandlerFunc {
	skuID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusBadRequest)
		return nil
	}

	sku, err := c.service.GetSKUByID(ctx, skuID)
	if err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusNotFound)
		return nil
	}
	oresponse.NewSuccessResponse(ctx, sku)

}

func (c *Controller) GetSKUBySellerID(ctx *gin.Context) gin.HandlerFunc {
	sellerID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusBadRequest)
		return nil
	}

	sku, err := c.service.GetSKUBySellerID(ctx, sellerID)
	if err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusNotFound)
		return nil
	}

	oresponse.NewSuccessResponse(ctx, sku)
	return nil
}

func (c *Controller) CreateSKU(ctx *gin.Context) gin.HandlerFunc {
	var sku domain.SKU
	if err := ctx.ShouldBindJSON(&sku); err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusBadRequest)
		return nil
	}

	newSKU, err := c.service.CreateSKU(ctx, sku)
	if err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusInternalServerError)
		return nil
	}

	oresponse.NewSuccessResponse(ctx, newSKU)
}

func (c *Controller) DeleteSKU(ctx *gin.Context) gin.HandlerFunc {
	skuID, err := strconv.Atoi(ctx.Param("sku_id"))
	if err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusBadRequest)
		return nil
	}

	err = c.service.DeleteSKU(ctx, skuID)
	if err != nil {
		oresponse.NewErrorResponseByStatusCode(ctx, http.StatusInternalServerError)
		return nil
	}

	oresponse.NewSuccessResponse(ctx, gin.H{"message": "SKU deleted successfully"})
	return nil
}
