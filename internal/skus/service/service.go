package skuservice

import (
	"awesomeProject/Project/WMS/internal/domain"
	skurepository "awesomeProject/Project/WMS/internal/skus/repository"
	"context"
	"errors"
)

type SKUService struct {
	repo *skurepository.Repository
}

type Service interface {
	CreateSKU(ctx context.Context, sku domain.SKU) (domain.SKU, error)
	GetSKUs(ctx context.Context) []domain.SKU
	GetSKUByID(ctx context.Context, skuID int) (domain.SKU, error)
	GetSKUByTenantId(ctx context.Context, tenantId int) ([]domain.SKU, error)
	GetSkuBySellerID(ctx context.Context, sellerID int) ([]domain.SKU, error)
	DeleteSKU(ctx context.Context, skuID int) error
}

type service struct {
	repo skurepository.Repository
}

func NewService(r skurepository.Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateSKU(ctx context.Context, sku domain.SKU) (domain.SKU, error) {
	return s.repo.CreateSKU(ctx, sku)
}

func (s *service) GetSKUs(ctx context.Context) []domain.SKU {
	return s.repo.GetAllSkus(ctx)
}

func (s *service) GetSKUByID(ctx context.Context, SKU_ID int) (domain.SKU, error) {
	if SKU_ID <= 0 {
		return domain.SKU{}, errors.New("invalid SKU ID")
	}

	sku, err := s.repo.GetSkuByID(ctx, SKU_ID)
	if err != nil {
		return domain.SKU{}, err // Return an empty SKU and the error
	}
	return sku, nil
}

func (s *service) GetSKUByTenantId(ctx context.Context, tenantId int) ([]domain.SKU, error) {
	return s.repo.GetSKUByTenantId(ctx, tenantId)
}

func (s *service) GetSkuBySellerID(ctx context.Context, sellerID int) ([]domain.SKU, error) {
	if sellerID <= 0 {
		return nil, errors.New("invalid Seller ID")
	}

	skus, err := s.repo.GetSkuBySellerID(ctx, sellerID)
	if err != nil {
		return nil, err
	}
	return skus, nil
}

func (s *service) DeleteSKU(ctx context.Context, skuID int) error {
	return s.repo.DeleteSKU(ctx, skuID)
}
