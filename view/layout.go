package view

import "github.com/rivo/tview"

var (
	SB *StatusBar
)

var splitItem *tview.Box = tview.NewBox().SetBorder(false)

func Setup(col0 *tview.Box, col1 *tview.Box, col2 *tview.Box) {

	app := tview.NewApplication()

	layout := tview.NewFlex().SetDirection(tview.FlexRow)

	content := tview.NewFlex().
		AddItem(splitItem, 1, 1, false).
		AddItem(col0, 30, 1, false).
		AddItem(splitItem, 1, 1, false).
		AddItem(col1, 50, 1, false).
		AddItem(splitItem, 1, 1, false).
		AddItem(col2, 0, 1, false).
		AddItem(splitItem, 1, 1, false)

	SB = newStatusBar(app)

	layout.AddItem(content, 0, 1, false).
		AddItem(
			tview.NewFlex().
				AddItem(splitItem, 2, 1, false).
				AddItem(SB, 0, 1, false).
				AddItem(splitItem, 2, 1, false),
			1, 1, false)

	if err := app.SetRoot(layout, true).SetFocus(content).Run(); err != nil {
		panic(err)
	}

}
