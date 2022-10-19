package controllers

import (
	"errors"
	"go-mygram/helpers"
	"go-mygram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := user.Create(s.DB); err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(201, gin.H{
		"age":      user.Age,
		"email":    user.Email,
		"id":       user.ID,
		"username": user.Username,
	})
}

func (s *Server) LoginUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	var payloadPass = user.Password
	if err := s.DB.Debug().Where("email = ?", user.Email).
		Take(&user).Error; err != nil {
		resError(ctx, http.StatusUnauthorized, errors.New("invalid email or password"))
		return
	}

	if comparePass := helpers.ComparePass(
		[]byte(user.Password), []byte(payloadPass),
	); !comparePass {
		resError(ctx, http.StatusUnauthorized, errors.New("invalid email or password"))
		return
	}

	ctx.JSON(200, gin.H{
		"token": helpers.GenerateToken(user.ID, user.Email),
	})
}

func (s *Server) UpdateUser(ctx *gin.Context) {
	var user, newUser models.User

	var userData = ctx.MustGet("userData").(jwt.MapClaims)
	var userID = uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	if newUser.Username == "" {
		resError(ctx, http.StatusBadRequest, errors.New("username: required"))
		return
	}

	user.ID = userID
	if err := user.Update(s.DB, newUser); err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	})
}

func (s *Server) DeleteUser(ctx *gin.Context) {
	var user models.User

	var userData = ctx.MustGet("userData").(jwt.MapClaims)
	var userID = uint(userData["id"].(float64))

	user.ID = userID

	if err := user.Delete(s.DB); err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "your account has been successfully deleted",
	})
}
