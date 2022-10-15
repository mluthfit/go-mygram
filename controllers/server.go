package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Router *gin.Engine
	db     *gorm.DB
}

func NewServer(router *gin.Engine, db *gorm.DB) *Server {
	return &Server{router, db}
}

func (s *Server) InitiateRoutes(router func(server *Server)) {
	router(s)
}

func (s *Server) Run(port string) {
	fmt.Printf("Server running at http://localhost%s\n", port)
	s.Router.Run(port)
}
