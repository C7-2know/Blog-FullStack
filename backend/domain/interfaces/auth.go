package interfaces

import (
	"backend-starter-project/domain/dtos"
	"backend-starter-project/domain/entities"
)

type AuthenticationService interface {
	RegisterUser(user *dtos.RegisterUserDto) (*entities.User, error)
	Login(emailOrUsername, password string) (*entities.RefreshToken,string,*entities.User, error)
	Logout(userId string) error
}

type PasswordResetService interface {
    RequestPasswordReset(email string) error
    ResetPassword(token, newPassword string) error
}