package note

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/shalldie/tnote/gist"
)

// 放全局，方便引用，，，反正是单例
var (
	note *TNote
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
		Gist:      gist.NewGist(token),
		App:       tview.NewApplication().EnableMouse(true),
		Pages:     tview.NewPages(),
		Modal:     tview.NewModal().AddButtons([]string{"确定", "取消"}),
		Layout:    tview.NewFlex().SetDirection(tview.FlexRow),
		Sidebar:   NewSidebarPanel(),
		View:      NewViewPanel(),
		StatusBar: NewStatusBar(),
	}

	return note
}

func (t *TNote) Setup() {

	t.initLayout()
	t.setKeyboardShortcuts()

	t.Sidebar.Setup()

	if err := t.App.SetRoot(t.Pages, true).SetFocus(t.Sidebar).Run(); err != nil {
		panic(err)
	}
}

// 初始化布局
func (t *TNote) initLayout() {
	// pages
	t.Pages.
		AddPage("main", t.Layout, true, true).
		AddPage("modal", t.Modal, true, false)

	splitItem := createSplitItem()

	// layout - 上 - 左中右
	content := tview.NewFlex().
		AddItem(splitItem, 1, 1, false).
		AddItem(t.Sidebar, 36, 1, false).
		AddItem(splitItem, 1, 1, false).
		AddItem(t.View, 0, 1, false).
		AddItem(splitItem, 1, 1, false)

	// layout - 下
	t.Layout.AddItem(content, 0, 1, false).
		AddItem(
			tview.NewFlex().
				AddItem(splitItem, 2, 1, false).
				AddItem(t.StatusBar, 0, 1, false).
				AddItem(splitItem, 2, 1, false),
			1, 1, false)
}

// 设置快捷键
func (t *TNote) setKeyboardShortcuts() *tview.Application {
	return t.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
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
