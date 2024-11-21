package Infrastructure

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type CloudinaryService struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryService(cloudName, apiKey, apiSecret string) *CloudinaryService {
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		panic(err)
	}
	return &CloudinaryService{cld: cld}
}
func (c *CloudinaryService) UploadImage(file *multipart.FileHeader) (string, string, error) {
	image, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer image.Close()
	res, err := c.cld.Upload.Upload(context.Background(), image, uploader.UploadParams{})
	if err != nil {
		return "", "", err
	}
	return res.SecureURL, res.PublicID, nil
}

func (c *CloudinaryService) DeleteImage(publicId string) error {
	_, err := c.cld.Upload.Destroy(context.Background(), uploader.DestroyParams{PublicID: publicId})
	if err != nil {
		return err
	}
	return nil
}
