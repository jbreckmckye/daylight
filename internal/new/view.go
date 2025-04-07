package new

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	width = 76

	brightBlue = lipgloss.AdaptiveColor{
		Light: "#003FDF",
		Dark:  "#005FFF",
	}
	brightGreen = lipgloss.Color("#00AF00")
	goldYellow  = lipgloss.AdaptiveColor{
		Light: "#CD8400",
		Dark:  "#FDC400",
	}
	offWhite = lipgloss.Color("#FFFDF5")
	pink     = lipgloss.Color("#B150E6")
	dimGrey  = lipgloss.AdaptiveColor{
		Light: "#555553",
		Dark:  "#353533",
	}
)

var titleBarStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.DoubleBorder()).
	BorderForeground(goldYellow).
	BorderBottom(true).
	MarginBottom(1).
	Foreground(goldYellow).
	Width(width)

func render(sv SummaryView) string {
	doc := strings.Builder{}

	doc.WriteString("\n")
	doc.WriteString(todayTitle())
	doc.WriteString(today(sv))

	doc.WriteString(lengthTitle())
	doc.WriteString(dayLength(sv))
	doc.WriteString(dayBar(sv))

	doc.WriteString(nextDaysTitle())
	doc.WriteString(projection(sv))

	doc.WriteString(about())
	doc.WriteString(statusBar(sv))
	doc.WriteString(linkString())

	return doc.String()
}

func todayTitle() string {
	return titleBarStyle.Render("Today's daylight") + "\n"
}

func today(sv SummaryView) string {
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
		col.Render("Rises\n"+sv.Rise),
		graphic.Render(noon),
		col.Render("Noon:\n"+sv.Noon),
		graphic.Render(sets),
		col.Render("Sets:\n"+sv.Sets),
	)

	return wrap.Render(contents) + "\n"
}

func lengthTitle() string {
	return titleBarStyle.Render("Day length") + "\n"
}

func dayLength(sv SummaryView) string {
	col := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(35).
		Height(2)

	summary := fmt.Sprintf("Daylight for:\n%s", sv.Len)
	yesterday := fmt.Sprintf("versus yesterday:\n%s", sv.Diff)

	return lipgloss.JoinHorizontal(lipgloss.Top,
		col.Render(summary),
		col.Render(yesterday),
	) + "\n"
}

func dayBar(sv SummaryView) string {
	barWidth := 72

	night := lipgloss.NewStyle().Background(dimGrey).Foreground(dimGrey)
	day := lipgloss.NewStyle().Background(brightBlue).Foreground(brightBlue)
	bar := lipgloss.NewStyle().Padding(1)

	if sv.DayEndRatio == 0 {
		// Polar night
		return bar.Render(night.Render(strings.Repeat(".", barWidth))) + "\n"
	} else if sv.DayEndRatio == 1 {
		// Polar day
		return bar.Render(day.Render(strings.Repeat("-", barWidth))) + "\n"
	}

	dayStart := int(math.Round(sv.DayStartRatio * float64(barWidth)))
	dayEnd := int(math.Round(sv.DayEndRatio * float64(barWidth)))

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

func projection(sv SummaryView) string {
	rows := make([][]string, len(sv.Next10Days))

	for i, v := range sv.Next10Days {
		rows[i] = []string{v.Day, v.Rise, v.Sets, v.Length}
	}

	var (
		headerStyle  = lipgloss.NewStyle().Foreground(brightBlue).Bold(true).Align(lipgloss.Center)
		cellStyle    = lipgloss.NewStyle().Padding(0, 3)
		oddRowStyle  = cellStyle.Foreground(pink)
		evenRowStyle = cellStyle
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
		Headers("DATE", "SUNRISE", " SUNSET", "LENGTH").
		Rows(rows...)

	return wrap.Render(t.Render()) + "\n"
}

func about() string {
	return titleBarStyle.Render("Your stats") + "\n"
}

func statusBar(sv SummaryView) string {
	spans := lipgloss.NewStyle().
		Foreground(offWhite).
		Padding(0, 1)

	w := lipgloss.Width

	blueTag := spans.Background(brightBlue)
	goldTag := spans.Background(goldYellow)
	greenTag := spans.Background(brightGreen)
	greyTag := spans.Background(dimGrey)

	sLocation := "LOCATION"
	sLatlong := fmt.Sprintf("Latitude %s, Longitude %s", sv.Lat, sv.Lng)
	sIPAddress := "IP ADDRESS"
	sIPData := sv.IP

	stretch := width - w(sLocation) - w(sIPAddress) - w(sIPData) - 6 // 6 padding

	return lipgloss.JoinHorizontal(lipgloss.Top,
		greenTag.Render(sLocation),
		greyTag.Width(stretch).Render(sLatlong),
		blueTag.Render(sIPAddress),
		goldTag.Render(sIPData),
	) + "\n"
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
