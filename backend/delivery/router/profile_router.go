package route

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProfileRouter(db *mongo.Database, group gin.RouterGroup) {
	profileRepo := repository.NewProfileRepository(context.TODO(), db)
	profileService := service.NewProfileService(profileRepo)
	profileController := controller.ProfileController{ProfileService: profileService}

	group.GET("profiles", profileController.GetProfiles)
	group.GET("profile", profileController.GetProfile)
	group.GET("profile/:id", profileController.GetUserProfile)
	group.POST("profile", profileController.CreateUserProfile)
	group.PUT("profile/:id", profileController.UpdateUserProfile)
	group.DELETE("profile/:id", profileController.DeleteUserProfile)

}
