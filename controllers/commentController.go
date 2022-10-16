package controllers

import (
	"go-mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreateComment(ctx *gin.Context) {
	var comment models.Comment

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var userData = ctx.MustGet("userData").(jwt.MapClaims)
	var userID = userData["id"].(uint)

	comment.UserID = userID
	if err := comment.Create(s.DB); err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(201, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	})
}

func (s *Server) GetAllComments(ctx *gin.Context) {
	var comment models.Comment
	var comments, err = comment.GetAllWithUserAndPhoto(s.DB)

	if err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(200, comments)
}

func (s *Server) UpdateComment(ctx *gin.Context) {
	var commentId = ctx.Param("commentId")
	var parseCommentId, _ = strconv.Atoi(commentId)

	var comment, payloadComment models.Comment
	if err := ctx.ShouldBindJSON(&payloadComment); err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	comment.ID = uint(parseCommentId)
	if err := comment.Update(s.DB, payloadComment); err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(200, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"updated_at": comment.UpdatedAt,
	})
}

func (s *Server) DeleteComment(ctx *gin.Context) {
	var commentId = ctx.Param("commentId")
	var parseCommentId, _ = strconv.Atoi(commentId)

	var comment models.Comment
	comment.ID = uint(parseCommentId)

	if err := comment.Delete(s.DB); err != nil {
		resError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Your comment has been succesfully deleted",
	})
}
