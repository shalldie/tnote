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

// 是否焦点在输入框
func ignoreKeyEvt() bool {
	// if detailPanel.detailView.HasFocus() {
	// 	return true
	// }

	textInputs := []string{"*tview.InputField", "*femto.View"}
	return gs.Contains(textInputs, reflect.TypeOf(note.App.GetFocus()).String())
}
