package dialog

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/shalldie/gog/gs"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/i18n"
	"github.com/shalldie/tnote/internal/utils"
)

// https://github.com/charmbracelet/lipgloss/pull/102/files

type DialogModel struct {
	*model.BoxModel
	Payload  *DialogPayload
	TabIndex int // 0-textarea 1-btnCancel 2-btnOK

	// components
	TextInput textinput.Model
	Select    list.Model
}

func (m *DialogModel) Focus() {
	m.BoxModel.Focus()
	store.State.DialogMode = true
}

func (m *DialogModel) Blur() {
	m.BoxModel.Blur()
	store.State.DialogMode = false
}

func (m *DialogModel) isSelect() bool {
	return len(m.Payload.SelectList) > 0
}

func (m *DialogModel) isPrompt() bool {
	return m.Payload.Mode == ModePrompt
}

func (m *DialogModel) nextTab() {
	m.TabIndex++

	// 边界处理
	if m.TabIndex > 2 {
		m.TabIndex = 0
	}
	// 非 prompt 没有 textarea
	if m.TabIndex == 0 && !m.isPrompt() {
		m.TabIndex++
	}
	// alert 没有 cancel
	if m.TabIndex == 1 && m.Payload.Mode == ModeAlert {
		m.TabIndex++
	}

	// focus
	if m.TabIndex == 0 {
		m.TextInput.Focus()
		m.TextInput.CursorEnd()
		store.State.InputFocus = true
	} else {
		m.TextInput.Blur()
		store.State.InputFocus = false
	}
}

func (m *DialogModel) Show(payload *DialogPayload) {
	m.Payload = payload
	m.Focus()
	m.TextInput.SetValue(payload.PromptValue)
	m.TabIndex = -1
	m.nextTab()
	if m.isPrompt() {
		m.TextInput.Focus()
	}
	// select
	m.updateSelect()
}

func (m *DialogModel) Close() {
	m.Active = false
	store.State.InputFocus = false
	go store.Send(store.CMD_APP_FOCUS(1))
}

func (m *DialogModel) FnOK() {
	ok := true
	if m.Payload.FnOK != nil {
		result := strings.TrimSpace(m.TextInput.Value())
		if m.isSelect() {
			if item, ok := m.Select.SelectedItem().(selectItem); ok {
				result = string(item)
			}
		}
		ok = m.Payload.FnOK(result)
	}
	if ok {
		m.Close()
	}
}

func (m DialogModel) Init() tea.Cmd {
	return tea.Batch(textinput.Blink)
}

func (m DialogModel) propagate(msg tea.Msg) (DialogModel, tea.Cmd) {
	// cmds := []tea.Cmd{}
	var cmd tea.Cmd
	var cmds []tea.Cmd

	m.TextInput, cmd = m.TextInput.Update(msg)
	cmds = append(cmds, cmd)

	m.Select, cmd = m.Select.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m DialogModel) Update(msg tea.Msg) (DialogModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// tab
		case "tab":
			m.nextTab()
			return m, nil

		// enter
		case "enter":
			if m.TabIndex == 1 {
				m.Close()
				return m, nil
			}
			if m.TabIndex == 2 {
				go m.FnOK()
				return m, nil
			}

		case "esc":
			m.Close()
			return m, nil

		}

	case tea.MouseMsg:
		if msg.Button != tea.MouseButtonLeft {
			return m, nil
		}
		if zone.Get(m.ID + "textarea").InBounds(msg) {
			m.TabIndex = -1 // 0-1
			m.nextTab()
		}
		if zone.Get(m.ID + "btn-cancel").InBounds(msg) {
			m.TabIndex = 0 // 1-1
			m.nextTab()
			m.Close()
		}
		if zone.Get(m.ID + "btn-ok").InBounds(msg) {
			m.TabIndex = 1 // 2-1
			m.nextTab()
			m.FnOK()
		}
		return m, nil

	}
	return m.propagate(msg)
}

func (m DialogModel) View() string {

	if m.Payload == nil || !m.Active {
		return ""
	}

	diaWidth := 46
	ui := ""

	// width
	if m.Payload.Width > 0 {
		diaWidth = m.Payload.Width
	}

	// title
	m.HTitle = m.Payload.Title

	// message
	message := lipgloss.NewStyle().Width(diaWidth).Align(lipgloss.Left).Render(m.Payload.Message)

	// select
	sl := utils.Ternary(m.isSelect(), m.Select.View(), "")

	// prompt
	prompt := zone.Mark(m.ID+"textarea", lipgloss.NewStyle().MarginTop(1).Render(m.TextInput.View()))

	stacks := gs.Filter([]string{
		message,
		sl,
		utils.Ternary(m.Payload.Mode == ModePrompt, prompt, ""),
	}, func(str string, index int) bool {
		return len(str) > 0
	})
	ui = lipgloss.JoinVertical(lipgloss.Top, stacks...)

	// btn
	btnCancel := zone.Mark(m.ID+"btn-cancel", utils.Ternary(m.TabIndex == 1, activeButtonStyle, buttonStyle).Render(i18n.Get(i18nTpl, "cancel")))
	btnOK := zone.Mark(m.ID+"btn-ok", utils.Ternary(m.TabIndex == 2, activeButtonStyle, buttonStyle).Render(i18n.Get(i18nTpl, "ok")))
	buttons := lipgloss.JoinHorizontal(lipgloss.Top,
		utils.Ternary(m.Payload.Mode != ModeAlert, btnCancel, ""),
		btnOK,
	)

	ui = dialogBoxStyle.Render(
		lipgloss.JoinVertical(lipgloss.Right, ui, buttons),
	)

	m.Resize(
		lipgloss.Width(ui)+2, //  border
		lipgloss.Height(ui),
	)

	return m.Render(ui)
}

func New() DialogModel {
	// box
	box := model.NewBoxModel()
	box.BorderActiveColor = lipgloss.Color("#874BFD")

	// input
	input := textinput.New()
	input.Placeholder = i18n.Get(i18nTpl, "placeholder")
	input.Width = 30
	input.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	input.TextStyle = input.PromptStyle

	return DialogModel{
		BoxModel:  box,
		TextInput: input,
		Select:    createSelect(),
	}
}
