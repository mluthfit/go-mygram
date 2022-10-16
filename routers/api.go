package routers

import (
	"go-mygram/controllers"
	"go-mygram/middlewares"
)

func ApiRoutes(s *controllers.Server) {
	var usersRouter = s.Router.Group("/users")
	{
		usersRouter.POST("/register", s.RegisterUser)
		usersRouter.POST("/login", s.LoginUser)

		usersRouter.Use(middlewares.Authentication())
		usersRouter.PUT("/", s.UpdateUser)
		usersRouter.DELETE("/", s.DeleteUser)
	}

	var photosRouter = s.Router.Group("/photos")
	{
		photosRouter.Use(middlewares.Authentication())
		photosRouter.POST("/", s.CreatePhoto)
		photosRouter.GET("/", s.GetAllPhotos)
		photosRouter.PUT("/:photoId", middlewares.PhotoAuthorization(s.DB), s.UpdatePhoto)
		photosRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(s.DB), s.DeletePhoto)
	}

	var commentsRouter = s.Router.Group("/comments")
	{
		commentsRouter.Use(middlewares.Authentication())
		commentsRouter.POST("/", s.CreateComment)
		commentsRouter.GET("/", s.GetAllComments)
		commentsRouter.PUT("/:commentId", middlewares.CommentAuthorization(s.DB), s.UpdateComment)
		commentsRouter.DELETE("/:commentId", middlewares.CommentAuthorization(s.DB), s.DeleteComment)
	}

	var socialMediasRouter = s.Router.Group("/socialMedias")
	{
		socialMediasRouter.Use(middlewares.Authentication())
		socialMediasRouter.POST("/", s.CreateSocialMedia)
		socialMediasRouter.GET("/", s.GetAllSocialMedias)
		socialMediasRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(s.DB), s.UpdateSocialMedia)
		socialMediasRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(s.DB), s.DeleteSocialMedia)
	}
}
