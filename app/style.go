package app

import "github.com/charmbracelet/lipgloss"

var primaryColor = lipgloss.Color("#874BFD")

var grayColor = lipgloss.AdaptiveColor{Light: "#3c3836", Dark: "#3c3836"}

var boxStyle = lipgloss.NewStyle().
	// Background(lipgloss.Color("#282a35")).
	Border(lipgloss.RoundedBorder(), true).
	BorderForeground(grayColor).
	Padding(0, 0)

var boxActiveStyle = boxStyle.Copy().
	BorderForeground(primaryColor)
