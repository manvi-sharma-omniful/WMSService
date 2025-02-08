package router

import (
	"awesomeProject/Project/WMS/internal/hubs/controller"
	hub_repository "awesomeProject/Project/WMS/internal/hubs/repository"
	"awesomeProject/Project/WMS/internal/hubs/service"
	"awesomeProject/Project/WMS/internal/inventory/controller"
	"awesomeProject/Project/WMS/internal/inventory/repository"
	"awesomeProject/Project/WMS/internal/inventory/service"
	"awesomeProject/Project/WMS/internal/skus/controller"
	"awesomeProject/Project/WMS/internal/skus/repository"
	"awesomeProject/Project/WMS/internal/skus/service"
	"awesomeProject/Project/WMS/pkg/db"
	"context"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
)

func InternalRoutes(ctx context.Context, s *http.Server) (err error) {

	s.Engine.Use(log.RequestLogMiddleware(log.MiddlewareOptions{
		Format:      config.GetString(ctx, "log.format"),
		Level:       config.GetString(ctx, "log.level"),
		LogRequest:  config.GetBool(ctx, "log.request"),
		LogResponse: config.GetBool(ctx, "log.response"),
	}))

	hub_engine := s.Engine.Group("/hub")
	sku_engine := s.Engine.Group("/sku")
	inventory_engine := s.Engine.Group("/inventory")

	hubRepository := hub_repository.NewRepository(pkg.GetCluster().DbCluster)
	hubService := hService.NewHubService(hubRepository)
	hubController := hcontroller.NewController(hubService)

	skuRepository := sku_repository.NewRepository(pkg.GetCluster().DbCluster)
	skuService := skuservice.NewService(skuRepository)
	skuController := controller.NewController(skuService)

	inventoryRepository := inventory_repository.NewInventoryRepository(pkg.GetCluster().DbCluster)
	inventoryService := iservice.NewInventoryService(inventoryRepository)
	inventoryController := icontroller.NewInventoryController(inventoryService)

	hub_engine.POST("/", hubController.CreateHub())
	//hub_engine.GET("/", hubController.GetHubs())
	hub_engine.GET("/:id", hubController.GetHubByID())
	hub_engine.DELETE("/:id", hubController.DeleteHub())

	sku_engine.POST("/", skuController.CreateSKU())
	sku_engine.GET("/", skuController.GetSKUs())
	sku_engine.GET("/:id", skuController.GetSKUByID())
	sku_engine.DELETE("/:id", skuController.DeleteSKU())
	sku_engine.GET("/seller/:id", skuController.GetSKUBySellerID())

	inventory_engine.GET("/", inventoryController.GetInventory())
	inventory_engine.PUT("/", inventoryController.UpdateInventory())
	inventory_engine.GET("/validate", inventoryController.ValidateInventory())

	return
}
