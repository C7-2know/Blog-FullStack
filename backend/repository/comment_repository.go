package repository

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type commentRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewCommentRepository(db *mongo.Database) interfaces.CommentRepository {
	return &commentRepository{
		db:         db,
		collection: db.Collection("comments"),
	}
}

func (r *commentRepository) AddComment(comment *entities.Comment) (*entities.Comment, error) {
	comment.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(context.Background(), comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *commentRepository) DeleteComment(commentId string) error {
	id, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(context.Background(), primitive.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (r *commentRepository) GetCommentsByBlogPostId(blogPostId string) ([]entities.Comment, error) {
	id, err := primitive.ObjectIDFromHex(blogPostId)
	if err != nil {
		return nil, err
	}
	cursor, err := r.collection.Find(context.Background(), primitive.M{"blogPostId": id})
	if err != nil {
		return nil, err
	}
	var comments []entities.Comment
	err = cursor.All(context.Background(), &comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepository) UpdateComment(comment *entities.Comment) (*entities.Comment, error) {
	if comment.ID == primitive.NilObjectID {
		return nil, errors.New("Invalid comment")
	}
	comment.Updated = true
	_, err := r.collection.ReplaceOne(context.TODO(), primitive.M{"_id": comment.ID}, comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
