package note

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type NoteModal struct {
	*CustomModal
}

func NewNoteModal() *NoteModal {
	return &NoteModal{
		CustomModal: NewCustomModal(),
	}
}

func (m *NoteModal) Show() {
	note.Pages.ShowPage("modal")
}

func (m *NoteModal) Close() {
	note.Pages.SwitchToPage("main")
}

func (m *NoteModal) Clear() {
	m.form.Clear(true)
	m.SetBackgroundColor(tcell.ColorBlue)
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

func (m *NoteModal) Prompt(title string, label string, value string, done func(text string)) {
	lastFocus := note.App.GetFocus()

	m.Clear()
	m.SetText(title).
		AddInputText(label, value).
		AddButtons([]string{"确定", "取消"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			m.Close()
			note.App.SetFocus(lastFocus)

			if buttonIndex == 0 {
				field := m.form.GetFormItemByLabel(label).(*tview.InputField)
				done(field.GetText())
			}
		})

	// m.SetBackgroundColor(tcell.ColorDarkSlateBlue)
	m.SetBackgroundColor(tcell.ColorMediumSlateBlue)

	note.Pages.ShowPage("modal")
}
