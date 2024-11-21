package service

import (
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	Infrastructure "backend-starter-project/infrastructure/cloudinary"
	"mime/multipart"
)

type defaultImageService struct {
	image_repo interfaces.DefaultImageRepository
	cld 	  Infrastructure.CloudinaryService
}

func NewDefaultImageService(image_repo interfaces.DefaultImageRepository,cld Infrastructure.CloudinaryService) interfaces.DefaultImageService {
	return &defaultImageService{image_repo: image_repo,cld: cld}
}

func (s *defaultImageService) CreateDefaultImage(file *multipart.FileHeader) error {
	url,publicId,err:=s.cld.UploadImage(file)
	if err!=nil{
		return err
	}
	err=s.image_repo.CreateDefaultImage(url,publicId)
	if err!=nil{
		return err
	}
	return nil
}

func (s *defaultImageService) DeleteDefaultImage(id string) error {
	image,err:=s.image_repo.GetDefaultImage(id)
	if err!=nil{
		return err
	}
	err=s.cld.DeleteImage(image.PublicID)
	if err!=nil{
		return err
	}
	return s.image_repo.DeleteDefaultImage(id)
}

func (s *defaultImageService) GetDefaultImage(id string) (*entities.DefaultImage, error) {
	return s.image_repo.GetDefaultImage(id)
}

func (s *defaultImageService) GetDefaultImages() ([]*entities.DefaultImage, error) {
	return s.image_repo.GetDefaultImages()
}