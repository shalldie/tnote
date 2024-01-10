package dialog

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/utils"
)

// https://github.com/charmbracelet/lipgloss/pull/102/files

type DialogModel struct {
	*model.BaseModel
	Payload  *DialogPayload
	TabIndex int // 0-textarea 1-btnCancel 2-btnOK

	// components
	TextInput textinput.Model

	//todo: add OnClose
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
	if m.TabIndex == 0 && !m.isPrompt() {
		m.TabIndex++
	}

	// focus
	if m.TabIndex == 0 {
		m.TextInput.Focus()
		store.State.InputFocus = true
	} else {
		m.TextInput.Blur()
		store.State.InputFocus = false
	}
}

func (m *DialogModel) Show(payload *DialogPayload) {
	m.Payload = payload
	m.Focus()
	m.TabIndex = -1
	m.nextTab()
	m.TextInput.SetValue(payload.PromptValue)
	if m.isPrompt() {
		m.TextInput.Focus()
	}
}

func (m *DialogModel) Close() {
	m.Active = false
	store.State.InputFocus = false
	go store.Send(store.CMD_APP_FOCUS(""))
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
				ok := true
				if m.Payload.FnOK != nil {
					ok = m.Payload.FnOK(strings.TrimSpace(m.TextInput.Value()))
				}
				if ok {
					m.Close()
				}
				return m, nil
			}

		case "esc":
			m.Close()
			return m, nil

		}

	}
	return m.propagate(msg)
}

func (m DialogModel) View() string {

	if m.Payload == nil || !m.Active {
		return ""
	}

	diaWidth := 42
	ui := ""

	// message
	message := lipgloss.NewStyle().Width(diaWidth).Align(lipgloss.Left).Render(m.Payload.Message)

	// prompt
	prompt := lipgloss.NewStyle().Render(m.TextInput.View())

	ui = lipgloss.JoinVertical(lipgloss.Top,
		message,
		utils.Ternary(m.Payload.Mode == 1, prompt, ""),
	)

	// btn
	btnCancel := utils.Ternary(m.TabIndex == 1, activeButtonStyle, buttonStyle).Render("取消")
	btnOK := utils.Ternary(m.TabIndex == 2, activeButtonStyle, buttonStyle).Render("确定")
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, btnCancel, btnOK)

	ui = lipgloss.JoinVertical(lipgloss.Right, ui, buttons)

	return dialogBoxStyle.Render(ui)
}

func New() DialogModel {
	input := textinput.New()
	input.Placeholder = "请输入..."
	input.Width = 20
	input.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	input.TextStyle = input.PromptStyle
	return DialogModel{
		BaseModel: model.NewBaseModel(),
		TextInput: input,
	}
}