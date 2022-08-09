package app

import "github.com/rivo/tview"

type UI struct {
	ui *tview.Box
}

func newUI(title string) *UI {
	return &UI{
		ui: tview.NewBox().SetBorder(true).SetTitle(title),
	}
}
