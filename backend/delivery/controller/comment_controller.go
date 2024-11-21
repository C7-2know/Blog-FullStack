package controller

import (
	"backend-starter-project/domain/dtos"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService interfaces.CommentService
}

func NewCommentController(commentService interfaces.CommentService) CommentController {
	return CommentController{
		commentService: commentService,
	}
}

func (controller *CommentController) AddComment(c *gin.Context){
	var comment dtos.CreateCommentRequestDTO
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	comment.AuthorID = c.GetString("userId")

	comm, err := controller.commentService.AddComment(&comment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, comm)
}

func (controller *CommentController) DeleteComment(commentId string) error {
	err := controller.commentService.DeleteComment(commentId)
	if err != nil {
		return err
	}
	return nil
}

func (controller *CommentController) GetCommentsByBlogPostId(c *gin.Context) ([]entities.Comment, error) {
	blogPostId := c.Param("blogPostId")
	comments, err := controller.commentService.GetCommentsByBlogPostId(blogPostId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
