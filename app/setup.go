package app

import (
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	app             *tview.Application
	layout, content *tview.Flex
	projectPanel    *ProjectPanel
	taskPanel       *TaskPanel
	detailPanel     *DetailPanel
	statusBar       *StatusBar
)

func Setup() {
	// 应用
	app = tview.NewApplication().EnableMouse(true)

	// panel
	projectPanel = NewProjectPanel()
	taskPanel = NewTaskPanel()
	detailPanel = NewDetailPanel()

	projectPanel.Next = taskPanel
	taskPanel.Prev = projectPanel
	taskPanel.Next = detailPanel
	detailPanel.Prev = taskPanel

	// layout
	prepareLayout(projectPanel, taskPanel, detailPanel)

	// shortcuts
	setKeyboardShortcuts()

	// project p.loadFromDB()
	projectPanel.Reset()

	if err := app.SetRoot(layout, true).SetFocus(projectPanel).Run(); err != nil {
		panic(err)
	}
}

func prepareLayout(col0 tview.Primitive, col1 tview.Primitive, col2 tview.Primitive) {

	// 容器 - 上下
	layout = tview.NewFlex().SetDirection(tview.FlexRow)

	splitItem := createSplitItem()

	// 容器 - 上 - 左中右
	content = tview.NewFlex().
		AddItem(splitItem, 1, 1, false).
		AddItem(col0, 30, 1, false).
		AddItem(splitItem, 1, 1, false).
		AddItem(col1, 50, 1, false).
		AddItem(splitItem, 1, 1, false).
		AddItem(col2, 0, 1, false).
		AddItem(splitItem, 1, 1, false)

	// 容器 - 下
	statusBar = NewStatusBar()

	layout.AddItem(content, 0, 1, false).
		AddItem(
			tview.NewFlex().
				AddItem(splitItem, 2, 1, false).
				AddItem(statusBar, 0, 1, false).
				AddItem(splitItem, 2, 1, false),
			1, 1, false)

}

func setKeyboardShortcuts() *tview.Application {
	return app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if ignoreKeyEvt() {
			return event
		}

		// Global shortcuts
		switch unicode.ToLower(event.Rune()) {
		case 'p':
			app.SetFocus(projectPanel)
			// contents.RemoveItem(taskDetailPane)
			return nil
		// case 'q':
		case 't':
			app.SetFocus(taskPanel)
			// contents.RemoveItem(taskDetailPane)
			return nil

			// case 'd':
			// 	// 3s 内按 `Y` 确定删除
			// 	app.SetFocus(taskPanel)
			// 	// contents.RemoveItem(taskDetailPane)
			// 	return nil

		}

		// Handle based on current focus. Handlers may modify event
		switch {
		case projectPanel.HasFocus():
			event = projectPanel.HandleShortcuts(event)
		case taskPanel.HasFocus():
			event = taskPanel.HandleShortcuts(event)
			// 	if event != nil && projectDetailPane.isShowing() {
			// 		event = projectDetailPane.handleShortcuts(event)
			// 	}
			// case taskDetailPane.HasFocus():
			// 	event = taskDetailPane.handleShortcuts(event)
		case detailPanel.HasFocus():
			event = detailPanel.handleShortcuts(event)
			// 	if event != nil && projectDetailPane.isShowing() {
			// 		event = projectDetailPane.handleShortcuts(event)
			// 	}
			// case taskDetailPane.HasFocus():
			// 	event = taskDetailPane.handleShortcuts(event)

		}

		return event
	})
}
