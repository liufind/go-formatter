package formatter

import (
	"time"
)

func setISO8601(t time.Time) string {
	return t.Format(time.RFC3339)
}
