package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/jbreckmckye/daylight/internal"
)

const (
	width = 100

	brightBlue  = lipgloss.Color("#005FFF")
	brightGreen = "#00aF00"
	goldYellow  = lipgloss.Color("#FDC400")
	sunYellow   = lipgloss.Color("#EDFF82")
	offWhite    = "#FFFDF5"
	pink        = lipgloss.Color("#A550DF")
	purple      = "#6124DF"
	midGrey     = lipgloss.Color("#525250")
	dimGrey     = lipgloss.Color("#353533")
)

var (
	titleBarStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#87CEEB")).
		BorderBottom(true).
		Foreground(lipgloss.Color("#87CEEB"))

	statusBarStyle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
		Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	locationTagStyle = lipgloss.NewStyle().
		Inherit(statusBarStyle).
		Foreground(lipgloss.Color(offWhite)).
		Background(lipgloss.Color(brightGreen)).
		Padding(0, 1).
		MarginRight(1)

	locationTextStyle = lipgloss.NewStyle().Inherit(statusBarStyle)

	// rename
	ipBlockStyles = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Padding(0, 1)

	// rename
	ipValStyle = ipBlockStyles.
		Background(goldYellow).
		Align(lipgloss.Right)

	// rename
	ipTagStyle = ipBlockStyles.Background(brightBlue)
)

func render(viewmodel internal.TodayViewModel) string {
	doc := strings.Builder{}

	w := lipgloss.Width

	ipT := ipTagStyle.Render("IP ADDRESS")
	ipV := ipValStyle.Render(viewmodel.IP)

	locationH := locationTagStyle.Render("LOCATION")
	locationT := locationTextStyle.
		Width(width - w(locationH) - w(ipT) - w(ipV)).
		Render(fmt.Sprintf("Latitude %s, Longitude %s", viewmodel.Lat, viewmodel.Lng))

	// hmm
	statusBar := lipgloss.JoinHorizontal(lipgloss.Top,
		locationH,
		locationT,
		ipT,
		ipV,
	)

	doc.WriteString(title())
	doc.WriteString(subHead())
	doc.WriteString(today())
	doc.WriteString("rises, noon, sets, length \n")
	doc.WriteString("progress bar thing \n")
	doc.WriteString(tableHead())
	doc.WriteString("date, rises, sets, length \n")
	doc.WriteString(statusBarStyle.Width(width).Render(statusBar) + "\n")
	doc.WriteString(linkString())
	return doc.String()
}

func title() string {
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(goldYellow).
		BorderBottom(true).
		Foreground(goldYellow)

	return style.Width(width).Render("Daylight") + "\n"
}

func subHead() string {
	style := lipgloss.NewStyle().
		BorderBottom(true).
		Foreground(brightBlue).
		Padding(1, 0, 0, 0)

	return style.Width(width).Render("Today's sun times") + "\n"
}

func today() string {
	graphic := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Width(16)

	col := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Width(25).
		Height(8)

	rises := drawPixels(sunrise)
	sets := drawPixels(sunset)

	return lipgloss.JoinHorizontal(lipgloss.Top,
		graphic.Render(rises),
		col.Align(lipgloss.Left).Render("Rises at"),
		graphic.Render(sets),
		col.Align(lipgloss.Left).Render("sets at"),
		col.Align(lipgloss.Right).Render("LENGTH------------------"))
	// ════════════════════════════════════════════════════════════════════════════════════════════════════
	//  RISES--------------------  NOON---------------------  SETS---------------------  LENGTH------------------ .
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

func tableHead() string {
	style := lipgloss.NewStyle().
		Foreground(brightBlue).
		Padding(1, 0, 0, 0)

	return style.Width(width).Render("10-day view") + "\n"
}

func linkString() string {
	style := lipgloss.NewStyle().
		Width(width).
		Foreground(pink).
		Padding(1, 0, 0)

	return style.Render("https://github.com/jbreckmckye/daylight") + "\n"
}
