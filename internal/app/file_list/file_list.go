package file_list

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/shalldie/gog/gs"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/gist"
	"github.com/shalldie/tnote/internal/i18n"
)

type FileListModel struct {
	*model.BoxModel

	spinner spinner.Model
	list    list.Model
}

func (m *FileListModel) Resize(width int, height int) {
	m.BoxModel.Resize(width, height)

	m.list.SetWidth(width - 2)
	m.list.SetHeight(height - 2)
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
		// list 默认的操作只有active时才能使用
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)

		// 所有的操作，选择、过滤等，同时同步到list、gist
		if item, ok := m.list.SelectedItem().(FileListItem); ok {
			m.selectFile(item.FileName)
		}
	}
	return m, tea.Batch(cmds...)
}

func (m FileListModel) Update(msg tea.Msg) (FileListModel, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.MouseMsg:
		if store.State.InputFocus || !zone.Get(m.ID).InBounds(msg) {
			return m, nil
		}
		isHover := m.Active && zone.Get(m.ID).InBounds(msg)
		// 向下滚动
		if isHover && msg.Button == tea.MouseButtonWheelDown && msg.Action == tea.MouseActionPress {
			m.list.CursorDown()
		}
		// 向上 滚动
		if isHover && msg.Button == tea.MouseButtonWheelUp && msg.Action == tea.MouseActionPress {
			m.list.CursorUp()
		}

		// click
		if msg.Button == tea.MouseButtonLeft && msg.Action == tea.MouseActionPress {
			// active
			if !m.Active {
				go store.Send(store.CMD_APP_FOCUS(1))
			}
			// 选择
			for _, listItem := range m.list.VisibleItems() {
				item, _ := listItem.(FileListItem)
				if zone.Get(item.ID+"title").InBounds(msg) || zone.Get(item.ID+"des").InBounds(msg) {
					// m.list.Select(i)
					m.selectFile(item.FileName)
					break
				}
			}
		}

		return m, nil

	case store.CMD_REFRESH_FILES:
		if len(string(msg)) > 0 {
			go store.Send(store.CMD_SELECT_FILE(string(msg)))
		}
		return m, m.refreshFiles()

	case store.CMD_SELECT_FILE:
		m.selectFile(string(msg))
		return m, nil

	case tea.KeyMsg:

		// ready 了才能操作
		if len(m.list.Items()) <= 0 {
			return m, nil
		}

		// active，输入框没有焦点，且不是正在输入过滤项
		if !store.State.InputFocus && m.list.FilterState() != list.Filtering {
			switch msg.String() {

			case "n":
				go m.newFile()
				return m, nil

			case "r":
				file := store.State.GetFile()
				if file != nil {
					go m.renameFile(file.FileName)
				}
				return m, nil

			case "e":
				go store.Send(store.CMD_INVOKE_EDIT(true))
				return m, nil

			case "d":
				file := store.State.GetFile()
				if file != nil {
					go m.delFile(file.FileName)
				}
				return m, nil
			}
		}

	}

	return m.propagate(msg)
}

func (m FileListModel) View() string {

	content := func() string {
		if len(m.list.Items()) > 0 {
			return m.list.View()
		}
		return fmt.Sprintf(" %v loading...", m.spinner.View())
	}()

	return zone.Mark(m.ID, m.Render(content))
}

func (m *FileListModel) refreshFiles() tea.Cmd {
	items := gs.Map(store.Gist.Files, func(f *gist.GistFile, i int) list.Item {
		return FileListItem{
			ID:       m.ID + f.FileName,
			GistFile: f,
		}
	})
	return m.list.SetItems(items)
}

func (m *FileListModel) selectFile(filename string) {
	// list
	targetIndex := gs.FindIndex[list.Item](m.list.VisibleItems(), func(item list.Item, i int) bool {
		if fli, ok := item.(FileListItem); ok {
			return fli.FileName == filename
		}
		return false
	})
	if m.list.Index() != targetIndex {
		m.list.Select(targetIndex)
	}

	// gist
	flItem, _ := m.list.SelectedItem().(FileListItem)
	file := flItem.GistFile
	curFile := store.State.GetFile()

	if file != curFile {
		go func() {
			store.State.SetFile(file)
			store.Send(store.CMD_UPDATE_FILE(""))
		}()
	}
}

func New() FileListModel {
	// loading
	s := spinner.New()
	s.Spinner = spinner.Line
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#fff"))

	// listDelegate，list item 选中色
	listDelegate := list.NewDefaultDelegate()
	listDelegate.Styles.SelectedTitle = listDelegate.Styles.SelectedTitle.Copy().
		Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#00acf8"}).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeftForeground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#00acf8"}).
		Bold(true)

	listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedDesc.Copy().
		Foreground(listDelegate.Styles.NormalDesc.GetForeground()).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeftForeground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#00acf8"})

	// list
	list := list.New([]list.Item{}, listDelegate, 0, 0)
	list.SetShowTitle(false)
	list.DisableQuitKeybindings()
	list.KeyMap = newListKeyMap()
	list.AdditionalFullHelpKeys = additionalKeyMap
	list.FilterInput.Prompt = i18n.Get(i18nTpl, "filelist_filter")

	// box
	boxModel := model.NewBoxModel()
	boxModel.HTitle = i18n.Get(i18nTpl, "filelist_title")

	// model
	model := FileListModel{
		BoxModel: boxModel,
		spinner:  s,
		list:     list,
	}

	model.Focus()

	return model
}
