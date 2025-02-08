package domain

type Seller struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(60);not null"`
	Email string `gorm:"type:varchar(200);unique;not null"`
}
