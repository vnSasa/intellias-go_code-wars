package task4kyu

import (
	"strconv"
)

func FormatDuration(seconds int64) string {
	if seconds == 0 {
		return "now"
	}

	// describe the structure for convenient storage of time components
	type durationComponent struct {
		value int64
		unit  string
	}

	// Set values for time units in seconds
	const (
		second = 1
		minute = 60 * second
		hour   = 60 * minute
		day    = 24 * hour
		year   = 365 * day
	)

	// calculate time components
	var components []durationComponent

	years := seconds / year
	if years > 0 {
		components = append(components, durationComponent{years, "year"})
	}
	seconds -= years * year

	days := seconds / day
	if days > 0 {
		components = append(components, durationComponent{days, "day"})
	}
	seconds -= days * day

	hours := seconds / hour
	if hours > 0 {
		components = append(components, durationComponent{hours, "hour"})
	}
	seconds -= hours * hour

	minutes := seconds / minute
	if minutes > 0 {
		components = append(components, durationComponent{minutes, "minute"})
	}
	seconds -= minutes * minute

	if seconds > 0 {
		components = append(components, durationComponent{seconds, "second"})
	}

	// form an answer
	var result string
	for i, c := range components {
		if c.value == 1 {
			result += strconv.FormatInt(c.value, 10) + " " + c.unit
		} else {
			result += strconv.FormatInt(c.value, 10) + " " + c.unit + "s"
		}

		if i < len(components)-2 {
			result += ", "
		} else if i == len(components)-2 {
			result += " and "
		}
	}

	return result
}