package util

import (
	"errors"
	"time"
)

func StringToDate(dateString string) (*time.Time, error) {
	if dateString == "" {
		return nil, nil
	}
	layout := "2006-01-02"
	date, err := time.Parse(layout, dateString)
	if err != nil {
		return nil, errors.New("invalid date format")
	}
	return &date, nil
}
