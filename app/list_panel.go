package app

import (
	"strings"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ListPanel[T any] struct {
	*BasePanel[T]
	List    *tview.List       // 列表组件
	NewItem *tview.InputField // 新加项组件
	Items   []*T              // 列表
	// model            *T                // 活动项
	LoadFromDB       func()            // 从db中获取数据
	AddNewItem       func(text string) // 添加新项
	OnSelectedChange func(item *T)     // 选择项改变
}

func NewListPanel[T any](title string, newItemText string) *ListPanel[T] {
	// instance
	l := &ListPanel[T]{
		BasePanel: NewBasePanel[T](),
		List: tview.NewList().ShowSecondaryText(false).SetHighlightFullLine(true).
			SetSelectedStyle(
				tcell.Style{}.Background(tcell.ColorBlue),
			),
		NewItem: makeLightTextInput(" + [" + newItemText + "] "),
	}

	// 组件
	l.SetTitle(" " + title + " ")
	l.AddItem(l.List, 0, 1, true).AddItem(l.NewItem, 1, 0, false)
	l.AddTip("新建：N ; 删除：D", "")

	// 兼容 powerlevel10k
	l.List.SetBorderPadding(0, 0, 1, 1)
	l.NewItem.SetBorderPadding(0, 0, 1, 1)

	// 事件 - list
	l.SetFocusFunc(func() {
		l.SetFocus()
	})
	// SetSelectedFunc
	l.List.SetChangedFunc(func(i int, s1, s2 string, r rune) {
		l.Model = l.Items[i]
		if l.OnSelectedChange != nil {
			// l.onSelectedChange(l.model)
			go app.QueueUpdateDraw(func() {
				l.Mu.Lock()
				defer l.Mu.Unlock()
				l.OnSelectedChange(l.Model)
				// app.Draw()
			})
			// go func() {
			// 	l.mu.Lock()
			// 	defer l.mu.Unlock()
			// 	l.onSelectedChange(l.model)
			// 	app.Draw()
			// }()
		}
	})

	// 事件 - newproject
	l.NewItem.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			l.AddNewItem(strings.TrimSpace(l.NewItem.GetText()))
			statusBar.ShowForSeconds("添加完毕...", 3)
		case tcell.KeyEsc:
			l.NewItem.SetText("")
			l.SetFocus()
		}
	})

	return l

}

// 重置数据、状态
func (l *ListPanel[T]) Reset() {
	// l.mu.Lock()
	// defer l.mu.Unlock()

	l.List.Clear()
	l.Items = make([]*T, 0)
	l.Model = nil
	if l.LoadFromDB != nil {
		l.LoadFromDB()
	}
	l.NewItem.SetText("")
}

// 设置焦点
func (l *ListPanel[T]) SetFocus() {
	app.SetFocus(l)
}

// 处理快捷键
func (l *ListPanel[T]) HandleShortcuts(event *tcell.EventKey) *tcell.EventKey {
	switch unicode.ToLower(event.Rune()) {
	// 新建
	case 'n':
		app.SetFocus(l.NewItem)
		return nil
	// 删除
	case 'd':
		app.SetFocus(l.NewItem)
		return nil
	}

	// 向左
	if event.Key() == tcell.KeyLeft && l.Prev != nil {
		l.Prev.SetFocus()
		return nil
	}
	// 向右
	if event.Key() == tcell.KeyRight && l.Next != nil {
		l.Next.SetFocus()
		return nil
	}

	return event
}
