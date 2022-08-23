package app

import (
	"reflect"

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
	model *T
	prev  IPanel
	next  IPanel
}

// 创建 panel 实例
func newBasePanel[T any]() *BasePanel[T] {
	p := &BasePanel[T]{
		Flex: tview.NewFlex().SetDirection(tview.FlexRow),
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
	id := reflect.ValueOf(*p.model).FieldByName("ID").String()
	db.Save(id, p.model)
	return p
}
