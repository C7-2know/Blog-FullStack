package service

import (
	"backend-starter-project/domain/dtos"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"errors"
)

type authService struct {
	userService     interfaces.UserService
	tokenRepo       interfaces.RefreshTokenRepository
	passwordService interfaces.PasswordService
	tokenService    interfaces.TokenService
}

func NewAuthService(userService interfaces.UserService, tokenRepo interfaces.RefreshTokenRepository,
	passService interfaces.PasswordService, tokenService interfaces.TokenService) interfaces.AuthenticationService {
	return &authService{
		userService:     userService,
		tokenRepo:       tokenRepo,
		tokenService:    tokenService,
		passwordService: passService,
	}
}

func (service *authService) RegisterUser(user *dtos.RegisterUserDto) (*entities.User, error) {
	hashedPass, err := service.passwordService.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPass
	newUser, err := service.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}
	
	return newUser, nil
}

func (service *authService) Login(emailOrUsername, password string) (*entities.RefreshToken, string,*entities.User, error) {
	user, err := service.userService.FindUserByEmail(emailOrUsername)
	if err != nil {
		return nil, "",nil, errors.New("User not found")
	}
	err = service.passwordService.ComparePassword(user.Password, password)
	if err != nil {
		return nil, "",nil, errors.New("Invalid password")
	}
	token, err := service.tokenService.GenerateAccessToken(user)
	if err != nil {
		return nil, "",nil, err
	}
	refresh_tok, err := service.tokenService.GenerateRefreshToken(user)
	if err != nil {
		return nil, "",nil, err
	}
	return refresh_tok, token,user, nil
}

func (service *authService) Logout(userId string) error {

	//delete the token from database
	err := service.tokenRepo.DeleteRefreshTokenByUserId(userId)
	if err != nil {
		return err
	}
	return nil

}
