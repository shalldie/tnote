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
	"github.com/shalldie/tnote/internal/i18n"
	"github.com/shalldie/tnote/internal/utils"
)

type EditorModel struct {
	*model.BoxModel

	TextArea textarea.Model
}

func (m *EditorModel) Resize(width int, height int) {
	m.BoxModel.Resize(width, height)
	m.TextArea.SetWidth(width - 2)
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
	file := store.State.GetFile()
	content := m.TextArea.Value()

	m.Blur()

	if file.Content == content {
		return
	}

	go store.Send(store.StatusPayload{
		Loading: true,
		Message: i18n.Get(i18nTpl, "saving"),
	})
	store.Gist.UpdateFile(file.FileName, &gist.UpdateGistPayload{Content: content})
	go store.Send(store.CMD_REFRESH_FILES(""))
	go store.Send(store.CMD_UPDATE_FILE(""))
	go store.Send(store.StatusPayload{
		Loading:  false,
		Message:  fmt.Sprintf(i18n.Get(i18nTpl, "saved"), file.FileName),
		Duration: 5,
	})
}

func (m EditorModel) Init() tea.Cmd {
	return tea.Batch(textarea.Blink)
}

func (m EditorModel) propagate(msg tea.Msg) (EditorModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.TextArea, cmd = m.TextArea.Update(msg)
	cmds = append(cmds, cmd)

	// if m.Active {
	// }

	return m, tea.Batch(cmds...)
}

func (m EditorModel) Update(msg tea.Msg) (EditorModel, tea.Cmd) {

	switch msg := msg.(type) {

	case store.CMD_UPDATE_FILE:
		// m.Resize(m.Width, m.Height)
		curFile := store.State.GetFile()
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
				return m, nil
			}
		}

	}

	return m.propagate(msg)
}

func (m EditorModel) View() string {
	if !m.Active {
		return ""
	}

	m.HTitle = func() string {
		// if store
		file := store.State.GetFile()
		if file == nil {
			return ""
		}
		return file.FileName
	}()

	m.FTitle = func() string {
		percent := float64(m.TextArea.Line()+1) / float64(m.TextArea.LineCount()) * 100
		return fmt.Sprintf("%3.f%%", percent)
	}()

	return m.Render(m.TextArea.View())
}

func (m EditorModel) headerView() string {
	titleStyle := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Padding(0, 1).Bold(m.Active)
	title := func() string {
		// if store
		file := store.State.GetFile()
		if file == nil {
			return titleStyle.Render("")
		}
		return titleStyle.Render(file.FileName)
	}()
	// title := titleStyle.Render("Mr. Pager")
	line := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Render(strings.Repeat(lipgloss.ThickBorder().Top, utils.MathMax(0, m.Width-lipgloss.Width(title))))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m EditorModel) footerView() string {
	// 不严谨，但就先这样了
	percent := float64(m.TextArea.Line()+1) / float64(m.TextArea.LineCount()) * 100

	infoStyle := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Padding(0, 1).Bold(true)
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", percent))

	line := lipgloss.NewStyle().Foreground(astyles.PRIMARY_ACTIVE_COLOR).Render(strings.Repeat(lipgloss.ThickBorder().Top, utils.MathMax(0, m.Width-lipgloss.Width(info))))

	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func NewEditorModel() EditorModel {
	ta := textarea.New()
	ta.Placeholder = i18n.Get(i18nTpl, "editor_placeholder")
	ta.CharLimit = 0
	ta.Prompt = " "

	return EditorModel{
		BoxModel: model.NewBoxModel(),
		TextArea: ta,
	}
}
