package models

import (
	"errors"
	"go-mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username string `json:"username" gorm:"uniqueIndex;not null" valid:"required"`
	Email    string `json:"email" gorm:"uniqueIndex;not null" valid:"required,email" binding:"required"`
	Password string `json:"password" gorm:"not null" valid:"required,minstringlength(6)"`
	Age      uint   `json:"age" gorm:"not null" valid:"required,numeric"`
	Photos   []Photo
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Age <= 8 {
		return errors.New("age: required and must be greater than 8")
	}

	if _, err := govalidator.ValidateStruct(u); err != nil {
		return err
	}

	u.Password = helpers.HashPass(u.Password)
	return nil
}

func (u *User) Create(db *gorm.DB) error {
	return db.Debug().Create(u).Error
}

func (u *User) Update(db *gorm.DB, newUser User) error {
	if err := db.Debug().Model(u).Updates(newUser).Error; err != nil {
		return err
	}

	return db.Debug().First(u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Debug().Delete(u).Error
}
