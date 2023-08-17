package storage

import (
	"github.com/fazanurfaizi/go-rest-template/pkg/config"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"go.uber.org/fx"
)

type FileStorage interface {
	Upload(input interface{}) (string, error)
	GetFile(filename string) (string, error)
}

func NewFileStorage(config *config.Config, logger logger.Logger) FileStorage {
	return NewCloudinaryStorage(config, logger)
}

var Module = fx.Options(
	fx.Provide(NewFileStorage),
)
