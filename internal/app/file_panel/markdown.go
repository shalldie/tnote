package file_panel

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/tnote/internal/app/astyles"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/gist"
	"github.com/shalldie/tnote/internal/utils"
)

var (
// titleStyle = func() lipgloss.Style {
// 	return lipgloss.NewStyle().Padding(0, 1)
// 	// b := lipgloss.RoundedBorder()
// 	// b.Right = "├"
// 	// return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
// }()

//	infoStyle = func() lipgloss.Style {
//		// b := lipgloss.RoundedBorder()
//		// b.Left = "┤"
//		// return titleStyle.Copy().BorderStyle(b)
//		return lipgloss.NewStyle().Padding(0, 1)
//	}()
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
	// model.Resize(100, 30)
	// model.Viewport.HighPerformanceRendering = true
	// model.Viewport.Style = lipgloss.NewStyle().
	// Border(lipgloss.RoundedBorder(), true).
	// BorderForeground(lipgloss.Color("#282a35"))
	// model.Viewport.Style = model.Viewport.Style.Padding(0)
	return model
}

// func (m *MarkdownModel) Focus() {
// 	m.BaseModel.Focus()
// }

func (m *MarkdownModel) Resize(width int, height int) {
	m.BaseModel.Resize(width, height)
	headerHeight := lipgloss.Height(m.headerView())
	footerHeight := lipgloss.Height(m.footerView())

	m.Viewport.Width = width
	m.Viewport.Height = height - headerHeight - footerHeight
}

func (m MarkdownModel) Init() tea.Cmd {
	return nil
}

func (m MarkdownModel) getMarkdownContent(content string) string {
	background := "light"

	if lipgloss.HasDarkBackground() {
		background = "dark"
	}

	// background := "notty"
	// background = "dracula"

	r, _ := glamour.NewTermRenderer(
		glamour.WithWordWrap(m.Width),
		glamour.WithStandardStyle(background),
		// glamour.WithAutoStyle(),
	)

	out, err := r.Render(content)
	if err != nil {
		// return "", errors.Unwrap(err)
		return errors.Unwrap(err).Error()
		// return content
	}

	// return content
	return out
}

func (m MarkdownModel) propagate(msg tea.Msg) (MarkdownModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	if m.Active {
		m.Viewport, cmd = m.Viewport.Update(msg)
		cmds = append(cmds, cmd)

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

	case store.CMD_UPDATE_FILE:
		// m.Resize(m.Width, m.Height)
		curFile := store.Gist.GetFile()
		m.file = curFile
		if curFile != nil {
			m.Viewport.SetContent(
				lipgloss.NewStyle().Width(m.Width).Height(m.Height).
					Render(m.getMarkdownContent(curFile.Content)),
			)
			m.Viewport.SetYOffset(0)
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
	// style := lipgloss.NewStyle().Width(m.Width).Height(m.Height).
	// 	Background(lipgloss.Color("#282a35"))

	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.Viewport.View(), m.footerView())
	// return m.Viewport.View()
}

func (m MarkdownModel) withActiveStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(utils.Ternary(m.Active, astyles.PRIMARY_ACTIVE_COLOR, astyles.PRIMARY_NORMAL_COLOR))
}

func (m MarkdownModel) headerView() string {
	titleStyle := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Padding(0, 1).Bold(m.Active)
	title := func() string {
		if m.file == nil {
			return titleStyle.Render("")
		}
		return titleStyle.Render(m.file.FileName)
	}()
	// title := titleStyle.Render("Mr. Pager")
	line := m.withActiveStyle().Render(strings.Repeat("─", utils.MathMax(0, m.Viewport.Width-lipgloss.Width(title))))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m MarkdownModel) footerView() string {
	infoStyle := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Padding(0, 1).Bold(true)
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := m.withActiveStyle().Render(strings.Repeat("─", utils.MathMax(0, m.Viewport.Width-lipgloss.Width(info))))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}