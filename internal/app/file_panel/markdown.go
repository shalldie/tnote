package file_panel

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/shalldie/tnote/internal/app/astyles"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/gist"
	"github.com/shalldie/tnote/internal/utils"
)

type MarkdownModel struct {
	*model.BaseModel

	Viewport viewport.Model

	file *gist.GistFile
}

func NewMarkdownModel() MarkdownModel {
	model := MarkdownModel{
		BaseModel: model.NewBaseModel(),
		Viewport:  viewport.New(0, 0),
	}

	return model
}

func (m *MarkdownModel) Resize(width int, height int) {
	m.BaseModel.Resize(width, height)
	headerHeight := lipgloss.Height(m.headerView())
	footerHeight := lipgloss.Height(m.footerView())

	m.Viewport.Width = width
	m.Viewport.Height = height - headerHeight - footerHeight

	m.renderFile()
}

func (m MarkdownModel) Init() tea.Cmd {
	return nil
}

func (m *MarkdownModel) renderFile() {
	if store.Gist == nil {
		return
	}
	curFile := store.State.GetFile()
	m.file = curFile
	if curFile != nil {
		m.Viewport.SetContent(
			lipgloss.NewStyle().Width(m.Width).Height(m.Height).
				Render(utils.RenderMarkdown(curFile.Content, m.Width)),
		)
		m.Viewport.SetYOffset(0)
	}
}

func (m MarkdownModel) propagate(msg tea.Msg) (MarkdownModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	if m.Active {
		m.Viewport, cmd = m.Viewport.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m MarkdownModel) Update(msg tea.Msg) (MarkdownModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	case store.CMD_UPDATE_FILE:
		// m.Resize(m.Width, m.Height)
		m.renderFile()
		return m, nil

	case tea.MouseMsg:
		if !zone.Get(m.ID).InBounds(msg) || store.State.DialogMode {
			return m, nil
		}

		// 向下滚动
		if msg.Button == tea.MouseButtonWheelDown && msg.Action == tea.MouseActionPress {
			m.Viewport.SetYOffset(m.Viewport.YOffset + 1)
		}
		// 向上 滚动
		if msg.Button == tea.MouseButtonWheelUp && msg.Action == tea.MouseActionPress {
			m.Viewport.SetYOffset(m.Viewport.YOffset - 1)
		}
		// 点击
		if msg.Button == tea.MouseButtonLeft && msg.Action == tea.MouseActionPress {
			go store.Send(store.CMD_APP_FOCUS(2))
		}

		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		}

	}
	m, cmd = m.propagate(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MarkdownModel) View() string {
	// style := lipgloss.NewStyle().Width(m.Width).Height(m.Height).
	// 	Background(lipgloss.Color("#282a35"))

	return zone.Mark(m.ID, lipgloss.JoinVertical(lipgloss.Center,
		m.headerView(),
		m.Viewport.View(),
		m.footerView(),
	))
	// return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.Viewport.View(), m.footerView())
	// return m.Viewport.View()
}

func (m MarkdownModel) withActiveStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(utils.Ternary(m.Active, astyles.PRIMARY_ACTIVE_COLOR, astyles.PRIMARY_NORMAL_COLOR))
}

func (m MarkdownModel) headerView() string {
	titleStyle := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Padding(0, 1).Bold(m.Active)
	title := func() string {
		if m.file == nil {
			return ""
		}
		return titleStyle.Render(m.file.FileName)
	}()
	// title := titleStyle.Render("Mr. Pager")
	line := m.withActiveStyle().Render(strings.Repeat(lipgloss.ThickBorder().Top, utils.MathMax(0, m.Viewport.Width-lipgloss.Width(title))))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m MarkdownModel) footerView() string {
	infoStyle := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Padding(0, 1).Bold(true)
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := m.withActiveStyle().Render(strings.Repeat(lipgloss.ThickBorder().Top, utils.MathMax(0, m.Viewport.Width-lipgloss.Width(info))))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}
