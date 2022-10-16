package controllers

import "github.com/gin-gonic/gin"

func resError(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"message": msg,
	})
}
