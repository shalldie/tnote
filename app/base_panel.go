package app

import (
	"reflect"
	"sync"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/shalldie/ttm/db"
)

type IPanel interface {
	SetFocus()
}

// panel 基类
type BasePanel[T any] struct {
	*tview.Flex
	Mu    *sync.Mutex // goroutine 的时候要加锁
	Model *T
	Prev  IPanel
	Next  IPanel
	Tips  []*tview.TextView
}

// 创建 panel 实例
func NewBasePanel[T any]() *BasePanel[T] {
	p := &BasePanel[T]{
		Flex: tview.NewFlex().SetDirection(tview.FlexRow),
		Mu:   &sync.Mutex{},
	}

	p.SetBorder(true)

	return p
}

// 设置焦点
func (p *BasePanel[T]) SetFocus() {
	app.SetFocus(p)
}

// 设置标题
func (p *BasePanel[T]) SetTitle(title string) *BasePanel[T] {
	p.Flex.SetTitle(" " + title + " ")
	return p
}

// 添加 tip
func (p *BasePanel[T]) AddTip(leftTip string, rightTip string) *BasePanel[T] {
	flexItem := tview.NewFlex()

	tipcom := tview.NewTextView().SetText(" " + leftTip + " ").SetTextColor(tcell.ColorYellow)
	flexItem.AddItem(tipcom, 0, 1, false)

	tipcomRight := tview.NewTextView().SetText(" " + rightTip + " ").SetTextColor(tcell.ColorYellow).SetTextAlign(tview.AlignRight)
	proportion := func() int {
		if utf8.RuneCountInString(rightTip) > 0 {
			return 1
		}
		return 0
	}()
	flexItem.AddItem(tipcomRight, 0, proportion, false)

	p.AddItem(flexItem, 1, 0, false)
	p.Tips = append(p.Tips, tipcom, tipcomRight)
	return p
}

// 保存当前 model 到 db
func (p *BasePanel[T]) SaveModel() *BasePanel[T] {
	id := reflect.ValueOf(*p.Model).FieldByName("ID").String()
	db.Save(id, p.Model)
	return p
}
