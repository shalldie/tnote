package note

import (
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/pgavlin/femto"
	"github.com/pgavlin/femto/runtime"
)

type ViewPanel struct {
	*BasePanel
	Editor *femto.View
}

func NewViewPanel() *ViewPanel {
	p := &ViewPanel{
		BasePanel: NewBasePanel(),
	}

	p.SetTitle("详情").SetBorderPadding(0, 0, 1, 1)

	p.prepareEditor()
	p.AddItem(p.Editor, 0, 1, false)
	p.AddTip("编辑：E ; 保存：Esc", " ")
	return p
}

func (p *ViewPanel) ActivateEditor() {
	p.Editor.Readonly = false
	p.Editor.SetBorderColor(tcell.ColorDarkOrange)
	app.SetFocus(p.Editor)
}

func (p *ViewPanel) DeactivateEditor() {
	p.Editor.Readonly = true
	p.Editor.SetBorderColor(tcell.ColorLightSlateGray)
	// app.SetFocus(p)
}

var colorScheme femto.Colorscheme

func (p *ViewPanel) prepareEditor() {
	p.Editor = femto.NewView(makeBufferFromString(""))
	p.Editor.SetRuntimeFiles(runtime.Files)

	if monokai := runtime.Files.FindFile(femto.RTColorscheme, "monokai"); monokai != nil {
		if data, err := monokai.Data(); err == nil {
			colorScheme = femto.ParseColorscheme(string(data))
		}
	}

	p.Editor.SetColorscheme(colorScheme)
	p.Editor.SetBorder(true).SetBorderPadding(0, 0, 1, 1)
	p.Editor.SetBorderColor(tcell.ColorLightSlateGray)

	p.Editor.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			p.DeactivateEditor()
			p.SaveContent(p.Editor.Buf.String())
			// if p.OnSave != nil {
			// 	p.OnSave(p.Editor.Buf.String())
			// }
			p.SetFocus()
			return nil
		}

		return event
	})
}

func (p *ViewPanel) SetContent(content string) {
	p.Editor.Buf = makeBufferFromString(content)
	p.Editor.SetColorscheme(colorScheme) // 不重新设置会丢失主题样式
}

func (p *ViewPanel) SaveContent(content string) {
	file := g.Files[g.CurrentIndex]
	// file.Content = content
	go g.UpdateFile(file.FileName, content)
}

func (p *ViewPanel) LoadFile(fileName string) {
	p.DeactivateEditor()
	go func() {
		content := g.FetchFile(fileName)
		p.SetContent(content)
		app.Draw()
	}()
}

// 处理快捷键
func (p *ViewPanel) HandleShortcuts(event *tcell.EventKey) *tcell.EventKey {

	switch unicode.ToLower(event.Rune()) {
	case 'e':
		p.ActivateEditor()
		return nil
	}

	if event.Key() == tcell.KeyLeft {
		sidebar.SetFocus()
		return nil
	}

	return event
}

func makeBufferFromString(content string) *femto.Buffer {
	buff := femto.NewBufferFromString(content, "")
	buff.Settings["filetype"] = "markdown"
	buff.Settings["keepautoindent"] = true
	buff.Settings["statusline"] = false
	buff.Settings["softwrap"] = true
	buff.Settings["scrollbar"] = true

	return buff
}
