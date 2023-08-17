package storage

import (
	"context"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/fazanurfaizi/go-rest-template/pkg/config"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

type cloudinaryStorage struct {
	config *config.Config
	logger logger.Logger
}

func NewCloudinaryStorage(config *config.Config, logger logger.Logger) FileStorage {
	return &cloudinaryStorage{
		config: config,
		logger: logger,
	}
}

func (s cloudinaryStorage) Upload(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create cloudinary instance
	cld, err := cloudinary.NewFromParams(
		s.config.Cloudinary.CloudName,
		s.config.Cloudinary.APIKey,
		s.config.Cloudinary.APISecret,
	)
	if err != nil {
		s.logger.Panicln(err.Error())
	}

	// Upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{
		Folder: s.config.Cloudinary.UploadFolder,
	})
	if err != nil {
		s.logger.Panicln(err.Error())
		return "", err
	}

	return uploadParam.PublicID, nil
}

func (s cloudinaryStorage) GetFile(filename string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create cloudinary instance
	cld, err := cloudinary.NewFromParams(
		s.config.Cloudinary.CloudName,
		s.config.Cloudinary.APIKey,
		s.config.Cloudinary.APISecret,
	)
	if err != nil {
		s.logger.Panicln(err.Error())
		return "", err
	}

	// Get details about the image with PublicID.
	response, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: filename})
	if err != nil {
		s.logger.Panicln(err.Error())
	}

	return response.SecureURL, nil
}
