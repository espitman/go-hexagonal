package postgre

import (
	"gorm.io/gorm"
	"time"
)

type CalendarModel struct {
	gorm.Model
	Date            time.Time `gorm:"type:date;uniqueIndex"`
	Jalali          string    `gorm:"uniqueIndex"`
	Type            string
	IsPeak          bool `gorm:"default: false"`
	IsCustomHoliday bool `gorm:"default: false"`
}

func (CalendarModel) TableName() string {
	return "calendar"
}
