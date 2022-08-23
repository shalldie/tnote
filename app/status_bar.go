package app

// edit from https://github.com/ajaxray/geek-life

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type StatusBar struct {
	*tview.Grid
	message *tview.TextView
}

// Name of page keys
const (
	defaultPage = "default"
	messagePage = "message"
)

// Used to skip queued restore of statusBar
// in case of new showForSeconds within waiting period
var restorInQ = 0

func newStatusBar() *StatusBar {
	sb := &StatusBar{
		Grid: tview.NewGrid(),
		// Pages:   tview.NewPages(),
		message: tview.NewTextView().SetDynamicColors(true).SetTextColor(tcell.ColorYellow),
	}

	sb.SetColumns(0, 0, 0, 0).
		SetRows(0).
		AddItem(sb.message, 0, 0, 1, 1, 0, 0, false).
		AddItem(tview.NewTextView().SetText("方向：↑↓←→ ；退出：Ctrl + C").SetTextAlign(tview.AlignRight), 0, 3, 1, 1, 0, 0, false) // ↑ ↓ ← →

	sb.message.SetChangedFunc(func() {
		app.Draw()
	})

	return sb
}

func (sb *StatusBar) Restore() {
	sb.ShowMessage("")
}

func (sb *StatusBar) ShowMessage(message string) {
	sb.message.SetText(message)
}

func (sb *StatusBar) ShowForSeconds(message string, timeout int) {

	sb.ShowMessage(message)
	restorInQ++

	go func() {
		time.Sleep(time.Second * time.Duration(timeout))

		// Apply restore only if this is the last pending restore
		if restorInQ == 1 {
			sb.Restore()
		}
		restorInQ--
	}()
}
