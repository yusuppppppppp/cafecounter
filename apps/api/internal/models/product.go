package models

type Product struct {
	BaseModel

	CategoryID  uint     `gorm:"not null;index"`
	Category    Category `gorm:"foreignKey:CategoryID"`
	Name        string   `gorm:"type:varchar(150);not null"`
	Description string   `gorm:"type:text"`
	Price       int64    `gorm:"not null"`
	ImageURL    string   `gorm:"type:text"`
	IsAvailable bool     `gorm:"default:true"`
}
