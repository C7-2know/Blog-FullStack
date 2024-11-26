package controller

import (
	"backend-starter-project/domain/dtos"
	"backend-starter-project/domain/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService interfaces.AuthenticationService
}

func (controller *AuthController) RegisterUser(c *gin.Context) {
	var user dtos.RegisterUserDto
	err := c.ShouldBindJSON(&user)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": "missing data", "input": user})
		return
	}

	created, err := controller.AuthService.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created", "user": created})

}

func (controller *AuthController) Login(c *gin.Context) {
	var data dtos.UserLoginDto
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	refreshToken, accessToken,user, err := controller.AuthService.Login(data.Email, data.Password)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	user_res:=dtos.UserResponseDto{
		ID: user.ID,
		Email: user.Email,
		Username: user.Username,
		Role: user.Role,
	}
	c.SetCookie("access_token", accessToken, 3600, "/", "localhost", false, false)
	c.SetCookie("refresh_token", refreshToken.Token, int(refreshToken.ExpiresAt.Unix()), "/", "localhost", false, false)
	c.Header("Authorization", "Bearer "+accessToken)
	c.JSON(200, gin.H{"message": "login successful", "user": user_res,"accessToken":accessToken})
	return
}
