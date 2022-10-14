package middlewares

import (
	"go-mygram/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var verifyToken, err = helpers.VerifyToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})

			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
