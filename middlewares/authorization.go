package middlewares

import "github.com/gin-gonic/gin"

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return Authorization()
}

func CommentAuthorization() gin.HandlerFunc {
	return Authorization()
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return Authorization()
}
