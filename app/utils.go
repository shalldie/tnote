package app

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

func confirm(text string, f func()) {

	activePane := app.GetFocus()
	modal := tview.NewModal().
		SetText(text).
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				f()
			}
			app.SetRoot(layout, true).EnableMouse(true)
			app.SetFocus(activePane)
		})

	pages := tview.NewPages().
		AddPage("background", layout, true, true).
		AddPage("modal", modal, true, true)
	_ = app.SetRoot(pages, true).EnableMouse(true)
}

// 是否焦点在输入框
func ignoreKeyEvt() bool {
	textInputs := []string{"*tview.InputField", "*femto.View"}
	return gs.Contains(textInputs, reflect.TypeOf(app.GetFocus()).String())
}
