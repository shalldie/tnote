package app

import (
	"reflect"
	"sync"

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
func (p *BasePanel[T]) AddTip(tip string) *BasePanel[T] {
	tipcom := tview.NewTextView().SetText(" " + tip + " ").SetTextColor(tcell.ColorYellow)
	p.AddItem(tipcom, 1, 0, false)
	return p
}

// 保存当前 model 到 db
func (p *BasePanel[T]) SaveModel() *BasePanel[T] {
	id := reflect.ValueOf(*p.Model).FieldByName("ID").String()
	db.Save(id, p.Model)
	return p
}
