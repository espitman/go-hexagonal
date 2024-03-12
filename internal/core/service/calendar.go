package service

import (
	"context"
	"github.com/espitman/go-hexagonal/internal/core/domain"
	"github.com/espitman/go-hexagonal/internal/core/port"
	"time"
)

/**
 * CalendarService implements port.CalendarService interface
 */

type CalendarService struct {
	CalendarRepository port.CalendarRepository
}

func NewCalendarService(calendarRepository port.CalendarRepository) *CalendarService {
	return &CalendarService{
		CalendarRepository: calendarRepository,
	}
}

func (s *CalendarService) Crete(ctx context.Context, calendar []domain.Calendar) ([]*domain.Calendar, error) {
	return s.CalendarRepository.Crete(ctx, calendar)
}

func (s *CalendarService) Get(ctx context.Context, startDate *time.Time, endDate *time.Time, limit int, offset int) ([]*domain.Calendar, error) {
	return s.CalendarRepository.Get(ctx, startDate, endDate, limit, offset)
}
