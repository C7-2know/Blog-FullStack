package service

import (
	"backend-starter-project/domain/dtos"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"fmt"
)

type userService struct {
	userRepository    interfaces.UserRepository
	profileRepository interfaces.ProfileRepository
}

func NewUserService(userRepository interfaces.UserRepository, profile_repo interfaces.ProfileRepository) interfaces.UserService {
	return &userService{
		userRepository:    userRepository,
		profileRepository: profile_repo,
	}
}

func (service *userService) CreateUser(user *dtos.RegisterUserDto) (*entities.User, error) {

	created, err := service.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	profile := entities.Profile{
		UserID:  created.ID,
		ProfilePicture: "",
		ContactInfo: entities.ContactInfo{
			Email:       user.Email,
			PhoneNumber: "",

		},
	}
	fmt.Println("profile before created", profile)
	_, err = service.profileRepository.CreateUserProfile(&profile)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (service *userService) FindUserByEmail(email string) (*entities.User, error) {
	return service.userRepository.FindUserByEmail(email)
}

func (service *userService) FindUserById(userId string) (*entities.User, error) {
	return service.userRepository.FindUserById(userId)
}

func (service *userService) UpdateUser(user *entities.User) (*entities.User, error) {
	return service.userRepository.UpdateUser(user)
}

func (service *userService) DeleteUser(userId string) error {
	return service.userRepository.DeleteUser(userId)
}

func (service *userService) PromoteUserToAdmin(userId string) error {
	user, err := service.userRepository.FindUserById(userId)
	if err != nil {
		return err
	}

	user.Role = "admin"
	_, err = service.userRepository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (service *userService) DemoteUserToRegular(userId string) error {
	user, err := service.userRepository.FindUserById(userId)
	if err != nil {
		return err
	}

	user.Role = "regular"
	_, err = service.userRepository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
