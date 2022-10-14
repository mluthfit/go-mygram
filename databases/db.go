package databases

import (
	"fmt"
	"go-mygram/helpers"
	"go-mygram/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	var dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		helpers.GetEnv("DB_USERNAME"),
		helpers.GetEnv("DB_PASSWORD"),
		helpers.GetEnv("DB_HOST"),
		helpers.GetEnv("DB_PORT"),
		helpers.GetEnv("DB_NAME"),
	)

	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Debug().AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
	return db
}
