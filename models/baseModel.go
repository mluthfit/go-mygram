package models

import "time"

type BaseModel struct {
	ID        uint      `json:"-" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
