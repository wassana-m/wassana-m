package appvalidator

import (
	"time"
)

const (
	ISO8601DateLayout = "2006-01-02T15:04:05"
)

func IsISO8601Date(field string) bool {
	if _, err := time.Parse(ISO8601DateLayout, field); err != nil {
		return false
	}

	return true
}
