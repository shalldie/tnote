package note

import "github.com/rivo/tview"

type NoteModal struct {
	*tview.Modal
}

func NewNoteModal() *NoteModal {
	return &NoteModal{
		Modal: tview.NewModal(),
	}
}

func (m *NoteModal) Show() {
	note.Pages.ShowPage("modal")
}

func (m *NoteModal) Close() {
	note.Pages.SwitchToPage("main")
}

func (m *NoteModal) Clear() {
	m.ClearButtons()
}

func (m *NoteModal) Confirm(title string, done func()) {
	lastFocus := note.App.GetFocus()

	m.Clear()
	m.SetText(title).
		AddButtons([]string{"确定", "取消"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			m.Close()
			note.App.SetFocus(lastFocus)

			if buttonIndex == 0 {
				done()
			}
		})

	note.Pages.ShowPage("modal")
}
