package note

import (
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type IPanel interface {
	SetFocus()
	HandleShortcuts(event *tcell.EventKey) *tcell.EventKey
}

type BasePanel struct {
	*tview.Flex
	Tips []*tview.TextView
}

func NewBasePanel() *BasePanel {
	p := &BasePanel{
		Flex: tview.NewFlex().SetDirection(tview.FlexRow),
	}

	p.SetBorder(true)

	return p
}

func (p *BasePanel) SetFocus() {
	app.SetFocus(p)
}

func (p *BasePanel) SetTitle(title string) *BasePanel {
	p.Flex.SetTitle(" " + title + " ")
	return p
}

// 添加 tip
func (p *BasePanel) AddTip(leftTip string, rightTip string) *BasePanel {
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
