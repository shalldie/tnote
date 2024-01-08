package app

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/gog/gs"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
	"github.com/shalldie/tnote/internal/gist"
)

type FileListModel struct {
	*model.BaseModel

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
		BaseModel: model.NewBaseModel(),
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

// --------- FileListItem ---------

type FileListItem struct {
	gistfile *gist.GistFile
}

func (item FileListItem) Title() string       { return item.gistfile.FileName }
func (item FileListItem) Description() string { return item.gistfile.Content }
func (item FileListItem) FilterValue() string { return item.Title() }

// --------- KeyMap ---------

func newListKeyMap() list.KeyMap {
	return list.KeyMap{
		// Browsing.
		CursorUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "上"),
		),
		CursorDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "下"),
		),
		PrevPage: key.NewBinding(
			// key.WithKeys("left", "h", "pgup", "b", "u"),
			key.WithKeys("h", "pgup"),
			// key.WithHelp("←/h/pgup", "prev page"),
			key.WithHelp("h/pgup", "上一页"),
		),
		NextPage: key.NewBinding(
			// key.WithKeys("right", "l", "pgdown", "f", "d"),
			key.WithKeys("l", "pgdown"),
			// key.WithHelp("→/l/pgdn", "next page"),
			key.WithHelp("l/pgdn", "下一页"),
		),
		// GoToStart: key.NewBinding(
		// 	key.WithKeys("home", "g"),
		// 	key.WithHelp("g/home", "开头"),
		// ),
		// GoToEnd: key.NewBinding(
		// 	key.WithKeys("end"),
		// 	key.WithHelp("end", "尾部"),
		// ),
		Filter: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "过滤"),
		),
		ClearFilter: key.NewBinding(
			key.WithKeys("esc"),
			// key.WithHelp("esc", "clear filter"),
			key.WithHelp("esc", "清空过滤条件"),
		),

		// Filtering.
		CancelWhileFiltering: key.NewBinding(
			key.WithKeys("esc"),
			// key.WithHelp("esc", "cancel"),
			key.WithHelp("esc", "取消"),
		),
		AcceptWhileFiltering: key.NewBinding(
			key.WithKeys("enter", "tab", "shift+tab", "ctrl+k", "up", "ctrl+j", "down"),
			// key.WithHelp("enter", "apply filter"),
			key.WithHelp("enter", "应用过滤"),
		),

		// Toggle help.
		ShowFullHelp: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "更多"),
		),
		CloseFullHelp: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "折叠"),
		),

		// Quitting.
		// Quit: key.NewBinding(
		// 	key.WithKeys("q", "esc"),
		// 	key.WithHelp("q", "quit"),
		// ),
		// ForceQuit: key.NewBinding(key.WithKeys("ctrl+c")),
	}
}

func additionalKeyMap() []key.Binding {
	return []key.Binding{
		key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "新建"),
		),
		key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "重命名"),
		),
		key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "编辑"),
		),
		key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "删除"),
		),
		// openDirectoryKey,
		// createFileKey,
		// createDirectoryKey,
		// deleteItemKey,
		// copyItemKey,
		// zipItemKey,
		// unzipItemKey,
		// toggleHiddenKey,
		// homeShortcutKey,
		// copyToClipboardKey,
		// escapeKey,
		// renameItemKey,
		// openInEditorKey,
		// submitInputKey,
		// moveItemKey,
	}
}
