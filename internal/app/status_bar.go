package app

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type StatusBarModel struct {
	*BaseModel
	// Width       int
	LoadingText string
	Loading     bool

	spinner spinner.Model
}

func (m StatusBarModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m StatusBarModel) propagate(msg tea.Msg) (StatusBarModel, tea.Cmd) {

	var cmd tea.Cmd
	var cmds []tea.Cmd

	// Propagate to all children.
	// m.tabs, _ = m.tabs.Update(msg)
	// m.dialog, _ = m.dialog.Update(msg)
	// m.list1, _ = m.list1.Update(msg)
	// m.list2, _ = m.list2.Update(msg)
	// m.LoadingText = "lalala"

	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)

	// switch msg := msg.(type) {
	// case spinner.TickMsg:
	// 	// println("lalala")
	// 	m.LoadingText = "lalala"
	// 	m.spinner, _ = m.spinner.Update(msg)
	// }

	// if msg, ok := msg.(tea.WindowSizeMsg); ok {
	// 	msg.Height -= m.tabs.(tabs).height + m.list1.(list).height
	// 	m.history, _ = m.history.Update(msg)
	// 	return m
	// }

	// m.history, _ = m.history.Update(msg)
	return m, tea.Batch(cmds...)
}

func (m StatusBarModel) Update(msg tea.Msg) (StatusBarModel, tea.Cmd) {
	switch msg := msg.(type) {
	// case tea.WindowSizeMsg:
	// 	m.Width = msg.Width
	// 	return m.propagate(msg), nil

	case CMD_APP_LOADING:
		m.Loading = len(msg) > 0
		m.LoadingText = string(msg)

		return m, nil

	}

	// case errMsg:
	// 	m.err = msg
	// 	return m, nil

	// default:
	// 	var cmd tea.Cmd
	// 	m.spinner, cmd = m.spinner.Update(msg)
	// 	return m, cmd
	// 	// return m.propagate(msg), nil
	// }

	return m.propagate(msg)
}

func (m StatusBarModel) View() string {
	// 基础样式
	baseStyle := lipgloss.NewStyle().
		// Foreground(lipgloss.Color("#FFFDF5")).
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.AdaptiveColor{Light: "#3c3836", Dark: "#3c3836"}).
		Padding(0, 1)

	// status
	statusStyle := baseStyle.Copy().
		// Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
		Background(lipgloss.AdaptiveColor{Light: "#F25D94", Dark: "#F25D94"})
	statusCol := statusStyle.Render(m.spinner.View())
	if !m.Loading {
		statusStyle = statusStyle.Copy().Background(lipgloss.Color("#42b883"))
		statusCol = statusStyle.Render("✔") // √,✓,✔
	}

	// help
	helpStyle := baseStyle.Copy().Background(lipgloss.Color("#A550DF"))
	helpCol := helpStyle.Render("🛎️  Help - F12")

	// version
	versionStyle := baseStyle.Copy().Background(lipgloss.Color("#6124DF"))
	versionCol := versionStyle.Render("TNOTE - V1.0.0")

	// SPACE
	w := lipgloss.Width
	spaceCol := baseStyle.Copy().
		// Foreground(lipgloss.Color("#FFFDF5")).
		// Background(lipgloss.Color("#6124DF")).
		Width(m.Width - w(statusCol) - w(versionCol) - w(helpCol)).
		Render(m.LoadingText)

	return lipgloss.JoinHorizontal(lipgloss.Top,
		statusCol,
		spaceCol,
		helpCol,
		versionCol,
	)

}

func NewStatusBar() StatusBarModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))

	return StatusBarModel{
		BaseModel: newBaseModel(),
		spinner:   s,
	}
}
