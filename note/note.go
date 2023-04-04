package note

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/shalldie/tnote/gist"
)

// 放全局，方便引用，，，反正是单例
var (
	note    *TNote
	app     *tview.Application
	g       *gist.Gist
	sidebar *SidebarPanel
	view    *ViewPanel
)

type TNote struct {
	Gist      *gist.Gist
	App       *tview.Application // 应用
	Pages     *tview.Pages       // pages
	Modal     *tview.Modal       // page - 弹框
	Layout    *tview.Flex        // page - 主容器
	Sidebar   *SidebarPanel      // 侧边栏
	View      *ViewPanel         // 右边视图
	StatusBar *StatusBar         // 状态栏
}

func NewTNote(token string) *TNote {
	note = &TNote{
		Gist: gist.NewGist(token),
	}
	g = note.Gist
	return note
}

func (t *TNote) Setup() {

	fmt.Println("loading...")

	t.initLayout()
	t.setKeyboardShortcuts()

	t.Sidebar.Setup()

	if err := t.App.SetRoot(t.Pages, true).SetFocus(t.Sidebar).Run(); err != nil {
		panic(err)
	}
}

func (t *TNote) initLayout() {
	// app
	app = tview.NewApplication().EnableMouse(true)
	t.App = app

	// pages
	t.Layout = tview.NewFlex().SetDirection(tview.FlexRow)
	t.Modal = tview.NewModal().AddButtons([]string{"确定", "取消"})

	t.Pages = tview.NewPages().
		AddPage("main", t.Layout, true, true).
		AddPage("modal", t.Modal, true, false)

	splitItem := createSplitItem()

	// layout - 上 - 左中右
	t.Sidebar = NewSidebarPanel()
	sidebar = t.Sidebar
	t.View = NewViewPanel()
	view = t.View
	content := tview.NewFlex().
		AddItem(splitItem, 1, 1, false).
		AddItem(t.Sidebar, 36, 1, false).
		AddItem(splitItem, 1, 1, false).
		AddItem(t.View, 0, 1, false).
		AddItem(splitItem, 1, 1, false)

	//
	// layout - 下
	t.StatusBar = NewStatusBar()

	t.Layout.AddItem(content, 0, 1, false).
		AddItem(
			tview.NewFlex().
				AddItem(splitItem, 2, 1, false).
				AddItem(t.StatusBar, 0, 1, false).
				AddItem(splitItem, 2, 1, false),
			1, 1, false)
}

func (t *TNote) setKeyboardShortcuts() *tview.Application {
	return app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if ignoreKeyEvt() {
			return event
		}

		// Global shortcuts
		// switch unicode.ToLower(event.Rune()) {
		// case '1':
		// 	app.SetFocus(projectPanel)
		// 	return nil
		// case '2':
		// 	app.SetFocus(taskPanel)
		// 	return nil

		// case '3':
		// 	app.SetFocus(detailPanel)
		// 	return nil
		// }

		// Handle based on current focus. Handlers may modify event
		switch {
		case t.Sidebar.HasFocus():
			event = t.Sidebar.HandleShortcuts(event)
		case t.View.HasFocus():
			event = t.View.HandleShortcuts(event)
			// 	if event != nil && projectDetailPane.isShowing() {
			// 		event = projectDetailPane.handleShortcuts(event)
			// 	}
			// case taskDetailPane.HasFocus():
			// 	event = taskDetailPane.handleShortcuts(event)

		}

		return event
	})
}
