package file_panel

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
	"github.com/shalldie/tnote/internal/app/store"
)

type FilePanelModel struct {
	*model.BaseModel
	Markdown MarkdownModel
	Editor   EditorModel
}

func (m *FilePanelModel) Resize(width int, height int) {
	m.BaseModel.Resize(width, height)
	m.Markdown.Resize(width, height)
	m.Editor.Resize(width, height)
}

func (m *FilePanelModel) Focus() {
	if m.Active {
		return
	}

	m.BaseModel.Focus()
	// m.Markdown.Focus()
	if !store.State.Editing {
		m.Markdown.Focus()
	} else {
		m.Editor.Focus()
	}
}

func (m *FilePanelModel) Blur() {
	if !m.Active {
		return
	}
	m.BaseModel.Blur()
	m.Markdown.Blur()
	m.Editor.Blur()
}

func (m FilePanelModel) Init() tea.Cmd {
	return tea.Batch(m.Markdown.Init(), m.Editor.Init())
}

func (m FilePanelModel) propagate(msg tea.Msg) (FilePanelModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.Markdown, cmd = m.Markdown.Update(msg)
	cmds = append(cmds, cmd)

	m.Editor, cmd = m.Editor.Update(msg)
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

	// case store.CMD_INVOKE_EDIT:
	// 	if store.State.Editing {
	// 		m.Markdown.Blur()
	// 		m.Editor.Focus()
	// 	} else {
	// 		m.Blur()
	// 		m.Focus()
	// 	}
	// 	return m, nil

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

	if store.State.Editing {
		return m.Editor.View()
	}
	return m.Markdown.View()
}

func New() FilePanelModel {
	return FilePanelModel{
		BaseModel: model.NewBaseModel(),
		Markdown:  NewMarkdownModel(),
		Editor:    NewEditorModel(),
	}
}
