package http

import (
	"github.com/espitman/go-hexagonal/internal/core/port"
	"github.com/espitman/go-hexagonal/internal/core/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CalendarHandler struct {
	validate        *validator.Validate
	calendarService port.CalendarService
}

func NewCalendarHandler(validate *validator.Validate, calendarService port.CalendarService) *CalendarHandler {
	return &CalendarHandler{
		validate:        validate,
		calendarService: calendarService,
	}
}

// Create
// @Summary Create
// @Description Create
// @Tags Calendar
// @Produce json
// @Param body body http.calendarCreateRequestDto true "body"
// @Success 200 {object} http.calendarGetResponseDto
// @Router /api/v1/calendar [post]
func (h *CalendarHandler) Create(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := fctx.Context()
	reqDto := new(calendarCreateRequestDto)
	if err := fctx.BodyParser(reqDto); err != nil {
		return fctx.BadRequest(err)
	}
	if err := h.createValidator(reqDto); err != nil {
		return fctx.BadRequest(err)
	}
	days := calendarCreateRequestDtoToCalendarDomainsMapper(*reqDto)
	createdDays, err := h.calendarService.Crete(ctx, days)
	resp := calendarDomainsToCalendarGetResponseDtoMapper(createdDays)
	if err != nil {
		return fctx.Conflict(err)
	}
	return fctx.ResponseOk(resp.Payload)
}

// Get
// @Summary Get
// @Description Get
// @Tags Calendar
// @Produce json
// @Param start_date query string false "Start date" format(YYYY-MM-DD)
// @Param end_date query string false "End date " format(YYYY-MM-DD)
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} http.calendarGetResponseDto
// @Router /api/v1/calendar [get]
func (h *CalendarHandler) Get(c *fiber.Ctx) error {
	var fctx = fiberCtx{c}
	ctx := fctx.Context()
	limit := fctx.QueryInt("limit", -1)
	offset := fctx.QueryInt("offset", 0)
	startDate, err := util.StringToDate(fctx.Query("start_date"))
	if err != nil {
		return fctx.BadRequest(err)
	}
	endDate, err := util.StringToDate(fctx.Query("end_date"))
	if err != nil {
		return fctx.BadRequest(err)
	}

	if err := h.getValidator(limit, offset, startDate, endDate); err != nil {
		return fctx.BadRequest(err)
	}

	days, err := h.calendarService.Get(ctx, startDate, endDate, limit, offset)
	if err != nil {
		return fctx.BadRequest(err)
	}
	resp := calendarDomainsToCalendarGetResponseDtoMapper(days)
	return fctx.ResponseOk(resp.Payload)
}
