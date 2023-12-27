package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/gog/gs"
	"github.com/shalldie/tnote/internal/gist"
)

type FileListModel struct {
	*BaseModel

	spinner spinner.Model
	list    list.Model
}

func (m *FileListModel) Resize(width int, height int) {
	m.BaseModel.Resize(width, height)

	m.list.SetWidth(width - 2)
	m.list.SetHeight(height)
}

func (m FileListModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m FileListModel) propagate(msg tea.Msg) (FileListModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)

	if m.Active {
		// curItem := m.list.SelectedItem()
		// if fli, ok := curItem.(FileListItem); ok {
		// 	curFilename := fli.gistfile.FileName
		// 	defer m.selectFile(curFilename)
		// }
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
		if item, ok := m.list.SelectedItem().(FileListItem); ok {
			m.selectFile(item.gistfile.FileName)
			// curIndex := gs.IndexOf(gt.Files, item.gistfile)
			// if curIndex != gt.CurrentIndex {
			// 	go func() {
			// 		gt.CurrentIndex = curIndex
			// 		app.Send(CMD_UPDATE_FILE(""))
			// 	}()
			// }
		}
	}
	return m, tea.Batch(cmds...)
}

func (m FileListModel) Update(msg tea.Msg) (FileListModel, tea.Cmd) {

	switch msg := msg.(type) {
	// case tea.WindowSizeMsg:
	// 	m.Resize(msg.Width, msg.Height)
	// 	return m, nil

	case CMD_REFRESH_FILES:
		return m, m.refreshFiles()

	case tea.KeyMsg:
		switch msg.String() {
		// case "left":
		// 	m.Active = true
		// case "right":
		// 	m.Active = false
		}

	}

	return m.propagate(msg)
}

func (m FileListModel) View() string {
	style := lipgloss.NewStyle().
		// Background(lipgloss.Color("#282a35")).
		Border(lipgloss.RoundedBorder(), true).
		// BorderForeground(grayColor).
		BorderForeground(PRIMARY_NORMAL_COLOR)
	if m.Active {
		style = style.BorderForeground(PRIMARY_ACTIVE_COLOR)
	}
	style = style.
		Width(m.Width).
		Height(m.Height)

	content := func() string {
		if len(m.list.Items()) > 0 {
			return m.list.View()
		}
		return fmt.Sprintf(" %vloading...", m.spinner.View())
	}()
	// content := utils.Ternary(len(m.list.Items()) > 0, m.list.View(), m.spinner.View())
	return style.Render(content)
}

func (m *FileListModel) refreshFiles() tea.Cmd {
	items := gs.Map(gt.Files, func(f *gist.GistFile, i int) list.Item {
		return FileListItem{gistfile: f}
	})
	return m.list.SetItems(items)
}

func (m *FileListModel) selectFile(filename string) {
	// list
	targetIndex := gs.FindIndex[list.Item](m.list.VisibleItems(), func(item list.Item, i int) bool {
		if fli, ok := item.(FileListItem); ok {
			return fli.gistfile.FileName == filename
		}
		return false
	})
	if m.list.Index() != targetIndex {
		m.list.Select(targetIndex)
	}

	// m.list.VisibleItems()

	// gist
	targetIndex = gs.FindIndex(gt.Files, func(item *gist.GistFile, i int) bool {
		return item.FileName == filename
	})
	if targetIndex != gt.CurrentIndex {
		go func() {
			gt.CurrentIndex = targetIndex
			app.Send(CMD_UPDATE_FILE(""))
		}()
	}
}

func newFileListModel() FileListModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#fff"))

	model := FileListModel{
		BaseModel: newBaseModel(),
		spinner:   s,
		list:      list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
	}
	model.Focus()
	model.list.Title = "文件列表"
	// model.list.SetShowHelp(false)
	model.list.DisableQuitKeybindings()
	// model.list.
	// model.list.SetShowPagination(false)
	return model
}

// --------- FileListItem ---------

type FileListItem struct {
	gistfile *gist.GistFile
}

func (item FileListItem) Title() string       { return item.gistfile.FileName }
func (item FileListItem) Description() string { return item.gistfile.Content }
func (item FileListItem) FilterValue() string { return item.Title() }