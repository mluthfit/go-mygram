package controllers

import (
	"go-mygram/models"
	"net/http"
	"strconv"

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
	if err := photo.Create(s.db); err != nil {
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
	var photo models.Photo
	var photos, err = photo.GetAllWithUser(s.db)

	if err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(200, photos)
}

func (s *Server) UpdatePhoto(ctx *gin.Context) {
	var photoId = ctx.Param("photoId")
	var parsePhotoId, err = strconv.ParseUint(photoId, 10, 32)
	if err != nil {
		resError(ctx, http.StatusBadRequest, "photo id must be unsigned integer")
		return
	}

	var photo, payloadPhoto models.Photo
	if err := ctx.ShouldBindJSON(&payloadPhoto); err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	photo.ID = uint(parsePhotoId)
	if err := photo.Update(s.db, payloadPhoto); err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(200, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoUrl,
		"user_id":    photo.UserID,
		"updated_at": photo.UpdatedAt,
	})
}

func (s *Server) DeletePhoto(ctx *gin.Context) {
	var photoId = ctx.Param("photoId")
	var parsePhotoId, err = strconv.ParseUint(photoId, 10, 32)
	if err != nil {
		resError(ctx, http.StatusBadRequest, "photo id must be unsigned integer")
		return
	}

	var photo models.Photo
	photo.ID = uint(parsePhotoId)

	if err := photo.Delete(s.db); err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
