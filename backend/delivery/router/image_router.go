package route

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/delivery/controller"
	Infrastructure "backend-starter-project/infrastructure/cloudinary"
	"backend-starter-project/repository"
	"backend-starter-project/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ImageRouter(db *mongo.Database, env *bootstrap.Env, group *gin.RouterGroup) {
	imageRepo:=repository.NewDefaultImageRepository(db)
	cloudService:=Infrastructure.NewCloudinaryService(env.CloudName,env.ApiKey,env.ApiSec)
	imageService:=service.NewDefaultImageService(imageRepo,*cloudService)
	imageController:=controller.DefaultImageController{ImageService: imageService}
	group.GET("/images",imageController.GetDefaultImages)
	group.POST("/image",imageController.CreateDefaultImage)
	group.DELETE("/image/:id",imageController.DeleteDefaultImage)
	// group.PUT("/image/:id",imageController.UpdateDefaultImage)
}