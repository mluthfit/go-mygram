package models

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	BaseModel
	UserID  uint   `json:"user_id" gorm:"not null" valid:"required"`
	PhotoID uint   `json:"photo_id" gorm:"not null" valid:"required"`
	Message string `json:"message" gorm:"not null" valid:"required" binding:"required"`
	User    User   `valid:"-" binding:"-"`
	Photo   Photo  `valid:"-" binding:"-"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	if _, err := govalidator.ValidateStruct(c); err != nil {
		return err
	}

	return nil
}

func (c *Comment) Create(db *gorm.DB) error {
	if err := db.Debug().First(&Photo{}, c.PhotoID).
		Error; err != nil {
		var name = "photo"
		return fmt.Errorf(fmt.Sprintf("The %s id %d was not found", name, c.PhotoID))
	}

	return db.Create(c).Error
}

func (c *Comment) GetAllWithUserAndPhoto(db *gorm.DB) ([]map[string]any, error) {
	var comments []Comment

	if err := db.Debug().Preload("User").
		Preload("Photo").Find(&comments).
		Error; err != nil {
		return nil, err
	}

	return c.mappingGetAll(comments), nil
}

func (c *Comment) Update(db *gorm.DB, newComment Comment) error {
	if err := db.Debug().Model(c).Updates(newComment).
		Error; err != nil {
		return err
	}

	return db.First(c).Error
}

func (c *Comment) Delete(db *gorm.DB) error {
	return db.Debug().Delete(c).Error
}

func (c *Comment) mappingGetAll(comments []Comment) (results []map[string]any) {
	for _, comment := range comments {
		var data = map[string]any{
			"id":         comment.ID,
			"message":    comment.Message,
			"photo_id":   comment.PhotoID,
			"user_id":    comment.UserID,
			"created_at": comment.CreatedAt,
			"updated_at": comment.UpdatedAt,
			"User": map[string]any{
				"id":       comment.User.ID,
				"email":    comment.User.Email,
				"username": comment.User.Username,
			},
			"Photo": map[string]any{
				"id":        comment.Photo.ID,
				"title":     comment.Photo.Title,
				"caption":   comment.Photo.Caption,
				"photo_url": comment.Photo.PhotoUrl,
				"user_id":   comment.Photo.UserID,
			},
		}

		results = append(results, data)
	}

	return
}
