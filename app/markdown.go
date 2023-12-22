package app

import (
	"errors"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

// RenderMarkdown renders the markdown content with glamour.
// func RenderMarkdown(width int, content string) string {
// 	background := "light"

// 	if lipgloss.HasDarkBackground() {
// 		background = "dark"
// 	}

// 	r, _ := glamour.NewTermRenderer(
// 		glamour.WithWordWrap(width),
// 		glamour.WithStandardStyle(background),
// 	)

// 	out, err := r.Render(content)
// 	if err != nil {
// 		// return "", errors.Unwrap(err)
// 		return errors.Unwrap(err).Error()
// 		// return content
// 	}

// 	return out
// }

type MarkdownModel struct {
	*BaseModel

	Viewport viewport.Model
}

func NewMarkdownModel() MarkdownModel {
	model := MarkdownModel{
		BaseModel: newBaseModel(),
		Viewport:  viewport.New(0, 0),
	}
	// model.Resize(100, 30)
	// model.Viewport.HighPerformanceRendering = true
	return model
}

// func (m *MarkdownModel) Focus() {
// 	m.BaseModel.Focus()
// }

func (m *MarkdownModel) Resize(width int, height int) {
	m.BaseModel.Resize(width, height)
	m.Viewport.Width = width
	m.Viewport.Height = height
}

func (m MarkdownModel) Init() tea.Cmd {
	return nil
}

func (m MarkdownModel) getMarkdownContent(content string) string {
	background := "light"

	if lipgloss.HasDarkBackground() {
		background = "dark"
	}

	r, _ := glamour.NewTermRenderer(
		glamour.WithWordWrap(m.Width),
		glamour.WithStandardStyle(background),
	)

	out, err := r.Render(content)
	if err != nil {
		// return "", errors.Unwrap(err)
		return errors.Unwrap(err).Error()
		// return content
	}

	return out
}

func (m MarkdownModel) propagate(msg tea.Msg) (MarkdownModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.Viewport, cmd = m.Viewport.Update(msg)
	cmds = append(cmds, cmd)
	if m.Active {
		// curItem := m.list.SelectedItem()
		// if fli, ok := curItem.(FileListItem); ok {
		// 	curFilename := fli.gistfile.FileName
		// 	defer m.selectFile(curFilename)
		// }
		// m.list, cmd = m.list.Update(msg)
		// cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m MarkdownModel) Update(msg tea.Msg) (MarkdownModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// m.Resize(msg.Width, msg.Height)
		// cmds = append(cmds, viewport.Sync(m.Viewport))
		// return m, nil

	case CMD_UPDATE_FILE:
		// m.Resize(m.Width, m.Height)
		curFile := gt.GetFile()
		if curFile != nil {
			m.Viewport.SetContent(
				lipgloss.NewStyle().Width(m.Width).Height(m.Height).
					Render(m.getMarkdownContent(curFile.Content)),
			)
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		// case "left":
		// 	m.Active = false
		// case "right":
		// 	m.Active = true
		}

	}
	m, cmd = m.propagate(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MarkdownModel) View() string {
	return m.Viewport.View()
}
