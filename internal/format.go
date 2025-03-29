package internal

import (
	"fmt"
	"time"
)

func LocalisedTime(t time.Time, tz *time.Location, hf string) string {
	if hf == "12h" {
		return t.In(tz).Format("03:04 PM")
	} else if hf == "24h" {
		return t.In(tz).Format("15:04")
	} else {
		return t.In(tz).Format("15:04 PM")
	}
}

func FormatDayLength(s SunTimes) string {
	if s.PolarDay {
		return "all day (polar sun)"
	}

	if s.PolarNight {
		return "no time (polar night)"
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
		return "the same"
	}

	prefix := "+"
	if direction == -1 {
		prefix = "-"
	}

	diff := (today.Length - yesterday.Length).Abs()
	h, m, s := durationHMS(diff)
	mins := m + (h * 60)

	return fmt.Sprintf("%s%dm %ds", prefix, mins, s)
}

func FormatNoon(s SunTimes, tz *time.Location, hourFormat string) string {
	if s.PolarDay {
		return "n/a"
	}

	if s.PolarNight {
		return "n/a"
	}

	noon := s.ApproximateNoon()
	return LocalisedTime(noon, tz, hourFormat)
}

func FormatRises(s SunTimes, tz *time.Location, hourFormat string) string {
	if (s.Rises == time.Time{}) {
		return "n/a"
	}
	return LocalisedTime(s.Rises, tz, hourFormat)
}

func FormatSets(s SunTimes, tz *time.Location, hourFormat string) string {
	if (s.Sets == time.Time{}) {
		return "n/a"
	}
	return LocalisedTime(s.Sets, tz, hourFormat)
}

func FormatDate(t time.Time) string {
	return t.Format("Mon Jan 02")
}

func FormatDayRatio(s SunTimes, tz *time.Location) (start float64, end float64) {
	if s.PolarDay {
		return 0, 1
	}

	if s.PolarNight {
		return 0, 0
	}

	dayMins := float64(24 * 60)

	riseH, riseM, _ := s.Rises.In(tz).Clock()
	rises := float64((riseH * 60) + riseM)

	setsH, setsM, _ := s.Sets.In(tz).Clock()
	sets := float64((setsH * 60) + setsM)

	return rises / dayMins, sets / dayMins
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
