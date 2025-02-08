package domain

import (
	"awesomeProject/Project/WMS/internal/hubs/responses"
	"context"
	error2 "github.com/omniful/go_commons/error"
)

type Hub struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"type:varchar(60);not null"`
	TenantID uint   `gorm:"not null"`
	Address1 string `gorm:"type:varchar(400)"`
	Address2 string `gorm:"type:varchar(400)"`
	City     string `gorm:"type:varchar(100)"`
	State    string `gorm:"type:varchar(100)"`
	Country  string `gorm:"type:varchar(100)"`
	Pincode  string `gorm:"type:varchar(10)"`
}
