package middlewares

import (
	"errors"
	"fmt"
	"go-mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Param Validation
// One Data Checker Validation
// Authorization Validation
func Authorization(param, modelName string, validate func(id, userID uint) (int, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var id = ctx.Param(param + "Id")
		var parseId, err = strconv.ParseUint(id, 10, 32)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "the" + modelName + " id must be unsigned integer",
			})
			return
		}

		var userData = ctx.MustGet("userData").(jwt.MapClaims)
		var userID = uint(userData["id"].(float64))

		if code, err := validate(uint(parseId), userID); err != nil {
			ctx.AbortWithStatusJSON(code, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.Next()
	}
}

func PhotoAuthorization(db *gorm.DB) gin.HandlerFunc {
	var name = "photo"
	var checkUserPhoto = func(id uint, userID uint) (int, error) {
		var photo models.Photo
		var err = db.Select("user_id").First(&photo, id).Error

		if err != nil {
			return http.StatusBadRequest,
				fmt.Errorf(fmt.Sprintf("the %s id %d was not found", name, id))
		}

		if photo.UserID != userID {
			return http.StatusUnauthorized,
				errors.New("you are not allowed to access the data")
		}

		return http.StatusOK, nil
	}

	return Authorization("photo", name, checkUserPhoto)
}

func CommentAuthorization(db *gorm.DB) gin.HandlerFunc {
	var name = "comment"
	var checkUserComment = func(id, userID uint) (int, error) {
		var comment models.Comment
		var err = db.Select("user_id").First(&comment, id).Error

		if err != nil {
			return http.StatusBadRequest,
				fmt.Errorf(fmt.Sprintf("the %s id %d was not found", name, id))
		}

		if comment.UserID != userID {
			return http.StatusUnauthorized,
				errors.New("you are not allowed to access the data")
		}

		return http.StatusOK, nil
	}

	return Authorization("comment", name, checkUserComment)
}

func SocialMediaAuthorization(db *gorm.DB) gin.HandlerFunc {
	var name = "social media"
	var checkUserSocialMedia = func(id, userID uint) (int, error) {
		var socialMedia models.SocialMedia
		var err = db.Select("user_id").First(&socialMedia, id).Error

		if err != nil {
			return http.StatusBadRequest,
				fmt.Errorf(fmt.Sprintf("the %s id %d was not found", name, id))
		}

		if socialMedia.UserID != userID {
			return http.StatusUnauthorized,
				errors.New("you are not allowed to access the data")
		}

		return http.StatusOK, nil
	}

	return Authorization("socialMedia", name, checkUserSocialMedia)
}
