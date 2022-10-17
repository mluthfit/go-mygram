package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	BaseModel
	Title    string `json:"title" gorm:"not null" valid:"required" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" gorm:"not null" valid:"required" binding:"required" `
	UserID   uint   `json:"user_id" gorm:"not null" valid:"required"`
	User     User   `valid:"-" binding:"-"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) error {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return err
	}

	return nil
}

func (p *Photo) GetAllWithUser(db *gorm.DB) (*[]Photo, error) {
	var photos []Photo

	if err := db.Debug().Preload("User").
		Find(&photos).Error; err != nil {
		return nil, err
	}

	return &photos, nil
}

func (p *Photo) Create(db *gorm.DB) error {
	return db.Debug().Create(p).Error
}

func (p *Photo) Update(db *gorm.DB, newPhoto Photo) error {
	if err := db.Debug().Model(p).
		Updates(newPhoto).Error; err != nil {
		return err
	}

	return db.First(p).Error
}

func (p *Photo) Delete(db *gorm.DB) error {
	return db.Debug().Delete(p).Error
}
