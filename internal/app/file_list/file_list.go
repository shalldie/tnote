package file_list

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/gog/gs"
	"github.com/shalldie/tnote/internal/app/astyles"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/gist"
)

type FileListModel struct {
	*model.BaseModel

	gist    *gist.Gist
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

	case store.CMD_REFRESH_FILES:
		if len(string(msg)) > 0 {
			go store.Send(store.CMD_SELECT_FILE(string(msg)))
		}
		return m, m.refreshFiles()

	case store.CMD_SELECT_FILE:
		m.selectFile(string(msg))
		return m, nil

	case tea.KeyMsg:
		// 输入框没有焦点，且不是正在输入过滤项
		if !store.State.InputFocus && m.list.FilterState() != list.Filtering {
			switch msg.String() {
			// case "left":
			// 	m.Active = true
			// case "right":
			// 	m.Active = false

			case "n":
				go m.newFile()
				return m, nil

			case "d":
				file := m.gist.GetFile()
				if file != nil {
					go m.delFile(file.FileName)
				}
				return m, nil

			case "r":
				file := m.gist.GetFile()
				if file != nil {
					go m.renameFile(file.FileName)
				}
				return m, nil

			}
		}

	}

	return m.propagate(msg)
}

func (m FileListModel) View() string {
	style := lipgloss.NewStyle().
		// Background(lipgloss.Color("#282a35")).
		Border(lipgloss.RoundedBorder(), true).
		// BorderForeground(grayColor).
		BorderForeground(astyles.PRIMARY_NORMAL_COLOR)
	if m.Active {
		style = style.BorderForeground(astyles.PRIMARY_ACTIVE_COLOR)
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
	items := gs.Map(m.gist.Files, func(f *gist.GistFile, i int) list.Item {
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
	targetIndex = gs.FindIndex(m.gist.Files, func(item *gist.GistFile, i int) bool {
		return item.FileName == filename
	})

	if targetIndex != m.gist.CurrentIndex {
		go func() {
			m.gist.CurrentIndex = targetIndex

			store.Send(store.CMD_UPDATE_FILE(""))
		}()
	}
}

func New(g *gist.Gist) FileListModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#fff"))

	model := FileListModel{
		BaseModel: model.NewBaseModel(),
		gist:      g,
		spinner:   s,
		list:      list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
	}
	model.Focus()
	model.list.Title = "文件列表"
	// model.list.SetShowHelp(false)
	model.list.DisableQuitKeybindings()
	model.list.KeyMap = newListKeyMap()
	model.list.AdditionalFullHelpKeys = func() []key.Binding {
		return additionalKeyMap()
	}

	// model.list.AdditionalShortHelpKeys = func() []key.Binding {
	// 	return additionalKeyMap()
	// }
	// model.list.SetShowPagination(false)
	return model
}
