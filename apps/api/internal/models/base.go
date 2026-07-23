package models

import "time"

type BaseModel struct {
	ID       uint `gorm:"primaryKey"`
	CreateAt time.Time
	UpdateAt time.Time
}
