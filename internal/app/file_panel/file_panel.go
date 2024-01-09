package file_panel

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
)

type FilePanelModel struct {
	*model.BaseModel
	Markdown MarkdownModel
}

func (m *FilePanelModel) Resize(width int, height int) {
	m.BaseModel.Resize(width, height)
	m.Markdown.Resize(width, height)
}

func (m *FilePanelModel) Focus() {
	m.BaseModel.Focus()
	m.Markdown.Focus()
}

func (m *FilePanelModel) Blur() {
	m.BaseModel.Blur()
	m.Markdown.Blur()
}

func (m FilePanelModel) Init() tea.Cmd {
	return tea.Batch(m.Markdown.Init(), m.Markdown.Init())
}

func (m FilePanelModel) propagate(msg tea.Msg) (FilePanelModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.Markdown, cmd = m.Markdown.Update(msg)
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

func (m FilePanelModel) Update(msg tea.Msg) (FilePanelModel, tea.Cmd) {

	switch msg := msg.(type) {
	// case tea.WindowSizeMsg:
	// 	m.Resize(msg.Width, msg.Height)
	// 	return m, nil
	// }
	case tea.KeyMsg:
		switch msg.String() {
		// case "left":
		// 	m.Active = false
		// case "right":
		// 	m.Active = true
		}

	}

	return m.propagate(msg)
}

func (m FilePanelModel) View() string {
	// style := utils.Ternary(m.Active, boxActiveStyle, boxStyle)
	style := lipgloss.NewStyle()
	style = style.Copy().
		Width(m.Width).
		Height(m.Height)
		// Padding(0, 1)

	return style.Render(m.Markdown.View())
}

func New() FilePanelModel {
	return FilePanelModel{
		BaseModel: model.NewBaseModel(),
		Markdown:  NewMarkdownModel(),
	}
}
