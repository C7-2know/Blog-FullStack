package interfaces


import "backend-starter-project/domain/entities"

type OTPService interface {
	GenerateOTP(email,userId string) (*entities.OTP, error)
	VerifyOTP(email, otp string) error
	SendOTP(userId,phoneNumber,email string) error
}

type TwilioService interface{
	SendOTP(phoneNumber, otp string) error
}


