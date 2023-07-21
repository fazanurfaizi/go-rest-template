package sanitize

import (
	"bytes"
	"encoding/json"

	"github.com/microcosm-cc/bluemonday"
)

var sanitizer *bluemonday.Policy

func init() {
	sanitizer = bluemonday.UGCPolicy()
}

// Sanitize JSON
func SanitizeJSON(s []byte) ([]byte, error) {
	d := json.NewDecoder(bytes.NewReader(s))
	d.UseNumber()
	var i interface{}
	err := d.Decode(&i)
	if err != nil {
		return nil, err
	}
	sanitize(i)
	return json.MarshalIndent(i, "", "    ")
}

func sanitize(data interface{}) {
	switch d := data.(type) {
	case map[string]interface{}:
		for i, v := range d {
			switch tv := v.(type) {
			case string:
				d[i] = sanitizer.Sanitize(tv)
			case map[string]interface{}:
				sanitize(tv)
			case []interface{}:
				sanitize(tv)
			case nil:
				delete(d, i)
			}
		}
	case []interface{}:
		if len(d) > 0 {
			switch d[0].(type) {
			case string:
				for i, v := range d {
					d[i] = sanitizer.Sanitize(v.(string))
				}
			case map[string]interface{}:
				for _, v := range d {
					sanitize(v)
				}
			case []interface{}:
				for _, v := range d {
					sanitize(v)
				}
			}
		}
	}
}
