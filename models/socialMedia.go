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

func (sm *SocialMedia) Create(db *gorm.DB) error {
	return db.Create(sm).Error
}

func (sm *SocialMedia) GetAllWithUser(db *gorm.DB) (*[]SocialMedia, error) {
	var socialMedias []SocialMedia

	if err := db.Preload("Users").Find(&socialMedias).Error; err != nil {
		return nil, err
	}

	return &socialMedias, nil
}

func (sm *SocialMedia) Update(db *gorm.DB, newSocialMedia SocialMedia) error {
	return db.Model(sm).Updates(newSocialMedia).Error
}

func (sm *SocialMedia) Delete(db *gorm.DB) error {
	return db.Delete(sm).Error
}
