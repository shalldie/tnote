package app

import (
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/pgavlin/femto"
	"github.com/pgavlin/femto/runtime"
	"github.com/rivo/tview"
)

type DetailPanel struct {
	*tview.Flex
	parent    IListPanel // 上一个panel
	editorTip *tview.TextView
	editor    *femto.View
}

func NewDetailPanel() *DetailPanel {
	detail := &DetailPanel{
		Flex: tview.NewFlex().SetDirection(tview.FlexRow),
	}

	detail.SetBorder(true).SetTitle(" 任务详情 ").SetBorderPadding(0, 0, 1, 1)

	detail.editorTip = tview.NewTextView().SetText(" 编辑：E ; 保存：Esc").SetTextColor(tcell.ColorYellow)

	detail.prepareEditor()

	detail.AddItem(detail.editor, 0, 1, false)
	detail.AddItem(detail.editorTip, 1, 0, false)

	detail.reset()

	return detail
}

func (d *DetailPanel) reset() {
	d.deactivateEditor()

	d.editor.Buf = makeBufferFromString("some thing new")
	// d.detailView
	// d.editor.set
}

func (d *DetailPanel) activateEditor() {
	d.editor.Readonly = false
	d.editor.SetBorderColor(tcell.ColorDarkOrange)
	app.SetFocus(d.editor)
}

func (d *DetailPanel) deactivateEditor() {
	d.editor.Readonly = true
	d.editor.SetBorderColor(tcell.ColorLightSlateGray)
	app.SetFocus(d)
}

func (d *DetailPanel) prepareEditor() {
	d.editor = femto.NewView(makeBufferFromString(""))

	var colorScheme femto.Colorscheme
	if monokai := runtime.Files.FindFile(femto.RTColorscheme, "monokai"); monokai != nil {
		if data, err := monokai.Data(); err == nil {
			colorScheme = femto.ParseColorscheme(string(data))
		}
	}

	d.editor.SetColorscheme(colorScheme)
	d.editor.SetBorder(true).SetBorderPadding(0, 0, 1, 1)
	d.editor.SetBorderColor(tcell.ColorLightSlateGray)

	d.editor.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			// d.updateTaskNote(d.taskDetailView.Buf.String())
			d.deactivateEditor()
			return nil
		}

		return event
	})
}

func (d *DetailPanel) setFocus() {
	app.SetFocus(d)
}

func makeBufferFromString(content string) *femto.Buffer {
	buff := femto.NewBufferFromString(content, "")
	// taskDetail.Settings["ruler"] = false
	buff.Settings["filetype"] = "markdown"
	buff.Settings["keepautoindent"] = true
	buff.Settings["statusline"] = false
	buff.Settings["softwrap"] = true
	buff.Settings["scrollbar"] = true

	return buff
}

// 处理快捷键
func (d *DetailPanel) handleShortcuts(event *tcell.EventKey) *tcell.EventKey {
	switch unicode.ToLower(event.Rune()) {
	case 'e':
		d.activateEditor()
		return nil
	}

	if event.Key() == tcell.KeyLeft && d.parent != nil {
		d.parent.setFocus()
		return nil
	}

	if event.Key() == tcell.KeyESC {
		d.deactivateEditor()
		return nil
	}

	// if event.Key() == tcell.KeyESC && l.parent != nil {
	// 	l.parent.setFocus()
	// } else if event.Key() == tcell.KeyEnter && l.child != nil {
	// 	l.child.setFocus()
	// }

	return event
}
