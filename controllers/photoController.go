package controllers

import (
	"go-mygram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreatePhoto(ctx *gin.Context) {
	var photo models.Photo

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var userData = ctx.MustGet("userData").(jwt.MapClaims)
	var userID = userData["id"].(uint)

	photo.UserID = userID
	if err := s.db.Create(&photo).Error; err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(201, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserID,
		"created_at": photo.CreatedAt,
	})
}

func (s *Server) GetAllPhotos(ctx *gin.Context) {
	var photos []models.Photo

	if err := s.db.Find(&photos).Error; err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(200, photos)
}

func (s *Server) UpdatePhoto(ctx *gin.Context) {
	var photoId = ctx.Param("paramId")
}

func (s *Server) DeletePhoto(ctx *gin.Context) {

}
