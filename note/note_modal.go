package note

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
	// m.ClearButtons()
	m.form.Clear(true)
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

func (m *NoteModal) Prompt(title string, label string, value string, done func()) {
	lastFocus := note.App.GetFocus()

	m.Clear()
	m.SetText(title).
		AddInputText(label, value).
		AddButtons([]string{"确定", "取消"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			m.Close()
			note.App.SetFocus(lastFocus)

			if buttonIndex == 0 {
				done()
			}
		})

	m.form.GetFormItemByLabel(label)

	note.Pages.ShowPage("modal")
}
