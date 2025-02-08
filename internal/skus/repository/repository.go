package sku_repository

import (
	"awesomeProject/Project/WMS/internal/domain"
	"context"
	"errors"
	"github.com/omniful/go_commons/db/sql/postgres"
	"sync"
)

type Repository struct {
	db *postgres.DbCluster
}

var repo *Repository
var repoOnce sync.Once

func NewRepository(db *postgres.DbCluster) *Repository {
	repoOnce.Do(func() {
		repo = &Repository{
			db: db,
		}
	})
	return repo
}

type SKURepository interface {
	CreateSku(ctx context.Context, sku domain.SKU) (domain.SKU, error)
	GetSkuByID(ctx context.Context, SKU_ID int) (domain.SKU, error)
	GetSkuBySellerID(ctx context.Context, sellerID int) ([]domain.SKU, error)
	GetAllSkus(ctx context.Context) []domain.SKU
	DeleteSku(ctx context.Context, skuID int) error
}

func (r *Repository) GetAllSkus(ctx context.Context) []domain.SKU {
	var skus []domain.SKU
	r.db.GetMasterDB(ctx).Find(&skus)
	return skus
}

func (r *Repository) GetSkuByID(ctx context.Context, skuID int) (domain.SKU, error) {
	var sku domain.SKU

	if skuID <= 0 {
		return sku, errors.New("invalid SKU ID")
	}

	result := r.db.GetMasterDB(ctx).Where("id = ?", skuID).First(&sku)
	if result.Error != nil {
		return domain.SKU{}, result.Error
	}
	return sku, nil
}

func (r *Repository) GetSKUByTenantId(ctx context.Context, tenantId int) ([]domain.SKU, error) {
	var skus []domain.SKU
	if tenantId <= 0 {
		return skus, errors.New("invalid tenant ID")
	}
	result := r.db.GetMasterDB(ctx).Where("tenant_id = ?", tenantId).Find(&skus)
	if result.Error != nil {
		return skus, result.Error
	}
	return skus, nil
}

func (r *Repository) GetSkuBySellerID(ctx context.Context, sellerID int) ([]domain.SKU, error) {
	var skus []domain.SKU
	if sellerID <= 0 {
		return skus, errors.New("invalid Seller ID")
	}

	result := r.db.GetMasterDB(ctx).Where("seller_id = ?", sellerID).Find(&skus)
	if result.Error != nil {
		return nil, result.Error
	}

	return skus, nil
}

func (r *Repository) CreateSKU(ctx context.Context, sku domain.SKU) (domain.SKU, error) {
	result := r.db.GetMasterDB(ctx).Create(&sku)
	return sku, result.Error
}

func (r *Repository) DeleteSKU(ctx context.Context, SKU_ID int) error {
	result := r.db.GetMasterDB(ctx).Delete(&domain.SKU{}, SKU_ID)
	return result.Error
}
