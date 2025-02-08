package domain

type Product struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(40);not null"`
	Description string `gorm:"type:varchar(200)"`
}
