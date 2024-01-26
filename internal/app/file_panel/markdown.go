package file_panel

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/gist"
	"github.com/shalldie/tnote/internal/utils"
)

type MarkdownModel struct {
	*model.BoxModel

	Viewport viewport.Model

	file *gist.GistFile
}

func NewMarkdownModel() MarkdownModel {
	model := MarkdownModel{
		BoxModel: model.NewBoxModel(),
		Viewport: viewport.New(0, 0),
	}

	return model
}

func (m *MarkdownModel) Resize(width int, height int) {
	m.BoxModel.Resize(width, height)

	m.Viewport.Width = width - 2
	m.Viewport.Height = height - 2

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
			lipgloss.NewStyle().Width(m.Viewport.Width).Height(m.Viewport.Height).
				Render(utils.RenderMarkdown(curFile.Content, m.Viewport.Width)),
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
		if msg.Button == tea.MouseButtonWheelDown && msg.Action == tea.MouseActionPress && m.Active {
			m.Viewport.SetYOffset(m.Viewport.YOffset + 1)
		}
		// 向上 滚动
		if msg.Button == tea.MouseButtonWheelUp && msg.Action == tea.MouseActionPress && m.Active {
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
	m.HTitle = func() string {
		if m.file == nil {
			return ""
		}
		return m.file.FileName
	}()

	m.FTitle = fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100)

	return zone.Mark(m.ID, m.Render(m.Viewport.View()))
}
