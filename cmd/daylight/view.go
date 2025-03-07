package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	"github.com/jbreckmckye/daylight/internal"
)

const (
	width = 76

	brightBlue  = lipgloss.Color("#005FFF")
	brightGreen = lipgloss.Color("#00AF00")
	goldYellow  = lipgloss.Color("#FDC400")
	offWhite    = lipgloss.Color("#FFFDF5")
	pink        = lipgloss.Color("#A550DF")
	dimGrey     = lipgloss.Color("#353533")
)

var (
	titleBarStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(goldYellow).
		BorderBottom(true).
		MarginBottom(1).
		Foreground(goldYellow).
		Width(width)
)

func render(vm internal.TodayViewModel) string {
	doc := strings.Builder{}

	doc.WriteString(todayTitle())
	doc.WriteString(today(vm))

	doc.WriteString(lengthTitle())
	doc.WriteString(dayLength(vm))
	doc.WriteString(dayBar(vm))

	doc.WriteString(nextDaysTitle())
	doc.WriteString(projection())

	doc.WriteString(about())
	doc.WriteString(statusBar(vm))
	doc.WriteString(linkString())

	return doc.String()
}

func todayTitle() string {
	return titleBarStyle.Render("Today's daylight") + "\n"
}

func today(vm internal.TodayViewModel) string {
	graphic := lipgloss.NewStyle().
		Width(5)

	col := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(17).
		Padding(1, 0).
		Height(5)

	wrap := lipgloss.NewStyle().
		Padding(0, 2)

	rises := drawPixels(sunriseGradient)
	sets := drawPixels(sunsetGradient)
	noon := drawPixels(noonGradient)

	contents := lipgloss.JoinHorizontal(lipgloss.Top,
		graphic.Render(rises),
		col.Render("Rises\n"+vm.Rise),
		graphic.Render(noon),
		col.Render("Noon\n"+vm.Noon),
		graphic.Render(sets),
		col.Render("Sets\n"+vm.Sets),
	)

	return wrap.Render(contents) + "\n"
}

func lengthTitle() string {
	return titleBarStyle.Render("Day length") + "\n"
}

func dayLength(vm internal.TodayViewModel) string {
	col := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(35).
		Height(2)

	summary := fmt.Sprintf("Daylight for:\n%s", vm.Len)
	yesterday := fmt.Sprintf("versus yesterday:\n%s", vm.Diff)

	return lipgloss.JoinHorizontal(lipgloss.Top,
		col.Render(summary),
		col.Render(yesterday),
	) + "\n"
}

func dayBar(vm internal.TodayViewModel) string {
	barWidth := 72

	night := lipgloss.NewStyle().Background(dimGrey).Foreground(dimGrey)
	day := lipgloss.NewStyle().Background(brightBlue).Foreground(brightBlue)
	bar := lipgloss.NewStyle().Padding(1)

	if vm.DayEndRatio == 0 {
		// Polar night
		return bar.Render(night.Width(barWidth).Render(" ")) + "\n"
	} else if vm.DayEndRatio == 1 {
		// Polar day
		return bar.Render(day.Width(barWidth).Render(" ")) + "\n"
	}

	dayStart := int(math.Round(vm.DayStartRatio * float64(barWidth)))
	dayEnd := int(math.Round(vm.DayEndRatio * float64(barWidth)))

	text := strings.Builder{}
	for i := 0; i <= barWidth; i++ {
		if i == dayStart {
			// Mark sunrise with "R"
			text.WriteString(day.Render("R"))
		} else if i == dayEnd {
			// Mark sunset with "S"
			text.WriteString(day.Render("S"))
		} else if i > dayStart && i < dayEnd {
			// Fill day with spaces
			text.WriteString(day.Render("-"))
		} else {
			// Fill night with dashes
			text.WriteString(night.Render("."))
		}
	}

	return bar.Render(text.String()) + "\n"
}

func nextDaysTitle() string {
	return titleBarStyle.Render("Ten day projection") + "\n"
}

func statusBar(viewmodel internal.TodayViewModel) string {
	spans := lipgloss.NewStyle().
		Foreground(offWhite).
		Padding(0, 1)

	w := lipgloss.Width

	blueTag := spans.Background(brightBlue)
	goldTag := spans.Background(goldYellow)
	greenTag := spans.Background(brightGreen)
	greyTag := spans.Background(dimGrey)

	sLocation := "LOCATION"
	sLatlong := fmt.Sprintf("Latitude %s, Longitude %s", viewmodel.Lat, viewmodel.Lng)
	sIPAddress := "IP ADDRESS"
	sIPData := viewmodel.IP

	stretch := width - w(sLocation) - w(sIPAddress) - w(sIPData) - 6 // 6 padding

	return lipgloss.JoinHorizontal(lipgloss.Top,
		greenTag.Render(sLocation),
		greyTag.Width(stretch).Render(sLatlong),
		blueTag.Render(sIPAddress),
		goldTag.Render(sIPData),
	) + "\n"
}

func projection() string {
	rows := [][]string{
		{"Fri 7 March", "afsf", "afsf", "afsfsd"},
		{"Sat 8 March", "afsfs", "afsf", "dfdf"},
		{"Sun 9 March", "fdfd", "afsf", "fdfd"},
		{"Mon 10 March", "Здравствуйте", "afsf", "Привет"},
		{"Tue 11 March", "Hola", "afsf", "¿Qué tal?"},
		{"Wed 12 March", "Hola", "afsf", "¿Qué tal?"},
		{"Thu 13 March", "Hola", "afsf", "¿Qué tal?"},
		{"Fri 14 March", "Hola", "afsf", "¿Qué tal?"},
		{"Sat 15 March", "Hola", "afsf", "¿Qué tal?"},
		{"Sun 16 March", "Hola", "afsf", "¿Qué tal?"},
	}

	var (
		headerStyle  = lipgloss.NewStyle().Foreground(brightBlue).Bold(true).Align(lipgloss.Center)
		cellStyle    = lipgloss.NewStyle().Padding(0, 1).Width(14)
		oddRowStyle  = cellStyle.Foreground(pink)
		evenRowStyle = cellStyle.Foreground(offWhite)
		wrap         = lipgloss.NewStyle().Width(width).Align(lipgloss.Center).Padding(0, 0, 1)
	)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(brightBlue)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		}).
		Headers("DATE", "SUNRISE", "SUNSET", "LENGTH").
		Rows(rows...)

	return wrap.Render(t.Render()) + "\n"
}

func about() string {
	return titleBarStyle.Render("Your stats") + "\n"
}

func linkString() string {
	style := lipgloss.NewStyle().
		Width(width).
		Foreground(pink).
		Padding(1, 0, 0)

	return style.Render("https://github.com/jbreckmckye/daylight") + "\n"
}

func drawPixels(px [][]uint) string {
	builder := strings.Builder{}
	style := lipgloss.NewStyle()

	for _, row := range px {
		for j, cell := range row {
			colour := lipgloss.Color(ABGRtoHex(cell))
			builder.WriteString(style.Background(colour).Render("  "))
			if j == len(row)-1 {
				builder.WriteString("\n")
			}
		}
	}

	return builder.String()
}
