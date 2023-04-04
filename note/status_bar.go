package note

// edit from https://github.com/ajaxray/geek-life

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type StatusBar struct {
	*tview.Grid
	message      *tview.TextView
	curTimestamp int64
}

func NewStatusBar() *StatusBar {
	sb := &StatusBar{
		Grid:    tview.NewGrid(),
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

func (sb *StatusBar) ShowMessage(message string) {
	app.QueueUpdateDraw(func() {
		sb.message.SetText(message)
	})
}

func (sb *StatusBar) ShowForSeconds(message string, timeout int) {

	sb.curTimestamp = time.Now().Unix()
	cts := sb.curTimestamp

	shouldStop := func() bool {
		return cts != sb.curTimestamp
	}

	go func() {
		for i := 0; i < timeout; i++ {
			if shouldStop() {
				return
			}
			sb.ShowMessage(fmt.Sprintf("%s...%ds", message, timeout-i))
			time.Sleep(time.Second)
		}
		if shouldStop() {
			return
		}
		sb.ShowMessage("")
	}()
}
