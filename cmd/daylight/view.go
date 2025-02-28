package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/jbreckmckye/daylight/internal"
)

const (
	width = 96
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
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#004889")).
		Padding(0, 1).
		MarginRight(1)

	locationTextStyle = lipgloss.NewStyle().Inherit(statusBarStyle)

	// rename
	ipBlockStyles = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Padding(0, 1)

	// rename
	ipValStyle = ipBlockStyles.
		Background(lipgloss.Color("#A550DF")).
		Align(lipgloss.Right)

	// rename
	ipTagStyle = ipBlockStyles.Background(lipgloss.Color("#6124DF"))
)

func render(viewmodel internal.TodayViewModel) string {
	doc := strings.Builder{}

	w := lipgloss.Width

	titleBar := titleBarStyle.Width(width).Render("Daylight")
	doc.WriteString(titleBar + "\n\n")
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

	doc.WriteString(statusBarStyle.Width(width).Render(statusBar))
	return doc.String()
}
