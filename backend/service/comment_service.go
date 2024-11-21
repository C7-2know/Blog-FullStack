package service

import (
	"backend-starter-project/domain/dtos"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commentService struct {
	comment_repo interfaces.CommentRepository
}

func NewCommentRepository(comment_repo interfaces.CommentRepository) interfaces.CommentService {
	return &commentService{
		comment_repo: comment_repo,
	}
}

func (service *commentService) AddComment(comment_dto *dtos.CreateCommentRequestDTO) (*entities.Comment, error) {
	BlogPostID,err:=primitive.ObjectIDFromHex(comment_dto.BlogPostID)
	if err!=nil{
		return nil,err
	}
	AuthorID,err:=primitive.ObjectIDFromHex(comment_dto.AuthorID)
	if err!=nil{
		return nil,err
	}
	comment:=entities.Comment{
		BlogPostID: BlogPostID,
		Content: comment_dto.Content,
		AuthorID: AuthorID,
	}
	comm,err:=service.comment_repo.AddComment(&comment)
	if err!=nil{
		return nil,err
	}
	return comm,nil
}

func (service *commentService) DeleteComment(commentId string) error {
	err:=service.comment_repo.DeleteComment(commentId)
	if err!=nil{
		return err
	}
	return nil
}

func (service *commentService) GetCommentsByBlogPostId(blogPostId string) ([]entities.Comment, error) {
	comments,err:=service.comment_repo.GetCommentsByBlogPostId(blogPostId)
	if err!=nil{
		return nil,err
	}
	return comments,nil
}

func (service *commentService) UpdateComment(comment *entities.Comment) (*entities.Comment, error) {
	comm,err:=service.comment_repo.UpdateComment(comment)
	if err!=nil{
		return nil,err
	}
	return comm,nil
}
