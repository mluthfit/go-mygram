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

func (p *Photo) GetAllWithUser(db *gorm.DB) (*[]Photo, error) {
	var photos *[]Photo

	if err := db.Preload("User").Find(photos).Error; err != nil {
		return nil, err
	}

	return photos, nil
}

func (p *Photo) Create(db *gorm.DB) error {
	return db.Create(p).Error
}

func (p *Photo) Update(db *gorm.DB, newPhoto Photo) error {
	return db.Model(p).Updates(newPhoto).Error
}

func (p *Photo) Delete(db *gorm.DB) error {
	return db.Delete(p).Error
}
