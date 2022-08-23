package app

import (
	"strings"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ListPanel[T any] struct {
	*BasePanel[T]
	list    *tview.List       // 列表组件
	newItem *tview.InputField // 新加项组件
	items   []*T              // 列表
	// model            *T                // 活动项
	loadFromDB       func()            // 从db中获取数据
	addNewItem       func(text string) // 添加新项
	onSelectedChange func(item *T)     // 选择项改变
}

func newListPanel[T any](title string, newItemText string) *ListPanel[T] {
	// instance
	l := &ListPanel[T]{
		BasePanel: newBasePanel[T](),
		list: tview.NewList().ShowSecondaryText(false).SetHighlightFullLine(true).
			SetSelectedStyle(
				tcell.Style{}.Background(tcell.ColorBlue),
			),
		newItem: makeLightTextInput(" + [" + newItemText + "] "),
	}

	// 组件
	l.SetTitle(" " + title + " ")
	l.AddItem(l.list, 0, 1, true).AddItem(l.newItem, 1, 0, false)
	l.AddTip("新建：N ; 删除：D")

	// 兼容 powerlevel10k
	l.list.SetBorderPadding(0, 0, 1, 1)
	l.newItem.SetBorderPadding(0, 0, 1, 1)

	// 事件 - list
	l.SetFocusFunc(func() {
		l.setFocus()
	})
	// SetSelectedFunc
	l.list.SetChangedFunc(func(i int, s1, s2 string, r rune) {
		l.model = l.items[i]
		if l.onSelectedChange != nil {
			go func() {
				l.onSelectedChange(l.model)
				app.Draw()
			}()
		}
	})

	// 事件 - newproject
	l.newItem.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			l.addNewItem(strings.TrimSpace(l.newItem.GetText()))
			statusBar.ShowForSeconds("添加完毕...", 3)
		case tcell.KeyEsc:
			l.newItem.SetText("")
			l.setFocus()
		}
	})

	return l

}

// 重置数据、状态
func (l *ListPanel[T]) reset() {
	l.list.Clear()
	l.items = make([]*T, 0)
	l.model = nil
	if l.loadFromDB != nil {
		l.loadFromDB()
	}
	l.newItem.SetText("")
}

// 设置焦点
func (l *ListPanel[T]) setFocus() {
	app.SetFocus(l)
}

// 处理快捷键
func (l *ListPanel[T]) handleShortcuts(event *tcell.EventKey) *tcell.EventKey {
	switch unicode.ToLower(event.Rune()) {
	// 新建
	case 'n':
		app.SetFocus(l.newItem)
		return nil
	// 删除
	case 'd':
		app.SetFocus(l.newItem)
		return nil
	}

	// 向左
	if event.Key() == tcell.KeyLeft && l.prev != nil {
		l.prev.SetFocus()
		return nil
	}
	// 向右
	if event.Key() == tcell.KeyRight && l.next != nil {
		l.next.SetFocus()
		return nil
	}

	return event
}
