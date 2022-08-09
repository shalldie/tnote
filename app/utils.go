package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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
