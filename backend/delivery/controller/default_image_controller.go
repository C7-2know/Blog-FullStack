package controller

import (
	"backend-starter-project/domain/dtos"
	"backend-starter-project/domain/interfaces"

	"github.com/gin-gonic/gin"
)

type DefaultImageController struct {
	ImageService interfaces.DefaultImageService
}

func NewDefaultImageController(image_service interfaces.DefaultImageService) *DefaultImageController {
	return &DefaultImageController{ImageService: image_service}
}

func (c *DefaultImageController) GetDefaultImages(ctx *gin.Context) {
	images, err := c.ImageService.GetDefaultImages()
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response:=[]dtos.ImageResponseDto{}
	for _,image:=range images{
		response=append(response,dtos.ImageResponseDto{
			Id: image.Id.Hex(),
			Url: image.Url,
		})
	}
	ctx.JSON(200, gin.H{"images": response})
	return
}

func (c *DefaultImageController) GetDefaultImage(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "id is required"})
		return
	}
	image, err := c.ImageService.GetDefaultImage(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"image": image})
	return
}

func (c *DefaultImageController) CreateDefaultImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = c.ImageService.CreateDefaultImage(file)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Image added successfully"})
	return
}
func (c *DefaultImageController) DeleteDefaultImage(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "id is required"})
		return
	}

	err := c.ImageService.DeleteDefaultImage(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "error while deleting the image"})
		return
	}
	ctx.JSON(200, gin.H{"message": "successfully deleted"})

}
