package app

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/tnote/gist"
)

var (
	app *tea.Program
	gt  *gist.Gist
)

type AppModel struct {
	*BaseModel
	// state
	// loading bool

	// components
	FileList  FileListModel
	FilePanel FilePanelModel
	StatusBar StatusBarModel
}

func newAppModel() AppModel {
	m := AppModel{
		BaseModel: newBaseModel(),
		FileList:  newFileListModel(),
		FilePanel: newFilePanelModel(),
		StatusBar: NewStatusBar(),
	}

	return m
}

func (m AppModel) Init() tea.Cmd {
	go func() {
		app.Send(CMD_APP_LOADING("loading..."))
		gt.Setup()
		// time.Sleep(time.Second * 3)
		app.Send(CMD_APP_LOADING(""))
		app.Send(CMD_REFRESH_FILES(""))
		app.Send(CMD_UPDATE_FILE(""))
	}()

	return tea.Batch(
		m.FileList.Init(),
		m.FilePanel.Init(),
		m.StatusBar.Init(),
	)
}

func (m AppModel) propagate(msg tea.Msg) (tea.Model, tea.Cmd) {
	// cmds := []tea.Cmd{}
	var cmd tea.Cmd
	var cmds []tea.Cmd
	// Propagate to all children.
	// m.tabs, _ = m.tabs.Update(msg)
	// m.dialog, _ = m.dialog.Update(msg)
	// m.list1, _ = m.list1.Update(msg)
	// m.list2, _ = m.list2.Update(msg)

	m.StatusBar, cmd = m.StatusBar.Update(msg)
	cmds = append(cmds, cmd)

	m.FileList, cmd = m.FileList.Update(msg)
	cmds = append(cmds, cmd)

	m.FilePanel, cmd = m.FilePanel.Update(msg)
	cmds = append(cmds, cmd)

	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		m.FileList.Resize(40, msg.Height-3)
		m.FilePanel.Resize(msg.Width-40-4, msg.Height-3)

		// msg.Height -= m.tabs.(tabs).height + m.list1.(list).height
		// m.history, _ = m.history.Update(msg)
		return m, nil
	}

	// m.history, _ = m.history.Update(msg)
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
		return m.propagate(msg)

	case CMD_APP_LOADING:
		// m.loading = len(msg) > 0
		return m.propagate(msg)

	// case spinner.TickMsg:
	// 	var cmd tea.Cmd
	// 	m.spinner, cmd = m.spinner.Update(msg)
	// 	return m, cmd

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		case "left":
			m.FileList.Focus()
			m.FilePanel.Blur()
			return m, nil
		case "right":
			m.FileList.Blur()
			m.FilePanel.Focus()
			return m, nil

		case "e":
			// if !m.textarea.Focused() {
			// 	m.textarea.Focus()
			// }

		case "esc":
			// if m.textarea.Focused() {
			// 	m.textarea.Blur()
			// }

		// These keys should exit the program.
		case "ctrl+c", "q":
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

	viewContainer := lipgloss.NewStyle().Height(m.Height - 1).Render(
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.FileList.View(),
			m.FilePanel.View(),
		),
	)

	// container := lipgloss.NewStyle().
	// 	// Background(lipgloss.AdaptiveColor{Light: "#F25D94", Dark: "#F25D94"}).
	// 	Height(m.height - 1)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		viewContainer,
		// container.Render(m.filelist.View()),
		m.StatusBar.View(),
	)

}

func RunAppModel(token string) {

	gt = gist.NewGist(token)
	app = tea.NewProgram(newAppModel(), tea.WithAltScreen())

	if _, err := app.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
