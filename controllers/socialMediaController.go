package controllers

import (
	"go-mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreateSocialMedia(ctx *gin.Context) {
	var socialMedia models.SocialMedia

	if err := ctx.ShouldBindJSON(&socialMedia); err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	var userData = ctx.MustGet("userData").(jwt.MapClaims)
	var userID = uint(userData["id"].(float64))

	socialMedia.UserID = userID
	if err := socialMedia.Create(s.DB); err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(201, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserID,
		"created_at":       socialMedia.CreatedAt,
	})
}

func (s *Server) GetAllSocialMedias(ctx *gin.Context) {
	var socialMedia models.SocialMedia
	var socialMedias, err = socialMedia.GetAllWithUser(s.DB)

	if err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, gin.H{
		"social_medias": socialMedias,
	})
}

func (s *Server) UpdateSocialMedia(ctx *gin.Context) {
	var socialMediaId = ctx.Param("socialMediaId")
	var parseSocialMediaId, _ = strconv.Atoi(socialMediaId)

	var socialMedia, payloadSocialMedia models.SocialMedia
	if err := ctx.ShouldBindJSON(&payloadSocialMedia); err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	socialMedia.ID = uint(parseSocialMediaId)
	if err := socialMedia.Update(s.DB, payloadSocialMedia); err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserID,
		"updated_at":       socialMedia.UpdatedAt,
	})
}

func (s *Server) DeleteSocialMedia(ctx *gin.Context) {
	var socialMediaId = ctx.Param("socialMediaId")
	var parseSocialMediaId, _ = strconv.Atoi(socialMediaId)

	var socialMedia models.SocialMedia
	socialMedia.ID = uint(parseSocialMediaId)

	if err := socialMedia.Delete(s.DB); err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "your social media has been succesfully deleted",
	})
}
