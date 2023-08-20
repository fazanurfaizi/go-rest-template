package converter

import (
	"reflect"
	"regexp"
)

var (
	ColumnNameRegexp = regexp.MustCompile(`(?m)column:(\w{1,}).*`)
	ParamNameRegexp  = regexp.MustCompile(`(?m)param:(\w{1,}).*`)
)

func GetColumnNameForField(field reflect.StructField) string {
	fieldTag := field.Tag.Get("gorm")
	res := ColumnNameRegexp.FindStringSubmatch(fieldTag)
	if len(res) == 2 {
		return res[1]
	}
	return field.Name
}
