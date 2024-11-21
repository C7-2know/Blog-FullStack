package repository

import (
	"backend-starter-project/domain/dtos"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type blogRepository struct {
	db *mongo.Database
}

func NewBlogRepository(db *mongo.Database) interfaces.BlogRepository {
	return &blogRepository{db: db}
}

func (repo *blogRepository) CreateBlogPost(blog *entities.BlogPost) (*entities.BlogPost, error) {
	_, err := repo.db.Collection("blogs").InsertOne(context.TODO(), blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (repo *blogRepository) GetBlogPostById(blogId string) (*entities.BlogPost, error) {
	var blog entities.BlogPost
	ID, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, err
	}
	res := repo.db.Collection("blogs").FindOne(context.TODO(), bson.M{"_id": ID})
	if res.Err() != nil {
		return nil, res.Err()
	}
	res.Decode(&blog)
	return &blog, nil
}

func (repo *blogRepository) DeleteBlogPost(blogId string) error {
	_, err := repo.db.Collection("blogs").DeleteOne(context.TODO(), bson.M{"_id": blogId})
	if err != nil {
		return err
	}
	return nil
}

func (repo *blogRepository) UpdateBlogPost(blog *dtos.UpdateBlogPostDto) (*entities.BlogPost, error) {
	var blogPost entities.BlogPost
	ID, err := primitive.ObjectIDFromHex(blog.ID)
	if err != nil {
		return nil, err
	}
	blogPost.ID = ID
	blogPost.Title = blog.Title
	blogPost.Content = blog.Content
	blogPost.Tags = blog.Tags
	blogPost.UpdatedAt = time.Now()

	filter := bson.M{"_id": blogPost.ID}
	_, err = repo.db.Collection("blogs").UpdateOne(context.TODO(), filter, bson.M{"$set": blogPost})
	if err != nil {
		return nil, err
	}
	return &blogPost, nil
}

// func (repo *blogRepository) GetBlogPosts(page, pageSize int, sortBy string) ([]entities.BlogPost, error) {
// 	var blogs []entities.BlogPost
// 	cursor, err := repo.db.Collection("blogs").Find(context.TODO(), bson.M{}, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	cursor.All(context.TODO(),blogs)
// 	return blogs, nil
// }

// Temporary blog post
func (repo *blogRepository) GetBlogPosts(dto dtos.BlogPostFilterDto) ([]entities.BlogPost, error) {
	var blogs []entities.BlogPost
	filter := bson.M{}
	if dto.Title != "" {
		filter["title"] = bson.M{"$regex": dto.Title, "$options": "i"}
	}
	if dto.Author != "" {
		filter["authorName"] = bson.M{"$regex": dto.Author, "$options": "i"}
	}
	if len(dto.Tags) > 0 {
		filter["tags"] = bson.M{"$in": dto.Tags}
	}
	fmt.Println("filter", filter)
	cursor, err := repo.db.Collection("blogs").Find(context.TODO(), filter, nil)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &blogs)
	if err != nil {
		return []entities.BlogPost{}, err
	}
	return blogs, nil
}

func (repo *blogRepository) SearchBlogPostsByAuthorName(name string) ([]entities.BlogPost, error) {
	var blogs []entities.BlogPost
	cursor, err := repo.db.Collection("blogs").Find(context.TODO(), bson.M{"authorName": name}, nil)
	if err != nil {
		return nil, err
	}
	cursor.All(context.TODO(), blogs)
	return blogs, nil
}

func (repo *blogRepository) SearchBlogPostsByTitle(title string) ([]entities.BlogPost, error) {
	var blogs []entities.BlogPost
	filter := bson.M{"$regex": title, "$options": "i"}
	cursor, err := repo.db.Collection("blogs").Find(context.TODO(), filter, nil)
	if err != nil {
		return nil, err
	}
	cursor.All(context.TODO(), blogs)
	return blogs, nil
}

//	func (repo *blogRepository) FilterBlogPosts(tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error) {
//		var blogs []entities.BlogPost
//		filter := bson.M{"tags": bson.M{"$in": tags}}
//		filter["createdAt"] = bson.M{"$gte": dateRange[0], "$lte": dateRange[1]}
//		cursor, err := repo.db.Collection("blogs").Find(context.TODO(), filter, nil)
//		if err != nil {
//			return nil, err
//		}
//		cursor.All(context.TODO(), blogs)
//		return blogs, nil
//	}
func (repo *blogRepository) FilterBlogPosts(tag string) ([]entities.BlogPost, error) {
	var blogs []entities.BlogPost
	filter := bson.M{"tags": bson.M{"$in": []string{tag}}}
	cursor, err := repo.db.Collection("blogs").Find(context.TODO(), filter, nil)
	if err != nil {
		return nil, err
	}
	cursor.All(context.TODO(), blogs)
	return blogs, nil
}

func (repo *blogRepository) LikeBlogPost(blogId, userId string) error {
	blog, err := repo.GetBlogPostById(blogId)
	if err != nil {
		return err
	}
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	if blog.Likes[id] {
		delete(blog.Likes, id)
		blog.LikeCount -= 1
	} else {
		blog.Likes[id] = true
		blog.LikeCount += 1
	}
	return nil
}

func (repo *blogRepository) DislikeBlogPost(blogId, userId string) error {
	blog, err := repo.GetBlogPostById(blogId)
	if err != nil {
		return err
	}
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	if blog.DisLikes[id] {
		delete(blog.DisLikes, id)
	}
	if blog.Likes[id] {
		delete(blog.Likes, id)
	}
	blog.DisLikes[id] = true

	return nil

}
