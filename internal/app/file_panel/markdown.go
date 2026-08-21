package file_panel

import (
	"fmt"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	zone "github.com/lrstanley/bubblezone/v2"
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
		Viewport: viewport.New(),
	}

	return model
}

func (m *MarkdownModel) Resize(width int, height int) {
	m.BoxModel.Resize(width, height)

	m.Viewport.SetWidth(width - 2)
	m.Viewport.SetHeight(height - 2)

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
			lipgloss.NewStyle().Width(m.Viewport.Width()).Height(m.Viewport.Height()).
				Render(utils.RenderMarkdown(curFile.Content, m.Viewport.Width())),
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

	case tea.MouseWheelMsg:
		if !zone.Get(m.ID).InBounds(msg) || store.State.DialogMode {
			return m, nil
		}

		// 向下滚动
		if msg.Button == tea.MouseWheelDown && m.Active {
			m.Viewport.SetYOffset(m.Viewport.YOffset() + 1)
		}
		// 向上 滚动
		if msg.Button == tea.MouseWheelUp && m.Active {
			m.Viewport.SetYOffset(m.Viewport.YOffset() - 1)
		}

		return m, nil

	case tea.MouseClickMsg:
		if !zone.Get(m.ID).InBounds(msg) || store.State.DialogMode {
			return m, nil
		}

		// 点击
		if msg.Button == tea.MouseLeft {
			go store.Send(store.CMD_APP_FOCUS(2))
		}

		return m, nil

	case tea.KeyPressMsg:
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
