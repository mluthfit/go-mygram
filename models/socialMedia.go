package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	BaseModel
	Name           string `json:"name" gorm:"not null" valid:"required" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" valid:"required" binding:"required"`
	UserID         uint   `json:"user_id" gorm:"not null" valid:"required"`
	User           User   `valid:"-" binding:"-"`
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) error {
	if _, err := govalidator.ValidateStruct(sm); err != nil {
		return err
	}

	return nil
}

func (sm *SocialMedia) Create(db *gorm.DB) error {
	return db.Debug().Create(sm).Error
}

func (sm *SocialMedia) GetAllWithUser(db *gorm.DB) (*[]SocialMedia, error) {
	var socialMedias []SocialMedia

	if err := db.Debug().Preload("User").
		Find(&socialMedias).Error; err != nil {
		return nil, err
	}

	return &socialMedias, nil
}

func (sm *SocialMedia) Update(db *gorm.DB, newSocialMedia SocialMedia) error {
	if err := db.Debug().Model(sm).
		Updates(newSocialMedia).Error; err != nil {
		return err
	}

	return db.First(sm).Error
}

func (sm *SocialMedia) Delete(db *gorm.DB) error {
	return db.Debug().Delete(sm).Error
}
