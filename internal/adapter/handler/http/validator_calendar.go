package http

import "time"

func (h *CalendarHandler) createValidator(reqDto *calendarCreateRequestDto) error {
	if err := h.validate.Struct(reqDto); err != nil {
		return err
	}
	return nil
}

func (h *CalendarHandler) getValidator(limit int, offset int, startDate *time.Time, endDate *time.Time) error {
	return nil
}
