package controller

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileService interfaces.ProfileService
}

func (controller *ProfileController) CreateUserProfile(ctx *gin.Context) {
	var profile entities.Profile
	err := ctx.ShouldBindJSON(&profile)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	controller.ProfileService.CreateUserProfile(&profile)
	ctx.JSON(200, gin.H{"message": "Profile created successfully"})

}
func (controller *ProfileController) GetProfiles(ctx *gin.Context) {
	profile, err := controller.ProfileService.GetProfiles()
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, profile)
}

func (controller *ProfileController) GetProfile(ctx *gin.Context) {
	userId:=ctx.GetString("userId")
	profile, err := controller.ProfileService.GetUserProfile(userId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Profile retrieved successfully", "profile": profile})
}


func (controller *ProfileController) GetUserProfile(ctx *gin.Context) {
	userId := ctx.Param("id")
	profile, err := controller.ProfileService.GetUserProfile(userId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, profile)
}

func (controller *ProfileController) UpdateUserProfile(ctx *gin.Context) {
	var profile entities.Profile
	err := ctx.ShouldBindJSON(&profile)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updated,err:=controller.ProfileService.UpdateUserProfile(&profile)
	if err!=nil{
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Profile updated successfully", "profile": updated})
}

func (controller *ProfileController) DeleteUserProfile(ctx *gin.Context) {
	userId := ctx.Param("userId")
	err := controller.ProfileService.DeleteUserProfile(userId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Profile deleted successfully"})
}