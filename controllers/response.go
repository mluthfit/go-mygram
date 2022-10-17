package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func resError(ctx *gin.Context, code int, err error) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"message": strings.ToLower(err.Error()),
	})
}
