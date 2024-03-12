package port

import (
	"context"
	"github.com/espitman/go-hexagonal/internal/core/domain"
	"time"
)

/**
 * CalendarService implemented by service.CalendarService interface
 */

type CalendarService interface {
	Crete(ctx context.Context, calendar []domain.Calendar) ([]*domain.Calendar, error)
	Get(ctx context.Context, startDate *time.Time, endDate *time.Time, limit int, offset int) ([]*domain.Calendar, error)
}
