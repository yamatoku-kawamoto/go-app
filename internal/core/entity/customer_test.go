package entity

import (
	"testing"
	"time"
)

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
