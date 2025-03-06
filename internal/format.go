package internal

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/term"

	templates "github.com/jbreckmckye/daylight/internal/templates"

	"github.com/fatih/color"
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

func FormatNoon(s SunTimes, tz *time.Location) string {
	if s.PolarDay {
		return "n/a"
	}

	if s.PolarNight {
		return "n/a"
	}

	noon := s.ApproximateNoon()
	return LocalisedTime(noon, tz)
}

func FormatRises(s SunTimes, tz *time.Location) string {
	if (s.Rises == time.Time{}) {
		return "n/a"
	}
	return LocalisedTime(s.Rises, tz)
}

func FormatSets(s SunTimes, tz *time.Location) string {
	if (s.Sets == time.Time{}) {
		return "n/a"
	}
	return LocalisedTime(s.Sets, tz)
}

func UsePrettyMode() bool {
	if !term.IsTerminal(0) {
		return false
	}

	width, _, err := term.GetSize(0)
	if err != nil {
		return false
	}

	if width < 80 {
		return false
	}

	return true
}

// Sunnify makes an input string... sunny :-D
func Sunnify(s string) string {
	ins := strings.Split(s, "\n")
	sunLines := strings.Split(templates.SunTxt, "\n")

	yellow := color.New(color.FgHiYellow, color.Bold)

	var output string
	for lineN, lineIn := range ins {
		lineOut := ""

		if lineN >= len(sunLines) {
			// "Picture" is complete, skip concatenations
			lineOut = lineIn

		} else {
			lineOut = padToLength(lineIn, 40)
			lineOut = lineOut + yellow.Sprint(sunLines[lineN])
		}

		output = output + lineOut + "\n"
	}

	return output
}

func padToLength(s string, to int) string {
	padding := to - len(s)
	if padding > 0 {
		return s + strings.Repeat(" ", padding)
	}
	return s
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
