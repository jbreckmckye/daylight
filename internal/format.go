package internal

import (
	"fmt"
	"time"
)

func LocalisedTime(t time.Time, tz *time.Location) string {
	return t.In(tz).Format("15:04 PM")
}

func FormatDayLength(s SunTimes) string {
	if s.PolarDay {
		return "all day (polar sun)"
	}

	if s.PolarNight {
		return "none (polar night)"
	}

	h, m, _ := durationHMS(s.Length)

	return fmt.Sprintf("%d hrs, %d mins", h, m)
}

func FormatLengthDiff(today SunTimes, yesterday SunTimes) string {
	direction := 0
	if today.Length > yesterday.Length {
		direction = 1
	}
	if today.Length < yesterday.Length {
		direction = -1
	}

	if direction == 0 {
		return "the same as yesterday"
	}

	prefix := "+"
	if direction == -1 {
		prefix = "-"
	}

	diff := (today.Length - yesterday.Length).Abs()
	h, m, s := durationHMS(diff)
	mins := m + (h * 60)

	return fmt.Sprintf("%s%dm %ds vs yesterday", prefix, mins, s)
}

func FormatNoon(s SunTimes, tz *time.Location) string {
	if s.PolarDay {
		return "--"
	}

	if s.PolarNight {
		return "--"
	}

	noon := s.ApproximateNoon()
	return LocalisedTime(noon, tz)
}

func durationHMS(d time.Duration) (hours int64, minutes int64, seconds int64) {
	// iterative subtraction
	seconds = int64(d.Round(time.Second).Seconds())

	hours = seconds / 3600
	seconds = seconds - (hours * 3600)

	minutes = seconds / 60
	seconds = seconds - (minutes * 60)

	return hours, minutes, seconds
}
