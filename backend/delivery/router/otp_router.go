package route

import (
	"backend-starter-project/delivery/controller"
	"backend-starter-project/infrastructure/twilio"
	"backend-starter-project/service"

	"github.com/gin-gonic/gin"
)

func NewOTPRouter(group *gin.RouterGroup) {
	twilioService := twilio.NewTwilioService("AC70e6cd6a8ea6f7b1052e0fa8fc89c0af", "215ede45175d11974d0e4e660d4c3335", "+251937380008")
	otpService := service.NewOTPService(twilioService)
	otpController := controller.NewOtpController(otpService)
	group.POST("/send", otpController.SendOtp)

}
