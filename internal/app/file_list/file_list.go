package file_list

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/gog/gs"
	"github.com/shalldie/tnote/internal/app/astyles"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/gist"
	"github.com/shalldie/tnote/internal/utils"
)

var listBorderStyle = lipgloss.NewStyle().Border(lipgloss.ThickBorder())

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

func (m *FileListModel) curForeground() lipgloss.Color {
	return utils.Ternary(m.Active, astyles.PRIMARY_ACTIVE_COLOR, astyles.PRIMARY_NORMAL_COLOR)
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

	// case store.CMD_INVOKE_EDIT:
	// 	if store.State.Editing {
	// 		m.Blur()
	// 	} else {
	// 		m.Focus()
	// 	}
	// 	return m, nil

	case tea.KeyMsg:
		// active，输入框没有焦点，且不是正在输入过滤项
		if !store.State.InputFocus && m.list.FilterState() != list.Filtering {
			switch msg.String() {
			// case "left":
			// 	m.Active = true
			// case "right":
			// 	m.Active = false

			case "n":
				go m.newFile()
				return m, nil

			case "r":
				file := store.Gist.GetFile()
				if file != nil {
					go m.renameFile(file.FileName)
				}
				return m, nil

			case "e":
				go store.Send(store.CMD_INVOKE_EDIT(true))
				return m, nil

			case "d":
				file := store.Gist.GetFile()
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
	// lipgloss.ThickBorder()
	style := listBorderStyle.Copy().
		BorderTop(false).
		BorderForeground(m.curForeground())

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
	header := m.headerView()
	body := style.Render(content)
	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		body,
	)
	// return header + "\n" + body
}

func (m FileListModel) headerView() string {
	topRight := listBorderStyle.GetBorderStyle().TopRight

	titleStyle := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Padding(0, 1).Bold(m.Active)
	title := titleStyle.Render("文件列表")
	// title := titleStyle.Render("Mr. Pager")
	// ━┓
	line := lipgloss.NewStyle().Foreground(m.curForeground()).Render(strings.Repeat("━", utils.MathMax(0, m.Width-lipgloss.Width(title)+1)) + topRight)
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m *FileListModel) refreshFiles() tea.Cmd {
	items := gs.Map(store.Gist.Files, func(f *gist.GistFile, i int) list.Item {
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

	// gist
	targetIndex = gs.FindIndex(store.Gist.Files, func(item *gist.GistFile, i int) bool {
		return item.FileName == filename
	})

	if targetIndex != store.Gist.CurrentIndex {
		go func() {
			store.Gist.CurrentIndex = targetIndex

			store.Send(store.CMD_UPDATE_FILE(""))
		}()
	}
}

func New() FileListModel {
	// loading
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#fff"))

	// list item 选中色
	listDelegate := list.NewDefaultDelegate()
	listDelegate.Styles.SelectedTitle = listDelegate.Styles.SelectedTitle.Copy().
		Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#00acf8"}).
		BorderLeftForeground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#00acf8"}).
		Bold(true)

	listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedDesc.Copy().
		Foreground(listDelegate.Styles.NormalDesc.GetForeground()).
		BorderLeftForeground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#00acf8"})

	// model
	model := FileListModel{
		BaseModel: model.NewBaseModel(),
		spinner:   s,
		list:      list.New([]list.Item{}, listDelegate, 0, 0),
	}

	model.Focus()
	model.list.SetShowTitle(false)
	model.list.DisableQuitKeybindings()
	model.list.KeyMap = newListKeyMap()
	model.list.AdditionalFullHelpKeys = additionalKeyMap

	return model
}
