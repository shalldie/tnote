package dialog

import "github.com/charmbracelet/lipgloss"

var dialogBoxStyle = lipgloss.NewStyle().
	// Border(lipgloss.ThickBorder()).
	// BorderForeground(lipgloss.Color("#874BFD")).
	Padding(1, 3)

var buttonStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFF7DB")).
	Background(lipgloss.Color("#888B7E")).
	Padding(0, 3).
	MarginTop(1).
	MarginLeft(2)

var activeButtonStyle = buttonStyle.Copy().
	Foreground(lipgloss.Color("#FFF7DB")).
	Background(lipgloss.Color("#F25D94")).
	// MarginRight(2).
	Underline(true)
