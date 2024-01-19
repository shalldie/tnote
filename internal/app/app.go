package app

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/shalldie/tnote/internal/app/file_list"
	"github.com/shalldie/tnote/internal/app/file_panel"
	"github.com/shalldie/tnote/internal/app/pkgs/dialog"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
	"github.com/shalldie/tnote/internal/app/status_bar"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/conf"
)

var (
	app *tea.Program
	// gt  *gist.Gist
)

type AppModel struct {
	*model.BaseModel

	// components
	FileList    file_list.FileListModel
	FilePanel   file_panel.FilePanelModel
	StatusBar   status_bar.StatusBarModel
	DialogModel dialog.DialogModel
}

func newAppModel() AppModel {

	m := AppModel{
		BaseModel: model.NewBaseModel(),

		FileList:    file_list.New(),
		FilePanel:   file_panel.New(),
		StatusBar:   status_bar.New(),
		DialogModel: dialog.New(),
	}

	return m
}

func (m *AppModel) Resize(width int, height int) {
	m.BaseModel.Resize(width, height)

	lWidth := conf.FileListWidth
	m.FileList.Resize(lWidth, height-3)
	m.FilePanel.Resize(width-lWidth-4, height-1)
	m.StatusBar.Resize(width, 1)
	m.DialogModel.Resize(width, height)
}

func (m *AppModel) Blur() {
	m.BaseModel.Blur()

	m.FileList.Blur()
	m.FilePanel.Blur()
	m.StatusBar.Blur()
	m.DialogModel.Blur()
}

// index: 1-filelist 2-filepanel
func (m *AppModel) focusPanel(index int) bool {
	if m.DialogModel.Active || store.State.InputFocus {
		return false
	}
	m.Blur()

	if index == 1 {
		m.FileList.Focus()
	}
	if index == 2 {
		m.FilePanel.Focus()
	}

	return true
}

func (m AppModel) Init() tea.Cmd {

	return tea.Batch(
		m.FileList.Init(),
		m.FilePanel.Init(),
		m.StatusBar.Init(),
		m.DialogModel.Init(),
	)
}

func (m AppModel) propagate(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	// Propagate to all children.

	m.StatusBar, cmd = m.StatusBar.Update(msg)
	cmds = append(cmds, cmd)

	m.FileList, cmd = m.FileList.Update(msg)
	cmds = append(cmds, cmd)

	m.FilePanel, cmd = m.FilePanel.Update(msg)
	cmds = append(cmds, cmd)

	if m.DialogModel.Active {
		m.DialogModel, cmd = m.DialogModel.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmds []tea.Cmd
	// var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Resize(msg.Width, msg.Height)
		// msg.Height -= 2
		// msg.Width -= 4
		return m, nil

	case dialog.DialogPayload:
		m.Blur()
		m.DialogModel.Show(&msg)
		return m, nil

	case store.CMD_APP_FOCUS:
		if m.focusPanel(int(msg)) {
			return m, nil
		}

	case store.CMD_INVOKE_EDIT:
		if store.State.Editing {
			m.focusPanel(2)
		} else {
			m.focusPanel(1)
		}
		return m, nil

	// case CMD_APP_LOADING:
	// 	// m.loading = len(msg) > 0
	// 	return m.propagate(msg)

	// case spinner.TickMsg:
	// 	var cmd tea.Cmd
	// 	m.spinner, cmd = m.spinner.Update(msg)
	// 	return m, cmd

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		case "left":
			if m.focusPanel(1) {
				return m, nil
			}

		case "right":
			if m.focusPanel(2) {
				return m, nil
			}

		case "e":
			// if !m.textarea.Focused() {
			// 	m.textarea.Focus()
			// }

		case "esc":
			// if m.textarea.Focused() {
			// 	m.textarea.Blur()
			// }

		// These keys should exit the program.
		case "ctrl+c":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":

		// The "down" and "j" keys move the cursor down
		case "down", "j":

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			// _, ok := m.selected[m.cursor]
			// if ok {
			// 	delete(m.selected, m.cursor)
			// } else {
			// 	m.selected[m.cursor] = struct{}{}
			// }
		}

	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	// return m, nil
	// m.textarea, cmd = m.textarea.Update(msg)
	// cmds = append(cmds, cmd)
	return m.propagate(msg)
}

func (m AppModel) View() string {

	viewContainer := lipgloss.NewStyle().
		Height(m.Height - 1).Render(
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.FileList.View(),
			m.FilePanel.View(),
		),
	)

	dialogView := m.DialogModel.View()
	viewContainer = dialog.PlaceOverlay(
		m.Width/2-lipgloss.Width(dialogView)/2, m.Height/2-lipgloss.Height(dialogView)/2-3,
		dialogView,
		viewContainer,
	)

	return zone.Scan(lipgloss.JoinVertical(
		lipgloss.Top,
		viewContainer,
		m.StatusBar.View(),
	))

}

func Run(token string) {

	zone.NewGlobal()

	app = tea.NewProgram(newAppModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())

	go func() {
		// utils.Log("init...")

		store.SendImpl = func(cmd any) {
			app.Send(cmd)
		}

		store.Send(store.StatusPayload{
			Loading: true,
			Message: "loading...",
		})

		store.Setup(token)

		store.Send(store.StatusPayload{Loading: false})
		store.Send(store.CMD_REFRESH_FILES(""))
		store.Send(store.CMD_UPDATE_FILE(""))

		// time.Sleep(time.Second * 3)
		// app.Send(dialog.DialogPayload{
		// 	Message:     "hello world",
		// 	Mode:        1,
		// 	PromptValue: "这个是默认值",
		// })
	}()

	if _, err := app.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
