package route

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/delivery/controller"
	"backend-starter-project/infrastructure/twilio"
	"backend-starter-project/service"

	"github.com/gin-gonic/gin"
)

func NewOTPRouter(group *gin.RouterGroup,env *bootstrap.Env) {
	twilioService := twilio.NewTwilioService(env.TwilloAccountSID, env.TwilloAuthToken,env.TwilloNumber)
	otpService := service.NewOTPService(twilioService)
	otpController := controller.NewOtpController(otpService)
	group.POST("/send", otpController.SendOtp)

}
