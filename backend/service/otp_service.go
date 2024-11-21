package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type otpService struct {
	store  map[string]entities.OTP
	twilio interfaces.TwilioService
}

func NewOTPService(twilio interfaces.TwilioService) interfaces.OTPService {
	return &otpService{store: make(map[string]entities.OTP),
		twilio: twilio}
}

func (service *otpService) GenerateOTP(email, Id string) (*entities.OTP, error) {
	code := generateRandomCode()
	userId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		fmt.Println("Error converting string to object id")
		return nil, err
	}
	otp := entities.OTP{
		ID:         userId,
		Code:       code,
		Email:      email,
		Expiration: time.Now().Add(time.Minute * 5),
	}
	service.store[email] = otp
	return &otp, nil
}

func (service *otpService) VerifyOTP(email, code string) error {
	stored_otp := service.store[email]
	if stored_otp.Code == code && stored_otp.Expiration.After(time.Now()) {
		delete(service.store, email)
		return nil
	}
	return errors.New("Invalid OTP")
}

func (service *otpService) SendOTP(code, phone, email string) error {
	fmt.Println("Sending OTP to ", phone)
	return service.twilio.SendOTP(phone, code)
}

func generateRandomCode() string {
	return "123466"
}
