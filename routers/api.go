package routers

import (
	"go-mygram/controllers"
)

func ApiRoutes(s *controllers.Server) {
	var userRouter = s.Router.Group("/users")
	{
		userRouter.POST("/register", s.UserRegister)
	}
}
