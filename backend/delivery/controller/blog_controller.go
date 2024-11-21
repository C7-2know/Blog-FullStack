package controller

import (
	"backend-starter-project/domain/dtos"
	"backend-starter-project/domain/interfaces"
	"fmt"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogService interfaces.BlogService
}

func (controller *BlogController) GetBlogPosts(c *gin.Context) {
	tag := c.QueryArray("tag")
	title := c.Query("title")
	author := c.Query("author")
	filter := dtos.BlogPostFilterDto{Tags: tag, Title: title, Author: author}
	result, err := controller.BlogService.GetBlogPosts(filter)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"blogs": result})
}

func (controller *BlogController) GetBlogPost(c *gin.Context) {
	id := c.Param("id")
	result, err := controller.BlogService.GetBlogPostById(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"blog": result})
}

func (controller *BlogController) CreateBlogPost(c *gin.Context) {
	
	var blog dtos.CreateBlogPostDto
	err := c.ShouldBindJSON(&blog)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("blog from controller", blog)
	fmt.Println("blog from controller", blog.ImageUrl)
	userId := c.GetString("userId")
	if userId == "" {
		c.JSON(400, gin.H{"error": "userId not found (doesn't set)"})
		return
	}
	blog.AuthorID = userId
	if blog.AuthorName == "" {
		author := c.GetString("userName")
		blog.AuthorName = author
	}
	result, err := controller.BlogService.CreateBlogPost(&blog, userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "blog post created", "blog": result})

}

func (controller *BlogController) UpdateBlogPost(c *gin.Context) {
	var blog dtos.UpdateBlogPostDto
	err := c.ShouldBindJSON(&blog)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	authorId := c.GetString("userId")
	blog.AuthorID = authorId
	result, err := controller.BlogService.UpdateBlogPost(&blog)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "blog post updated", "blog": result})
}
