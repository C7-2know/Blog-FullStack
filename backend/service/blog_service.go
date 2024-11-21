package service

import (
	"backend-starter-project/domain/dtos"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type blogService struct {
	Blog_repo interfaces.BlogRepository
}

func NewBlogService(blog_repo interfaces.BlogRepository) interfaces.BlogService {
	return &blogService{Blog_repo: blog_repo}
}

func (service *blogService) CreateBlogPost(blogPost *dtos.CreateBlogPostDto, userId string) (*entities.BlogPost, error) {
	authorId, err := primitive.ObjectIDFromHex(blogPost.AuthorID)
	if err != nil {
		return nil, err
	}
	blog := entities.BlogPost{
		Title:     blogPost.Title,
		Content:   blogPost.Content,
		Tags:      blogPost.Tags,
		AuthorID:  authorId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ImgUrl:    blogPost.ImageUrl,
	}
	fmt.Println("from service", blog)
	result, err := service.Blog_repo.CreateBlogPost(&blog)
	return result, nil
}
func (service *blogService) GetBlogPostById(blogPostId string) (*entities.BlogPost, error) {
	blog, err := service.Blog_repo.GetBlogPostById(blogPostId)
	if err != nil {
		return nil, err
	}
	return blog, nil
}
func (service *blogService) GetBlogPosts(filter dtos.BlogPostFilterDto) ([]entities.BlogPost, error) {
	blogs, err := service.Blog_repo.GetBlogPosts(filter)
	if err != nil {
		return []entities.BlogPost{}, err
	}
	return blogs, nil
}

func (service *blogService) UpdateBlogPost(blogPost *dtos.UpdateBlogPostDto) (*entities.BlogPost, error) {
	if blogPost.ID == "" {
		return nil, errors.New("Invalid blog post id")
	}

	blog, err := service.GetBlogPostById(blogPost.ID)
	if err != nil || blog.AuthorID.Hex() != blogPost.AuthorID {
		return nil, errors.New("only the author can edit the post")
	}

	res, err := service.Blog_repo.UpdateBlogPost(blogPost)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *blogService) DeleteBlogPost(blogPostId, userId, role string) error {
	err := service.Blog_repo.DeleteBlogPost(blogPostId)
	if err != nil {
		return err
	}
	return nil
}

//	func (service *blogService) GetBlogPosts(page, limit int, sortBy string) ([]entities.BlogPost, int, error) {
//		blogs,err:=service.Blog_repo.GetBlogPosts(page,limit,sortBy)
//		if err!=nil{
//			return nil,0,err
//		}
//		return blogs,0,nil
//	}
func (service *blogService) SearchBlogPostsByTitle(title string) ([]entities.BlogPost, error) {
	blogs, err := service.Blog_repo.SearchBlogPostsByTitle(title)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (service *blogService) SearchBlogPostsByAuthorName(authorName string) ([]entities.BlogPost, error) {
	blogs, err := service.Blog_repo.SearchBlogPostsByAuthorName(authorName)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (service *blogService) FilterBlogPosts(tag string) ([]entities.BlogPost, error) {
	blogs, err := service.Blog_repo.FilterBlogPosts(tag)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (service *blogService) LikeBlogPost(blogPostId, userId string) error {
	err := service.Blog_repo.LikeBlogPost(blogPostId, userId)
	if err != nil {
		return errors.ErrUnsupported
	}
	return nil
}

func (service *blogService) DislikeBlogPost(blogId, userId string) error {
	err := service.Blog_repo.DislikeBlogPost(blogId, userId)
	if err != nil {
		return err
	}
	return nil
}
