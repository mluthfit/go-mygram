package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	BaseModel
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message" gorm:"not null" valid:"required"`
	User    User   `json:"user"`
	Photo   Photo  `json:"photo"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	if _, err := govalidator.ValidateStruct(c); err != nil {
		return err
	}

	return nil
}
