package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	BaseModel
	Name           string `json:"name" gorm:"not null" valid:"required"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" valid:"required"`
	UserID         uint   `json:"user_id"`
	User           User
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) error {
	if _, err := govalidator.ValidateStruct(sm); err != nil {
		return err
	}

	return nil
}
