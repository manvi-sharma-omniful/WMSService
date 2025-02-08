package domain

import "time"

type Inventory struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	HubID             uint      `gorm:"not null"`
	TenantID          uint      `gorm:"not null"`
	SellerID          uint      `gorm:"not null"`
	SKU_ID            uint      `gorm:"not null"`
	AvailableQuantity int       `gorm:"not null;check:available_quantity >= 0"`
	LastUpdated       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
