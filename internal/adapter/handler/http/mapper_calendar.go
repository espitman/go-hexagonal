package http

import (
	"github.com/espitman/go-hexagonal/internal/core/domain"
	"time"
)

func calendarCreateRequestDtoToCalendarDomainsMapper(dto calendarCreateRequestDto) []domain.Calendar {
	resp := make([]domain.Calendar, len(dto.Days))
	for i, d := range dto.Days {
		date, _ := time.Parse("2006-01-02", d.Date)
		resp[i] = domain.Calendar{
			Date:            date,
			Jalali:          d.Jalali,
			Type:            d.Type,
			IsPeak:          d.IsPeak,
			IsCustomHoliday: d.IsCustomHoliday,
		}
	}
	return resp
}

func calendarDomainsToCalendarGetResponseDtoMapper(ds []*domain.Calendar) calendarGetResponseDto {
	var resp calendarGetResponseDto
	days := make([]calendarDayDto, len(ds))
	for i, d := range ds {
		days[i] = calendarDayDto{
			Date:            d.Date.Format("2006-01-02"),
			Jalali:          d.Jalali,
			Type:            d.Type,
			IsPeak:          d.IsPeak,
			IsCustomHoliday: d.IsCustomHoliday,
		}
	}
	resp.Payload.Days = days
	return resp
}
