package media

import (
	"database/sql/driver"
	"image"
	"io"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

// Media is an interface including methods that needs for a media library storage
type Media interface {
	Scan(value interface{}) error
	Value() (driver.Value, error)

	GetUrlTemplate(*Option) string
	GetUrl(option *Option, scope *gorm.DB, field reflect.StructField, templater UrlTemplater) string

	GetFileHeader() FileHeader
	GetFileName() string

	GetSizes() map[string]*Size
	NeedCrop() bool
	Cropped(values ...bool) bool
	GetCropOption(name string) *image.Rectangle

	Store(url string, option *Option, reader io.Reader) error
	Retrieve(url string) (FileInterface, error)

	IsImage() bool

	URL(style ...string) string
	Ext() string
	String() string
}

type FileInterface interface {
	io.ReadSeeker
	io.Closer
}

type Size struct {
	Width   int
	Height  int
	Padding bool
}

type UrlTemplater interface {
	GetUrlTemplate(*Option) string
}

type Option map[string]string

func (option Option) Get(key string) string {
	return option[strings.ToUpper(key)]
}

func (option Option) Set(key string, val string) {
	option[strings.ToUpper(key)] = val
}
