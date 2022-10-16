package models

import (
	"fmt"

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

func (c *Comment) Create(db *gorm.DB) error {
	if err := db.First(&Photo{}, c.PhotoID).Error; err != nil {
		var name = "photo"
		return fmt.Errorf(fmt.Sprintf("The %s id %d was not found", name, c.PhotoID))
	}

	return db.Create(c).Error
}

func (c *Comment) GetAllWithUserAndPhoto(db *gorm.DB) (*[]Comment, error) {
	var comments []Comment

	if err := db.Preload("Users").Preload("Photos").
		Find(&comments).Error; err != nil {
		return nil, err
	}

	return &comments, nil
}

func (c *Comment) Update(db *gorm.DB, newComment Comment) error {
	return db.Model(c).Updates(newComment).Error
}

func (c *Comment) Delete(db *gorm.DB) error {
	return db.Delete(c).Error
}
