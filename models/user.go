package models

import (
	"go-mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username string `json:"username" gorm:"uniqueIndex;not null" valid:"required"`
	Email    string `json:"email" gorm:"uniqueIndex;not null" valid:"required,email"`
	Password string `json:"password" gorm:"not null" valid:"required,minstringlength(6)"`
	Age      uint   `json:"age" gorm:"not null" binding:"gt=8" valid:"required,numeric"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if _, err := govalidator.ValidateStruct(u); err != nil {
		return err
	}

	u.Password = helpers.HashPass(u.Password)
	return nil
}
