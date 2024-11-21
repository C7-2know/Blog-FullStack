package controller

import (
	"backend-starter-project/domain/interfaces"
	"fmt"

	"github.com/gin-gonic/gin"
)

type OTPController struct {
	otpService interfaces.OTPService
}

func NewOtpController(otpService interfaces.OTPService) *OTPController {
	return &OTPController{otpService: otpService}
}

func (controller *OTPController) SendOtp(c *gin.Context) {
	var request struct {
		Email       string `json:"email" binding:"required"`
		PhoneNumber string `json:"phone_number" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"bind error": err.Error()})
		return
	}
	userId := "507f1f77bcf86cd799439011"
	otp, err := controller.otpService.GenerateOTP(request.Email, userId)
	if err != nil {
		c.JSON(500, gin.H{"generate error": err.Error()})
		return
	}
	err = controller.otpService.SendOTP(otp.Code, request.PhoneNumber, request.Email)
	fmt.Println("phone number", request.PhoneNumber)
	if err != nil {
		c.JSON(500, gin.H{"sending error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "OTP sent successfully"})
}
