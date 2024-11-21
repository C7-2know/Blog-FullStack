package route

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/infrastructure/middleware"
	"backend-starter-project/repository"
	"backend-starter-project/service"
	"context"

	"github.com/gin-gonic/gin"
)

func Setup(gin *gin.Engine){
	env:=bootstrap.NewEnv()
	client:=bootstrap.NewMongoDatabase(env)
	db:=client.Database(env.DBName)
	otp:=gin.Group("/otp")

	userRepo:=repository.NewUserRepository(db.Collection("users"))
	tokenRepo:=repository.NewTokenRepository(db)
	tokenService:=service.NewTokenService(env.AccessTokenSecret,env.RefreshTokenSecret,tokenRepo,userRepo)
	authmiddleware:=middleware.NewAuthMiddleware(tokenService)

	NewOTPRouter(otp)
	
	publicAuthRouter:=gin.Group("/auth")
	publicAuthRouter.Use()
	NewAuthRouter(*env,context.TODO(),db,*publicAuthRouter)

	
	publicBlogRouter:=gin.Group("")
	NewBlogRouter(db,*publicBlogRouter)
	
	privateBlogRouter:=gin.Group("")
	privateBlogRouter.Use(authmiddleware.AuthMiddleware(""))
	PrivateBlogRouter(db,*privateBlogRouter)

	profileRouter:=gin.Group("/user")
	profileRouter.Use(authmiddleware.AuthMiddleware(""))
	NewProfileRouter(db,*profileRouter)

	defaultImageRouter:=gin.Group("/default")
	// defaultImageRouter.Use(authmiddleware.AuthMiddleware("admin"))
	ImageRouter(db,env,defaultImageRouter)
	gin.Run(":8080")

}
