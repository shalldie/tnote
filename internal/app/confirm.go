package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
)

// https://github.com/charmbracelet/lipgloss/pull/102/files

type ConfirmModel struct {
	*model.BaseModel
	Content string
}

func (m *ConfirmModel) Show(content string) {
	m.Content = content
	m.Focus()
}

func (m ConfirmModel) Init() tea.Cmd {
	return nil
}

func (m ConfirmModel) Update(msg tea.Msg) (ConfirmModel, tea.Cmd) {
	return m, nil
}

func (m ConfirmModel) View() string {
	dialogBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#874BFD")).
		Padding(1, 3).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)

	buttonStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#888B7E")).
		Padding(0, 3).
		MarginTop(1).
		MarginLeft(2)

	activeButtonStyle := buttonStyle.Copy().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#F25D94")).
		// MarginRight(2).
		Underline(true)

	okButton := activeButtonStyle.Render("确定")
	cancelButton := buttonStyle.Render("取消")

	question := lipgloss.NewStyle().Width(42).Align(lipgloss.Left).Render("确定要删除吗？")
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, cancelButton, okButton)
	ui := lipgloss.JoinVertical(lipgloss.Right, question, buttons)
	return dialogBoxStyle.Render(ui)
}

func NewConfirmModel() ConfirmModel {
	return ConfirmModel{
		BaseModel: model.NewBaseModel(),
	}
}
