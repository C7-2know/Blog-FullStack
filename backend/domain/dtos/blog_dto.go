package dtos

import "time"

type CreateBlogPostDto struct {
	Title      string   `json:"title" binding:"required"`
	Content    string   `json:"content"`
	AuthorName string   `json:"authorName" `
	Tags       []string `json:"tags"`
	AuthorID   string   `json:"-"`
	ImageUrl   string   `json:"imgUrl"`
}

type UpdateBlogPostDto struct {
	ID        string   `json:"id" binding:"required"`
	Title     string   `json:"title" `
	Content   string   `json:"content"`
	ImgUrl    string   `json:"imgUrl"`
	Tags      []string `json:"tags"`
	AuthorID  string   `json:"-"`
	UpdatedAt time.Time	`json:"-" `
}

type BlogPostFilterDto struct{
	Title string `json:"title"`
	Author string `json:"author"`
	Tags []string `json:"tags"`
}