package app

import (
	"fmt"
	"time"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/pgavlin/femto"
	"github.com/pgavlin/femto/runtime"
	"github.com/shalldie/ttm/model"
)

type DetailPanel struct {
	*BasePanel[model.Detail]
	Editor *femto.View
}

func NewDetailPanel() *DetailPanel {
	p := &DetailPanel{
		BasePanel: NewBasePanel[model.Detail](),
	}

	p.SetTitle("任务详情").SetBorderPadding(0, 0, 1, 1)

	p.PrepareEditor()
	p.AddItem(p.Editor, 0, 1, false)
	p.AddTip("编辑：E ; 保存：Esc", " ")
	// p.AddRightTip("创建时间：2022-12-08 23:22:33")

	// p.reset()

	return p
}

func (d *DetailPanel) IsReady() bool {
	return taskPanel.Model != nil
}

func (p *DetailPanel) Reset() {
	// p.deactivateEditor()
	p.Model = nil
	if !p.IsReady() {
		return
	}
	p.LoadModel()
}

func (p *DetailPanel) LoadModel() {
	p.Model = nil
	did := taskPanel.Model.DetailId

	// 查找是否存在对应的 detail
	if len(did) > 0 {
		targetList := model.FindDetails(did)
		if len(targetList) > 0 {
			p.Model = targetList[0]
		}
	}

	// 不存在则创建
	if p.Model == nil {
		p.Model = model.NewDetail()
		taskPanel.Model.DetailId = p.Model.ID
		taskPanel.SaveModel()
		p.SaveModel()
	}

	p.SetContent(p.Model.Content)
	timeContent := fmt.Sprintf(" 创建时间：%s ", time.Unix(p.Model.CreatedTime, 0).Format("2006-01-02 15:04:05"))
	p.Tips[1].SetText(timeContent)
}

func (p *DetailPanel) ActivateEditor() {
	p.Editor.Readonly = false
	p.Editor.SetBorderColor(tcell.ColorDarkOrange)
	app.SetFocus(p.Editor)
}

func (p *DetailPanel) DeactivateEditor() {
	p.Editor.Readonly = true
	p.Editor.SetBorderColor(tcell.ColorLightSlateGray)
	// app.SetFocus(p)
}

func (p *DetailPanel) SetContent(content string) {
	p.Editor.Buf = makeBufferFromString(content)
	p.Editor.SetColorscheme(colorScheme) // 不重新设置会丢失主题样式
}

var colorScheme femto.Colorscheme

func (p *DetailPanel) PrepareEditor() {
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
			// d.updateTaskNote(d.taskDetailView.Buf.String())
			p.DeactivateEditor()
			p.Model.Content = p.Editor.Buf.String()
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
	if !p.IsReady() {
		return event
	}
	switch unicode.ToLower(event.Rune()) {
	case 'e':
		p.ActivateEditor()
		return nil
	}

	if event.Key() == tcell.KeyLeft && p.Prev != nil {
		p.Prev.SetFocus()
		return nil
	}

	return event
}
