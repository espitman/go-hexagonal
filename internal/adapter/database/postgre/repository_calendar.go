package postgre

import (
	"context"
	"github.com/espitman/go-hexagonal/internal/core/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type CalendarRepository struct {
	db *gorm.DB
}

/**
 * CalendarRepository implements port.CalendarRepository interface
 */

func NewCalendarRepository(db *gorm.DB) *CalendarRepository {
	return &CalendarRepository{
		db: db,
	}
}

func (r *CalendarRepository) Crete(ctx context.Context, calendar []domain.Calendar) ([]*domain.Calendar, error) {
	days := calendarDomainsToCalendarModelsMapper(calendar)
	columns := []string{"type", "is_peak", "is_custom_holiday"}
	result := r.db.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "date"}},
			DoUpdates: clause.AssignmentColumns(columns),
		},
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "jalali"}},
			DoUpdates: clause.AssignmentColumns(columns),
		},
	).Create(&days)
	if result.Error != nil {
		return nil, result.Error
	}
	response := calendarModelsToCalendarDomainsPointerMapper(days)
	return response, nil
}

func (r *CalendarRepository) Get(ctx context.Context, startDate *time.Time, endDate *time.Time, limit int, offset int) ([]*domain.Calendar, error) {
	var result []CalendarModel
	q := r.db.Limit(limit).Offset(offset)
	if startDate != nil {
		q.Where("date >= ?", startDate)
	}
	if endDate != nil {
		q.Where("date <= ?", endDate)
	}
	q.Order("date")
	q.Find(&result)
	response := calendarModelsToCalendarDomainsPointerMapper(result)
	return response, nil
}
