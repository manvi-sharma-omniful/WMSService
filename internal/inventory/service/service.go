package iservice

import (
	"awesomeProject/Project/WMS/internal/domain"
	hub_repository "awesomeProject/Project/WMS/internal/hubs/repository"
	inventory_repository "awesomeProject/Project/WMS/internal/inventory/repository"
	error3 "awesomeProject/Project/WMS/pkg/error"
	"context"
	"fmt"
	error2 "github.com/omniful/go_commons/error"
)

type inventoryService struct {
	repo inventory_repository.InventoryRepository
}

func NewInventoryService(repo inventory_repository.InventoryRepository) *inventoryService {
	return &inventoryService{repo: repo}
}

type InventoryService interface {
	GetInventory(ctx context.Context, hubID, SKU_ID int) ([]domain.Inventory, error2.CustomError)
	UpdateInventory(ctx context.Context, inventory domain.Inventory) error2.CustomError
	ValidateInventory(ctx context.Context, skuID, hubID, quantity int) (bool, error2.CustomError)
}

func (s *inventoryService) GetInventory(ctx context.Context, hubID, SKU_ID int) ([]domain.Inventory, error2.CustomError) {
	return s.repo.GetInventory(ctx, hubID, SKU_ID)
}
func (s *inventoryService) UpdateInventory(ctx context.Context, inventory domain.Inventory) error2.CustomError {
	if inventory.HubID == 0 || inventory.SKU_ID == 0 {
		return error2.NewCustomError(error3.SqlUpdateError, fmt.Sprintf("Could not update inventory, Created its entry"))
	}
	return s.repo.UpdateInventory(ctx, inventory)
}

func (s *inventoryService) ValidateInventory(ctx context.Context, SKU_ID, hubID, quantity int) (bool, error2.CustomError) {
	return s.repo.ValidateInventory(ctx, SKU_ID, hubID, quantity)
}

func NewService(repository *hub_repository.Repository) interface{} {

}
