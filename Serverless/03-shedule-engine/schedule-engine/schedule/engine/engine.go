package engine

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// Service interface
type Service interface {
	Schedule(ctx context.Context, schedule *Schedule) (Periods, error)
}

// Engine -
type Engine struct {
	logger log.Logger
}

// NewService is the func that initialize Engine service
func NewService(logger log.Logger) Service {
	engine := &Engine{logger: logger}
	return engine
}

// Day represents a day of the week
type Day int

// Days of the week
const (
	monday    = 1
	tuesday   = 2
	wednesday = 3
	thursday  = 4
	friday    = 5
	saturday  = 6
	sunday    = 0
)

// Periodicity of a schedule
type Periodicity int

// List of periodicities
const (
	everyWeek          = 0
	twoWeeksOnOddWeek  = 1
	twoWeeksOnEvenWeek = 2
)

// Schedule struct
type Schedule struct {
	StartDate  time.Time
	EndDate    time.Time
	Settings   []Setting
	Exceptions []time.Time
}

// Setting struct
type Setting struct {
	WeekDay     Day
	DaySettings []DaySetting
}

// DaySetting struct
type DaySetting struct {
	Periodicity Periodicity
	Time        Period
}

// AddDaySetting func
func (s *Setting) AddDaySetting(start time.Time, end time.Time, periodicity Periodicity) {
	s.DaySettings = append(s.DaySettings, DaySetting{
		Periodicity: periodicity,
		Time: Period{
			Start: start,
			End:   end,
		},
	})
}

// Period a range of two dates
type Period struct {
	Start time.Time
	End   time.Time
}

// Periods array of Period
type Periods struct {
	list []Period
}

// AddPeriod -
func (p *Periods) AddPeriod(start time.Time, end time.Time) {
	p.list = append(p.list, Period{Start: start, End: end})
}

// Index candidate day for the engine
type Index struct {
	date time.Time
}

// NewIndex builds index
func NewIndex(startDate time.Time) *Index {
	return &Index{date: startDate}
}

// NextDay add one day to index date
func (i *Index) NextDay() {
	i.date = i.date.AddDate(0, 0, 1)
}

// IsWeekDaySatisfied checks if week days math
func (i *Index) IsWeekDaySatisfied(s *Setting) bool {
	return s.WeekDay == Day(i.date.Weekday())
}

// IsPeriodicitySatisfied checks periodicity
func (i *Index) IsPeriodicitySatisfied(ds *DaySetting) bool {
	_, week := i.date.ISOWeek()
	switch ds.Periodicity {
	case everyWeek:
		return true
	case twoWeeksOnOddWeek:
		return week%2 == 0
	case twoWeeksOnEvenWeek:
		return week%2 != 0
	default:
		panic("unsuported periodicity")
	}
}

// IsExceptionSatisfied checks if candidate is in exception
func (i *Index) IsExceptionSatisfied(exceptions *[]time.Time) bool {
	for _, e := range *exceptions {
		if e.Day() == i.date.Day() && e.Month() == i.date.Month() && e.Year() == i.date.Year() {
			return true
		}
	}
	return false
}

// Schedule performs the schedule work
func (e *Engine) Schedule(ctx context.Context, schedule *Schedule) (Periods, error) {
	index := NewIndex(schedule.StartDate)
	periods := Periods{}

	// Loop for every day until end date.
	// On everyday evaluate certain things to see if the date is a eligible day.
	// Index will be the candidate date
	for index.date.Before(schedule.EndDate) || index.date.Equal(schedule.EndDate) {

		// Evaluate if the candidate is on exception list
		// If candidate date is on exception list next day will be evaluated
		if !index.IsExceptionSatisfied(&schedule.Exceptions) {

			// Loop through all settings.
			for _, setting := range schedule.Settings {

				// Check if candidate week day is the same week day as the setting that is evaluation.
				if index.IsWeekDaySatisfied(&setting) {
					for _, t := range setting.DaySettings {

						if index.IsPeriodicitySatisfied(&t) {
							start := time.Date(index.date.Year(), index.date.Month(), index.date.Day(), t.Time.Start.Hour(), t.Time.Start.Minute(), t.Time.Start.Second(), t.Time.Start.Nanosecond(), index.date.Location())
							end := time.Date(index.date.Year(), index.date.Month(), index.date.Day(), t.Time.End.Hour(), t.Time.End.Minute(), t.Time.End.Second(), t.Time.End.Nanosecond(), index.date.Location())
							periods.AddPeriod(start, end)
						}
					}
				}
			}
		}
		index.NextDay()
	}
	return periods, nil
}
