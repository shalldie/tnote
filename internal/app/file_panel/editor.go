package file_panel

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/tnote/internal/app/astyles"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/gist"
	"github.com/shalldie/tnote/internal/utils"
)

type EditorModel struct {
	*model.BaseModel

	TextArea textarea.Model
}

func (m *EditorModel) Resize(width int, height int) {
	m.BaseModel.Resize(width, height)
	m.TextArea.SetWidth(width)
	m.TextArea.SetHeight(height - 2)
}

func (m *EditorModel) Focus() {
	m.BaseModel.Focus()
	m.TextArea.Focus()
	store.State.InputFocus = true
}

func (m *EditorModel) Blur() {
	m.BaseModel.Blur()
	m.TextArea.Blur()
	store.State.InputFocus = false
}

func (m *EditorModel) Save() {
	file := store.Gist.GetFile()
	content := m.TextArea.Value()

	m.Blur()

	if file.Content == content {
		return
	}

	go store.Send(store.StatusPayload{
		Loading: true,
		Message: "保存中...",
	})
	store.Gist.UpdateFile(file.FileName, &gist.UpdateGistPayload{Content: content})
	go store.Send(store.CMD_REFRESH_FILES(""))
	go store.Send(store.CMD_UPDATE_FILE(""))
	go store.Send(store.StatusPayload{
		Loading:  false,
		Message:  fmt.Sprintf("「%v」保存完毕", file.FileName),
		Duration: 5,
	})
}

func (m EditorModel) Init() tea.Cmd {
	return tea.Batch(textarea.Blink)
}

func (m EditorModel) propagate(msg tea.Msg) (EditorModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	// m.Markdown, cmd = m.Markdown.Update(msg)
	// cmds = append(cmds, cmd)

	m.TextArea, cmd = m.TextArea.Update(msg)
	cmds = append(cmds, cmd)

	if m.Active {
		// curItem := m.list.SelectedItem()
		// if fli, ok := curItem.(FileListItem); ok {
		// 	curFilename := fli.gistfile.FileName
		// 	defer m.selectFile(curFilename)
		// }
		// m.list, cmd = m.list.Update(msg)
		// cmds = append(cmds, cmd)
		// cmds = append(cmds, m.TextArea.Focus())
	}
	// cmds = append(cmds, m.TextArea.Focus())
	return m, tea.Batch(cmds...)
}

func (m EditorModel) Update(msg tea.Msg) (EditorModel, tea.Cmd) {

	switch msg := msg.(type) {
	// case tea.WindowSizeMsg:
	// 	m.Resize(msg.Width, msg.Height)
	// 	return m, nil
	// }

	case store.CMD_UPDATE_FILE:
		// m.Resize(m.Width, m.Height)
		curFile := store.Gist.GetFile()
		if curFile != nil {
			m.TextArea.SetValue(curFile.Content)
			for i := 0; i < 2*len(strings.Split(curFile.Content, "\n")); i++ {
				m.TextArea.CursorUp()
			}
			m.TextArea.CursorStart()
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.TextArea.Focused() {
				go store.Send(store.CMD_INVOKE_EDIT(false))
				go m.Save()
				// m.Blur()
				// go store.Send(store.CMD_INVOKE_EDIT(false))
				// todo: 更新gist文件

				return m, nil
			}

			// case "right":
			// 	m.Active = true
		}

	}

	return m.propagate(msg)
}

func (m EditorModel) View() string {
	if !m.Active {
		return ""
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.TextArea.View(), m.footerView())
	// style := lipgloss.NewStyle().BorderTop(true)
	// return style.Render(m.TextArea.View())
}

func (m EditorModel) headerView() string {
	titleStyle := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Padding(0, 1).Bold(m.Active)
	title := func() string {
		// if store
		file := store.Gist.GetFile()
		if file == nil {
			return titleStyle.Render("")
		}
		return titleStyle.Render(file.FileName)
	}()
	// title := titleStyle.Render("Mr. Pager")
	line := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Render(strings.Repeat("━", utils.MathMax(0, m.Width-lipgloss.Width(title))))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m EditorModel) footerView() string {
	// 不严谨，但就先这样了
	percent := float64(m.TextArea.Line()+1) / float64(m.TextArea.LineCount()) * 100

	infoStyle := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Padding(0, 1).Bold(true)
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", percent))
	// info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	// line := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Render(strings.Repeat("─", m.Width))
	line := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Render(strings.Repeat("━", utils.MathMax(0, m.Width-lipgloss.Width(info))))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func NewEditorModel() EditorModel {
	ta := textarea.New()
	ta.Placeholder = "please enter..."
	ta.CharLimit = 0

	return EditorModel{
		BaseModel: model.NewBaseModel(),
		TextArea:  ta,
	}
}
