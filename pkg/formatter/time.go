package formatter

import "time"

const (
	YYYYMMDD       = "2006-01-02"
	DDMMYYYY       = "02/01/2006"
	YYYYMMDDhhmmss = "2006-01-02 15:04:05"
)

// Format time layout
func FormatTime(value time.Time, format string) string {
	layout := "2006-01-02 15:04:05 -0700 MST"
	date, _ := time.Parse(layout, value.String())
	utc := date.UTC()
	return utc.Format(format)
}
