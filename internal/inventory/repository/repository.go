package inventory_repository

import (
	"awesomeProject/Project/WMS/internal/domain"
	error3 "awesomeProject/Project/WMS/pkg/error"
	"context"
	"fmt"
	"github.com/omniful/go_commons/db/sql/postgres"
	error2 "github.com/omniful/go_commons/error"
	"sync"
)

type inventoryRepository struct {
	db *postgres.DbCluster
}

var repo *inventoryRepository
var repoOnce sync.Once

func NewRepository(db *postgres.DbCluster) *inventoryRepository {
	repoOnce.Do(func() {
		repo = &inventoryRepository{
			db: db,
		}
	})
	return repo
}

type InventoryRepository interface {
	GetInventory(ctx context.Context, hubID, skuID int) ([]domain.Inventory, error2.CustomError)
	UpdateInventory(ctx context.Context, inventory domain.Inventory) error2.CustomError
	ValidateInventory(ctx context.Context, skuID, hubID, quantity int) (bool, error2.CustomError)
}

func (r *inventoryRepository) GetInventory(ctx context.Context, hubID, skuID int, tenantID int) ([]domain.Inventory, error2.CustomError) {
	var inventories []domain.Inventory
	query := r.db.GetMasterDB(ctx)
	if hubID != 0 {
		query = query.Where("hub_id = ?", hubID)
	}
	if skuID != 0 {
		query = query.Where("sku_id = ?", skuID)
	}
	if tenantID != 0 {
		query = query.Where("tenant_id = ?", hubID)
	}

	inventoryResult := query.Find(&inventories)
	if resultErr := inventoryResult.Error; resultErr != nil {
		err := error2.NewCustomError(error3.SqlFetchError, fmt.Sprintf("Could not get inventory for condition: %+v, err: %v, Created its entry", resultErr))
		return []domain.Inventory{}, err
	}
	return inventories, error2.CustomError{}
}

func (r *inventoryRepository) UpdateInventory(ctx context.Context, inventory domain.Inventory) error2.CustomError {
	result := r.db.GetMasterDB(ctx).Model(&inventory).Where(" hub_id = ? AND sku_id = ?", inventory.HubID, inventory.SKU_ID).Updates(inventory)

	if resultErr := result.Error; resultErr != nil {
		err := error2.NewCustomError(error3.SqlUpdateError, fmt.Sprintf("Could not update inventory for condition: %+v, err: %v, Created its entry", resultErr))
		return err
	}

	if result.RowsAffected == 0 {
		return error2.NewCustomError(error3.SqlUpdateError, fmt.Sprintf("no inventory record found to update"))
	}
	return error2.CustomError{}
}
func (r *inventoryRepository) ValidateInventory(ctx context.Context, skuID, hubID, quantity int) (bool, error2.CustomError) {
	var totalQuantity int
	result := r.db.GetMasterDB(ctx).Table("inventories").Where("sku_id = ? AND hub_id = ?", skuID, hubID).
		Select("SUM(quantity)").Row().Scan(&totalQuantity)

	if result != nil {
		return false, error2.NewCustomError(error3.SqlFetchError, fmt.Sprintf("Cannot validate Inventory"))
	}

	return totalQuantity >= quantity, error2.CustomError{}
}
