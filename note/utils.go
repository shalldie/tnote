package note

import (
	"reflect"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/shalldie/gog/gs"
)

func makeLightTextInput(placeholder string) *tview.InputField {
	return tview.NewInputField().
		SetPlaceholder(placeholder).
		SetPlaceholderTextColor(tcell.ColorLightGray).
		SetFieldTextColor(tcell.ColorBlack).
		SetFieldBackgroundColor(tcell.ColorLightBlue)
}

// 单纯占位用，兼容一些 terminal 主题
func createSplitItem() *tview.Box {
	return tview.NewBox().SetBorder(false)
}

// func confirm(text string, f func()) {

// 	activePane := app.GetFocus()
// 	modal := tview.NewModal().
// 		SetText(text).
// 		AddButtons([]string{"Yes", "No"}).
// 		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
// 			if buttonLabel == "Yes" {
// 				f()
// 			}
// 			app.SetRoot(layout, true).EnableMouse(true)
// 			app.SetFocus(activePane)
// 		})

// 	pages := tview.NewPages().
// 		AddPage("background", layout, true, true).
// 		AddPage("modal", modal, true, true)
// 	_ = app.SetRoot(pages, true).EnableMouse(true)
// }

// 是否焦点在输入框
func ignoreKeyEvt() bool {
	// if detailPanel.detailView.HasFocus() {
	// 	return true
	// }

	textInputs := []string{"*tview.InputField", "*femto.View"}
	return gs.Contains(textInputs, reflect.TypeOf(app.GetFocus()).String())
}

func makeConfirm(content string, done func()) {
	lastFocus := note.App.GetFocus()
	note.Modal.SetText(content).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			note.Pages.HidePage("modal")
			note.Pages.HidePage("main")

			go app.QueueUpdateDraw(func() {
				note.Pages.SwitchToPage("main")
				app.SetFocus(lastFocus)
				if buttonIndex == 0 {
					done()
				}
			})
		})

	note.Pages.ShowPage("modal")
}
