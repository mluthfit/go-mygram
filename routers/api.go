package routers

import "github.com/gin-gonic/gin"

func ApiRoutes(router *gin.Engine) {
	var userRouter = router.Group("/users")
	{
		userRouter.POST("/register", func(ctx *gin.Context) {})
	}
}
