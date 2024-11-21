package interfaces

import (
	"backend-starter-project/domain/entities"
	"mime/multipart"
)

type DefaultImageRepository interface {
	CreateDefaultImage(url,publicId string) error
	DeleteDefaultImage(id string) error
	GetDefaultImage(id string) (*entities.DefaultImage, error)
	GetDefaultImages() ([]*entities.DefaultImage, error)
}
type DefaultImageService interface {
	CreateDefaultImage(file *multipart.FileHeader) error
	DeleteDefaultImage(id string) error
	GetDefaultImage(id string) (*entities.DefaultImage, error)
	GetDefaultImages() ([]*entities.DefaultImage, error)
}