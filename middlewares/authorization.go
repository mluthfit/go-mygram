package middlewares

import (
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
func Authorization(param string, validate func(id uint, userID uint) (int, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var id = ctx.Param(param + "Id")
		var parseId, err = strconv.ParseUint(id, 10, 32)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": param + " id must be unsigned integer",
			})
			return
		}

		var userData = ctx.MustGet("userData").(jwt.MapClaims)
		var userID = userData["id"].(uint)

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
				fmt.Errorf(fmt.Sprintf("The photo id %d was not found", id))
		}

		if photo.UserID != userID {
			return http.StatusUnauthorized,
				fmt.Errorf(fmt.Sprintf("You are not allowed to access the %s data", name))
		}

		return http.StatusOK, nil
	}

	return Authorization(name, checkUserPhoto)
}

func CommentAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(*gin.Context) {}
}

func SocialMediaAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(*gin.Context) {}
}
