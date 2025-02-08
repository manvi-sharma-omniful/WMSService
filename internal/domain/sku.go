package domain

import (
	"context"
)

type SKU struct {
	ID             uint                   `gorm:"primaryKey;autoIncrement"`
	ProductID      uint                   `gorm:"not null"`
	Price          float64                `gorm:"type:decimal(10,2);not null"`
	Fragile        string                 `gorm:"type:varchar(10);not null;check:fragile IN ('Yes', 'No')"`
	Specifications map[string]interface{} `gorm:"type:jsonb"`
}

type SKUService interface {
	GetSkus(ctx context.Context) []SKU
	FetchSkuByID(ctx context.Context, SKU_ID int) (SKU, error)
}
