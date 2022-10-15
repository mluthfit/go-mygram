package controllers

import (
	"errors"
	"go-mygram/helpers"
	"go-mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) UserRegister(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := s.db.Debug().Create(&user).Error; err != nil {
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

func (s *Server) UserLogin(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		resError(ctx, http.StatusBadRequest, err)
		return
	}

	var payloadPass = user.Password
	if err := s.db.Debug().Where("email = ?", user.Email).
		Take(&user).Error; err != nil {
		err = errors.New("invalid email or password")
		resError(ctx, http.StatusUnauthorized, err)
		return
	}

	if comparePass := helpers.ComparePass(
		[]byte(user.Password), []byte(payloadPass),
	); !comparePass {
		var err = errors.New("invalid email or password")
		resError(ctx, http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(200, gin.H{
		"token": helpers.GenerateToken(user.ID, user.Email),
	})
}

func (s *Server) UserUpdate(ctx *gin.Context) {

}
