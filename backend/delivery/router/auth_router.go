package route

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/delivery/controller"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"backend-starter-project/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

func NewAuthRouter(env bootstrap.Env,ctx context.Context, db *mongo.Database, group gin.RouterGroup) {
	tokenRepo := repository.NewTokenRepository(db)
	userRepo := repository.NewUserRepository(db.Collection("users"))
	profileRepo := repository.NewProfileRepository(ctx, db)

	tokenService:=service.NewTokenService(env.AccessTokenSecret,env.RefreshTokenSecret,tokenRepo,userRepo)
	passService := utils.NewPasswordService()
	userService := service.NewUserService(userRepo, profileRepo)
	authService := service.NewAuthService(userService, tokenRepo, passService, tokenService)
	authController:=controller.AuthController{AuthService:authService}
	group.POST("/register",authController.RegisterUser)	
	group.POST("/login",authController.Login)
	
}
