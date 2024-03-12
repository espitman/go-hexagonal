package http

type calendarDayDto struct {
	Date            string `json:"date" validate:"required,IsDate"`
	Jalali          string `json:"jalali" validate:"required"`
	Type            string `json:"type" validate:"required,oneof=base weekend holiday"`
	IsPeak          bool   `json:"isPeak" validate:"boolean"`
	IsCustomHoliday bool   `json:"isCustomHoliday" validate:"boolean"`
}

type calendarCreateRequestDto struct {
	Days []calendarDayDto `json:"days" validate:"required,dive"`
}

type calendarGetResponseDto struct {
	ResponseDto
	Payload struct {
		Days []calendarDayDto `json:"days"`
	} `json:"payload"`
}
