package interfaces

import "backend-starter-project/domain/entities"

type ProfileRepository interface {
	GetProfiles() ([]*entities.Profile, error)
	GetUserProfile(userId string) (*entities.Profile, error)
	UpdateUserProfile(profile *entities.Profile) (*entities.Profile, error)
	CreateUserProfile(profile *entities.Profile) (*entities.Profile, error)
	DeleteUserProfile(user_id string) error
}

type ProfileService interface {
	GetProfiles() ([]*entities.Profile, error)
	GetUserProfile(userId string) (*entities.Profile, error)
	UpdateUserProfile(profile *entities.Profile) (*entities.Profile, error)
	CreateUserProfile(profile *entities.Profile) (*entities.Profile, error)
	DeleteUserProfile(user_id string) error
}
