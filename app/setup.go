package app

import "github.com/rivo/tview"

var (
	app          *tview.Application
	projectPanel *ProjectPanel
	taskPanel    *TaskPanel
	detailPanel  *DetailPanel
	statusBar    *StatusBar
)

func Setup() {
	projectPanel = NewProjectPanel()
	taskPanel = NewTaskPanel()
	detailPanel = NewDetailPanel()

	prepareLayout(projectPanel, taskPanel, detailPanel)
}

func prepareLayout(col0 tview.Primitive, col1 tview.Primitive, col2 tview.Primitive) {

	app = tview.NewApplication().EnableMouse(true)

	layout := tview.NewFlex().SetDirection(tview.FlexRow)

	splitItem := createSplitItem()

	content := tview.NewFlex().
		AddItem(splitItem, 1, 1, false).
		AddItem(col0, 30, 1, false).
		AddItem(splitItem, 1, 1, false).
		AddItem(col1, 50, 1, false).
		AddItem(splitItem, 1, 1, false).
		AddItem(col2, 0, 1, false).
		AddItem(splitItem, 1, 1, false)

	statusBar = newStatusBar(app)

	layout.AddItem(content, 0, 1, false).
		AddItem(
			tview.NewFlex().
				AddItem(splitItem, 2, 1, false).
				AddItem(statusBar, 0, 1, false).
				AddItem(splitItem, 2, 1, false),
			1, 1, false)

	if err := app.SetRoot(layout, true).SetFocus(content).Run(); err != nil {
		panic(err)
	}

}
