package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	BaseModel
	Title    string `json:"title" gorm:"not null" valid:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" gorm:"not null" valid:"required"`
	UserID   uint   `json:"user_id"`
	User     User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return err
	}

	return nil
}
