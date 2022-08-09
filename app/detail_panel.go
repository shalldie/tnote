package app

import "github.com/rivo/tview"

type DetailPanel struct {
	*tview.Box
}

func NewDetailPanel() *DetailPanel {
	return &DetailPanel{
		Box: tview.NewBox().SetBorder(true).SetTitle(" 详情 "),
	}
}
