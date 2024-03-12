package domain

import (
	"time"
)

type Calendar struct {
	Date            time.Time
	Jalali          string
	Type            string
	IsPeak          bool
	IsCustomHoliday bool
}
