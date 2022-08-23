package app

import (
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/pgavlin/femto"
	"github.com/pgavlin/femto/runtime"
	"github.com/shalldie/ttm/db"
	"github.com/shalldie/ttm/model"
)

type DetailPanel struct {
	*BasePanel[model.Detail]
	editor *femto.View
}

func NewDetailPanel() *DetailPanel {
	p := &DetailPanel{
		BasePanel: newBasePanel[model.Detail](),
	}

	p.SetTitle("任务详情").SetBorderPadding(0, 0, 1, 1)

	p.prepareEditor()
	p.AddItem(p.editor, 0, 1, false)
	p.AddTip("编辑：E ; 保存：Esc")

	p.reset()

	return p
}

func (d *DetailPanel) isReady() bool {
	return taskPanel.model != nil
}

func (d *DetailPanel) reset() {
	d.deactivateEditor()
	if !d.isReady() {
		return
	}
	d.loadModel()
}

func (d *DetailPanel) loadModel() {
	did := taskPanel.model.DetailId
	if len(did) <= 0 {
		d.model = model.NewDetail()
		taskPanel.model.DetailId = did
		db.Save(taskPanel.model.ID, taskPanel.model)
		db.Save(d.model.ID, d.model)
	} else {
		d.model = model.FindDetails(did)[0]
	}
	d.editor.Buf = makeBufferFromString(d.model.Content)
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
	if !d.isReady() {
		return event
	}
	switch unicode.ToLower(event.Rune()) {
	case 'e':
		d.activateEditor()
		return nil
	}

	if event.Key() == tcell.KeyLeft && d.prev != nil {
		d.prev.SetFocus()
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
