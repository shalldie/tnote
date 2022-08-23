package app

import (
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/pgavlin/femto"
	"github.com/pgavlin/femto/runtime"
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

	// p.reset()

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

func (p *DetailPanel) loadModel() {
	did := taskPanel.model.DetailId
	if len(did) <= 0 {
		p.model = model.NewDetail()
		taskPanel.model.DetailId = p.model.ID
		// go func() {
		taskPanel.SaveModel()
		p.SaveModel()
		// }()
	} else {
		p.model = model.FindDetails(did)[0]
	}
	p.editor.Buf = makeBufferFromString(p.model.Content)
}

func (p *DetailPanel) activateEditor() {
	p.editor.Readonly = false
	p.editor.SetBorderColor(tcell.ColorDarkOrange)
	app.SetFocus(p.editor)
}

func (p *DetailPanel) deactivateEditor() {
	p.editor.Readonly = true
	p.editor.SetBorderColor(tcell.ColorLightSlateGray)
	// app.SetFocus(p)
}

func (p *DetailPanel) prepareEditor() {
	p.editor = femto.NewView(makeBufferFromString(""))

	var colorScheme femto.Colorscheme
	if monokai := runtime.Files.FindFile(femto.RTColorscheme, "monokai"); monokai != nil {
		if data, err := monokai.Data(); err == nil {
			colorScheme = femto.ParseColorscheme(string(data))
		}
	}

	p.editor.SetColorscheme(colorScheme)
	p.editor.SetBorder(true).SetBorderPadding(0, 0, 1, 1)
	p.editor.SetBorderColor(tcell.ColorLightSlateGray)

	p.editor.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			// d.updateTaskNote(d.taskDetailView.Buf.String())
			p.deactivateEditor()
			p.model.Content = p.editor.Buf.String()
			p.SaveModel()
			p.SetFocus()
			return nil
		}

		return event
	})
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
func (p *DetailPanel) handleShortcuts(event *tcell.EventKey) *tcell.EventKey {
	if !p.isReady() {
		return event
	}
	switch unicode.ToLower(event.Rune()) {
	case 'e':
		p.activateEditor()
		return nil
	}

	if event.Key() == tcell.KeyLeft && p.prev != nil {
		p.prev.SetFocus()
		return nil
	}

	return event
}
