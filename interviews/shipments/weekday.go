package main

import "fmt"

// weekDay represents a day of the work week
type weekDay byte

// weekday constants
const (
	Monday    weekDay = 'M'
	Tuesday           = 'T'
	Wednesday         = 'W'
	Thursday          = 'R'
	Friday            = 'F'
	EOD               = 'Z'
)

// parseWeekDay parses a weekday from it's string representation (single char)
func parseWeekDay(day string) (weekDay, error) {
	weekDay := weekDay(day[0])
	switch weekDay {
	case Monday, Tuesday, Wednesday, Thursday, Friday:
		return weekDay, nil
	default:
		return Monday, fmt.Errorf("'%s' is an invalid weekday value. Must be 'M', 'T', 'W', 'R' or 'F'", string(day))
	}
}

// nextday returns the chronologically following day or EOD if we have reached the week's end.
func (w weekDay) nextDay() weekDay {
	switch w {
	case 'M':
		return Tuesday
	case 'T':
		return Wednesday
	case 'W':
		return Thursday
	case 'R':
		return Friday
	}
	return EOD
}
