package models

type Category struct {
	BaseModel

	Name     string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}
