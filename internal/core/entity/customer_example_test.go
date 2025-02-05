package entity_test

import (
	"testing"
	"time"
)

type CustomerJobInfo string

type Customer struct {
	ID           int64
	Name         string
	Email        string
	Phone        string
	FuriganaName string
	Birth        time.Time
	Job          CustomerJobInfo
}

func (c *Customer) Age(baseDate time.Time) int {
	age := baseDate.Year() - c.Birth.Year()
	if c.Birth.YearDay() > baseDate.YearDay() {
		age--
	}
	return age
}

func (c *Customer) IsOver18Age(baseDate time.Time) bool {
	return c.Age(baseDate) >= 18
}

func (c *Customer) IsMinor(baseDate time.Time) bool {
	return !c.IsOver18Age(baseDate)
}

func (c *Customer) IsAdult(baseDate time.Time) bool {
	return c.IsOver18Age(baseDate)
}

// ---------------------------------------------------------

func TestCustomerAge(t *testing.T) {
	tests := []struct {
		name     string
		birth    time.Time
		baseDate time.Time
		correct  int
	}{
		{
			name:     "adult",
			birth:    time.Date(2000, 2, 1, 0, 0, 0, 0, time.UTC),
			baseDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			correct:  24,
		},
		{
			name:     "minor",
			birth:    time.Date(2010, 12, 31, 0, 0, 0, 0, time.UTC),
			baseDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			correct:  14,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			customer := &Customer{
				Birth: test.birth,
			}
			age := customer.Age(test.baseDate)
			if age != test.correct {
				t.Errorf("Expected age %d, got %d", test.correct, age)
			}
		})
	}
}
