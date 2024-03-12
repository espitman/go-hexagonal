package postgre

import "github.com/espitman/go-hexagonal/internal/core/domain"

func calendarDomainToCalendarModel(d domain.Calendar) CalendarModel {
	return CalendarModel{
		Date:            d.Date,
		Jalali:          d.Jalali,
		Type:            d.Type,
		IsPeak:          d.IsPeak,
		IsCustomHoliday: d.IsCustomHoliday,
	}
}

func calendarDomainsToCalendarModelsMapper(ds []domain.Calendar) []CalendarModel {
	resp := make([]CalendarModel, len(ds))
	for i, d := range ds {
		resp[i] = calendarDomainToCalendarModel(d)
	}
	return resp
}

func calendarModelToCalendarDomainMapper(c CalendarModel) domain.Calendar {
	return domain.Calendar{
		Date:            c.Date,
		Jalali:          c.Jalali,
		Type:            c.Type,
		IsPeak:          c.IsPeak,
		IsCustomHoliday: c.IsCustomHoliday,
	}
}

func calendarModelsToCalendarDomainsPointerMapper(cs []CalendarModel) []*domain.Calendar {
	resp := make([]*domain.Calendar, len(cs))
	for i, c := range cs {
		d := calendarModelToCalendarDomainMapper(c)
		resp[i] = &d
	}
	return resp
}
