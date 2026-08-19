package status_bar

import (
	"time"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	zone "github.com/lrstanley/bubblezone/v2"
	"github.com/shalldie/tnote/internal/app/pkgs/model"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/conf"
	"github.com/shalldie/tnote/internal/i18n"
)

var S_ID = 1

type StatusBarModel struct {
	*model.BaseModel

	spinner spinner.Model
}

func (m StatusBarModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m StatusBarModel) propagate(msg tea.Msg) (StatusBarModel, tea.Cmd) {

	var cmd tea.Cmd
	var cmds []tea.Cmd

	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m StatusBarModel) Update(msg tea.Msg) (StatusBarModel, tea.Cmd) {
	switch msg := msg.(type) {

	// 	return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {

		case "f10":
			go store.Send(store.CMD_SHOW_PLATFORM(true))
			return m, nil

		case "f12":
			go m.showAbout()
			return m, nil

		}

	case store.CMD_SHOW_PLATFORM:
		if msg {
			go m.showPlatform()
		}
		return m, nil

	case tea.MouseClickMsg:
		if msg.Button != tea.MouseLeft || store.State.InputFocus {
			return m, nil
		}
		if zone.Get(PLATFORM_ID).InBounds(msg) {
			go store.Send(store.CMD_SHOW_PLATFORM(true))
		}
		if zone.Get(ABOUT_ID).InBounds(msg) {
			go m.showAbout()
		}
		return m, nil

	case store.StatusPayload:
		if store.State.Status.Duration > 0 {
			go func() {
				S_ID++
				curId := S_ID
				time.Sleep(time.Second * time.Duration(store.State.Status.Duration))
				if S_ID != curId {
					return
				}
				store.Send(store.StatusPayload{
					Loading: false,
					Message: "",
				})
			}()
		}
		return m, nil

	}

	return m.propagate(msg)
}

func (m StatusBarModel) View() string {
	// 基础样式
	baseStyle := lipgloss.NewStyle().
		// Foreground(lipgloss.Color("#FFFDF5")).
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("#3c3836")).
		Padding(0, 1)

	// status
	statusStyle := baseStyle.Copy().
		// Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
		Background(lipgloss.Color("#F25D94"))
	statusCol := statusStyle.Render(m.spinner.View())
	if !store.State.Status.Loading {
		statusStyle = statusStyle.Copy().Background(lipgloss.Color("#42b883"))
		statusCol = statusStyle.Render("✔") // √,✓,✔
	}

	// help
	// helpStyle := baseStyle.Copy().Background(lipgloss.Color("#A550DF"))
	// helpCol := helpStyle.Render("🛎️  Help - F12")

	// platform
	pfStyle := baseStyle.Copy().Background(lipgloss.Color("#A550DF"))
	pfCol := zone.Mark(PLATFORM_ID, pfStyle.Render(conf.PF_CURRENT+" - F10"))

	// version
	versionStyle := baseStyle.Copy().Background(lipgloss.Color("#6124DF"))
	versionCol := zone.Mark(ABOUT_ID, versionStyle.Render(i18n.Get(i18nTpl, "about")))

	// SPACE
	w := lipgloss.Width
	spaceCol := baseStyle.Copy().
		// Foreground(lipgloss.Color("#FFFDF5")).
		// Background(lipgloss.Color("#6124DF")).
		// Width(m.Width - w(statusCol) - w(versionCol) - w(helpCol)).
		Width(m.Width - w(statusCol) - w(pfCol) - w(versionCol)).
		Render(store.State.Status.Message)

	return lipgloss.JoinHorizontal(lipgloss.Top,
		statusCol,
		spaceCol,
		// helpCol,
		pfCol,
		versionCol,
	)

}

func New() StatusBarModel {
	s := spinner.New()
	s.Spinner = spinner.Line
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))

	return StatusBarModel{
		BaseModel: model.NewBaseModel(),
		spinner:   s,
	}
}
