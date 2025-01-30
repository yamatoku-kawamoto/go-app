package entity

import "time"

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
