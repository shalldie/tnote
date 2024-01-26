package model

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/tnote/internal/app/astyles"
	"github.com/shalldie/tnote/internal/utils"
)

var boxBorder = lipgloss.ThickBorder()

// 基础容器模型
type BoxModel struct {
	*BaseModel

	// header title
	HTitle string

	// footer title
	FTitle string

	BorderNormalColor lipgloss.Color

	BorderActiveColor lipgloss.Color
}

func (m *BoxModel) curForeground() lipgloss.Color {
	return utils.Ternary(m.Active, m.BorderActiveColor, m.BorderNormalColor)
}

func (m *BoxModel) curStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(m.curForeground())
}

// render with inner view
func (m *BoxModel) Render(innerView string) string {
	style := lipgloss.NewStyle().
		Margin(0).
		Width(m.Width - 2).Height(m.Height - 2).
		Border(boxBorder).
		BorderTop(false).BorderBottom(false).
		BorderForeground(m.curForeground())

	body := style.Render(innerView)

	return lipgloss.JoinVertical(lipgloss.Left,
		m.headerView(),
		body,
		m.footerView(),
	)
}

func (m *BoxModel) headerView() string {

	curs := m.curStyle()
	topLeft := curs.Render(boxBorder.TopLeft)
	topRight := curs.Render(boxBorder.TopRight)

	padding := utils.Ternary(len(m.HTitle) > 0, 1, 0)
	titleStyle := lipgloss.NewStyle().Foreground(m.BorderActiveColor).Padding(0, padding).Bold(m.Active)
	title := titleStyle.Render(m.HTitle)

	line := curs.Render(
		strings.Repeat(boxBorder.Top,
			utils.MathMax(0, m.Width-2-lipgloss.Width(title)), // 这里要 -2，去除 topLeft、topRight
		),
	)

	return lipgloss.JoinHorizontal(lipgloss.Center,
		topLeft, title, line, topRight,
	)
}

func (m *BoxModel) footerView() string {

	curs := m.curStyle()
	bottomLeft := curs.Render(boxBorder.BottomLeft)
	bottomRight := curs.Render(boxBorder.BottomRight)

	padding := utils.Ternary(len(m.FTitle) > 0, 1, 0)
	titleStyle := lipgloss.NewStyle().Foreground(m.BorderActiveColor).Padding(0, padding).Bold(m.Active)
	title := titleStyle.Render(m.FTitle)

	line := curs.Render(
		strings.Repeat(boxBorder.Top,
			utils.MathMax(0,
				m.Width-2-lipgloss.Width(title),
			),
		),
	)

	return lipgloss.JoinHorizontal(lipgloss.Center,
		bottomLeft, line, title, bottomRight,
	)
}

func (m *BoxModel) GetBoxStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(boxBorder).
		BorderForeground(utils.Ternary(m.Active, m.BorderNormalColor, m.BorderActiveColor))
}

func NewBoxModel() *BoxModel {
	m := &BoxModel{
		BaseModel:         NewBaseModel(),
		BorderNormalColor: astyles.PRIMARY_NORMAL_COLOR,
		BorderActiveColor: astyles.PRIMARY_ACTIVE_COLOR,
	}

	lipgloss.NewStyle().Border(lipgloss.ThickBorder()).BorderForeground()
	return m
}
