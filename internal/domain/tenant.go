package domain

import "time"

type Tenant struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	Name          string    `gorm:"type:varchar(100);not null"`
	ContactNumber string    `gorm:"type:varchar(10);not null"`
	Status        string    `gorm:"type:enum('Active', 'Suspended', 'Terminated');not null"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
