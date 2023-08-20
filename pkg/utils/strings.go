package utils

import (
	"strings"

	"github.com/fazanurfaizi/go-rest-template/pkg/constants"
	"github.com/gosimple/slug"
)

// ParseTagOption parse tag options to hash
func ParseTagOption(str string) map[string]string {
	tags := strings.Split(str, ";")
	setting := map[string]string{}
	for _, tag := range tags {
		v := strings.Split(tag, ":")
		k := strings.TrimSpace(strings.ToUpper(v[0]))
		if len(v) == 2 {
			setting[k] = v[1]
		} else {
			setting[k] = k
		}
	}
	return setting
}

func ToParamString(str string) string {
	if constants.AsicsiiRegexp.MatchString(str) {
		return strings.Replace(str, " ", "_", -1)
	}

	return slug.Make(str)
}
