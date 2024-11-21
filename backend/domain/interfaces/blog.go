package interfaces

import (
	"backend-starter-project/domain/dtos"
	"backend-starter-project/domain/entities"
)

type BlogRepository interface {
	CreateBlogPost(blogPost *entities.BlogPost) (*entities.BlogPost, error)
	GetBlogPostById(blogPostId string) (*entities.BlogPost, error)
	UpdateBlogPost(blogPost *dtos.UpdateBlogPostDto) (*entities.BlogPost, error)
	DeleteBlogPost(blogPostId string) error
	// GetBlogPosts(page, pageSize int, sortBy string) ([]entities.BlogPost, error)
	GetBlogPosts(filter dtos.BlogPostFilterDto) ([]entities.BlogPost, error)
	SearchBlogPostsByAuthorName(authorName string) ([]entities.BlogPost, error)
	SearchBlogPostsByTitle(title string) ([]entities.BlogPost, error)
	// SearchBlogPostsByContent(content string) ([]entities.BlogPost, error)

	FilterBlogPosts(tag string) ([]entities.BlogPost, error)
	// FilterBlogPosts(tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error)
	LikeBlogPost(blogId, userId string) error
	DislikeBlogPost(blogId, userId string) error
}

type BlogService interface {
	CreateBlogPost(blogPost *dtos.CreateBlogPostDto, userId string) (*entities.BlogPost, error)
	GetBlogPostById(blogPostId string) (*entities.BlogPost, error)
	UpdateBlogPost(blogPost *dtos.UpdateBlogPostDto) (*entities.BlogPost, error)
	DeleteBlogPost(blogPostId, userId, role string) error
	// GetBlogPosts(page, pageSize int, sortBy string) ([]entities.BlogPost,int, error)
	GetBlogPosts(filter dtos.BlogPostFilterDto) ([]entities.BlogPost, error)
	SearchBlogPostsByAuthorName(authorName string) ([]entities.BlogPost, error)
	SearchBlogPostsByTitle(title string) ([]entities.BlogPost, error)
	FilterBlogPosts(tag string) ([]entities.BlogPost, error)
	// FilterBlogPosts(tags []string, dateRange []time.Time, sortBy string) ([]entities.BlogPost, error)
	LikeBlogPost(blogId, userId string) error
	DislikeBlogPost(blogId, userId string) error
}

type PopularityTrackingService interface {
	IncrementViewCount(blogPostId string) error
	LikeBlogPost(blogPostId, userId string) error
	DislikeBlogPost(blogPostId, userId string) error
	GetPopularityMetrics(blogPostId string) (map[string]int, error)
}
