package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Title        string               `bson:"title"`
	Content      string               `bson:"content"`
	AuthorName   string               `bson:"authorName"`
	AuthorID     primitive.ObjectID   `bson:"authorId"`
	ImgUrl       string               `bson:"imgUrl"`
	Tags         []string             `bson:"tags"`
	CreatedAt    time.Time            `bson:"createdAt"`
	UpdatedAt    time.Time            `bson:"updatedAt"`
	ViewCount    int                  `bson:"viewCount"`
	LikeCount    int                  `bson:"likeCount"`
	DislikeCount int                  `bson:"dislikeCount"`
	Likes        map[primitive.ObjectID]bool `bson:"likes"`
	DisLikes     map[primitive.ObjectID]bool `bson:"dislikes"`
	Comments []primitive.ObjectID `bson:"comments"`
}
type Like struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	BlogPostID primitive.ObjectID `bson:"blogPostId"`
	AuthorID   primitive.ObjectID `bson:"authorId"`
	CreatedAt  time.Time          `bson:"createdAt"`
}
type Dislike struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	BlogPostID primitive.ObjectID `bson:"blogPostId"`
	AuthorID   primitive.ObjectID `bson:"authorId"`
	CreatedAt  time.Time          `bson:"createdAt"`
}
