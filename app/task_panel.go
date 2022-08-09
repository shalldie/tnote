package app

import "github.com/rivo/tview"

type TaskPanel struct {
	*tview.Box
}

func NewTaskPanel() *TaskPanel {
	return &TaskPanel{
		Box: tview.NewBox().SetBorder(true).SetTitle(" 任务 "),
	}
}
