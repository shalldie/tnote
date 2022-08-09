package view

// edit from https://github.com/ajaxray/geek-life

import (
	"time"

	"github.com/rivo/tview"
)

type StatusBar struct {
	*tview.Pages
	app     *tview.Application
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

func newStatusBar(app *tview.Application) *StatusBar {
	sb := &StatusBar{
		Pages:   tview.NewPages(),
		app:     app,
		message: tview.NewTextView().SetDynamicColors(true).SetText("Loading..."),
	}

	sb.AddPage(messagePage, sb.message, true, true)
	sb.AddPage(defaultPage,
		tview.NewGrid(). // Content will not be modified, So, no need to declare explicitly
					SetColumns(0, 0, 0, 0).
					SetRows(0).
					AddItem(tview.NewTextView().SetText("方向：↑ ↓"), 0, 0, 1, 1, 0, 0, false). // ↑ ↓ ← →
					AddItem(tview.NewTextView().SetText("新建：N").SetTextAlign(tview.AlignCenter), 0, 1, 1, 1, 0, 0, false).
					AddItem(tview.NewTextView().SetText("上一步：Esc").SetTextAlign(tview.AlignCenter), 0, 2, 1, 1, 0, 0, false).
					AddItem(tview.NewTextView().SetText("退出：Ctrl + C").SetTextAlign(tview.AlignRight), 0, 3, 1, 1, 0, 0, false),
		true,
		true,
	)

	return sb
}

func (fo *StatusBar) Restore() {
	fo.app.QueueUpdateDraw(func() {
		fo.SwitchToPage(defaultPage)
	})
}

func (fo *StatusBar) ShowForSeconds(message string, timeout int) {
	if fo.app == nil {
		return
	}

	fo.message.SetText(message)
	fo.SwitchToPage(messagePage)
	restorInQ++

	go func() {
		time.Sleep(time.Second * time.Duration(timeout))

		// Apply restore only if this is the last pending restore
		if restorInQ == 1 {
			fo.Restore()
		}
		restorInQ--
	}()
}
