package route

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func PrivateBlogRouter(db *mongo.Database, group gin.RouterGroup) {
	blogRepo := repository.NewBlogRepository(db)
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.BlogController{BlogService: blogService}
	group.POST("/blog", blogController.CreateBlogPost)
	group.PUT("/blog/:id", blogController.UpdateBlogPost)
}


func NewBlogRouter(db *mongo.Database, group gin.RouterGroup) {
	blogRepo := repository.NewBlogRepository(db)
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.BlogController{BlogService: blogService}
	
	group.GET("/blog/:id", blogController.GetBlogPost)
	group.GET("/blogs", blogController.GetBlogPosts)
	
}
