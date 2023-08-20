package media

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/fazanurfaizi/go-rest-template/pkg/converter"
	"gorm.io/gorm"
)

var (
	MediaLibraryURL = ""
)

func cropField(field reflect.StructField, scope *gorm.DB) (cropped bool) {
	columnName := converter.GetColumnNameForField(field)
	fieldValue := reflect.ValueOf(columnName)

	if fieldValue.CanAddr() {
		// Handle scanner
		if media, ok := fieldValue.Addr().Interface().(Media); ok && !media.Cropped() {
			option := ParseMediaTagOption(field.Tag.Get("media_library"))
			if MediaLibraryURL != "" {
				option.Set("url", MediaLibraryURL)
			}
			if media.GetFileHeader() != nil || media.NeedCrop() {
				var mediaFile FileInterface
				var err error
				if fileHeader := media.GetFileHeader(); fileHeader != nil {
					mediaFile, err = media.GetFileHeader().Open()
				} else {
					mediaFile, err = media.Retrieve(media.URL("original"))
				}

				if err != nil {
					scope.AddError(err)
					return false
				}

				media.Cropped(true)

				if url := media.GetUrl(option, scope, field, media); url == "" {
					scope.AddError(errors.New("invalid URL"))
				} else {
					result, _ := json.Marshal(map[string]string{"Url": url})
					media.Scan(string(result))
				}

				if mediaFile != nil {
					defer mediaFile.Close()
					var handled = false
					for _, handler := range mediaHandlers {
						if handler.CouldHandle(media) {
							mediaFile.Seek(0, 0)
							if scope.AddError(handler.Handle(media, mediaFile, option)) == nil {
								handled = true
							}
						}
					}

					if !handled {
						scope.AddError(media.Store(media.URL(), option, mediaFile))
					}
				}
				return true
			}
		}
	}
	return false
}

// func saveAndCropImage(isCreate bool) func(scope *gorm.DB) {
// 	return func(scope *gorm.DB) {
// 		if scope.Error == nil {
// 			var updatedColumns = map[string]interface{}{}

// 			// Handle SerializableMeta
// 			if value, ok := scope.Statement.ReflectValue.(se)
// 		}
// 	}
// }
