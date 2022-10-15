package controllers

import (
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
