package note

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/shalldie/gog/gs"
	"github.com/shalldie/tnote/gist"
)

type SidebarPanel struct {
	*BasePanel
	List    *tview.List       // 列表组件
	NewItem *tview.InputField // 新加项组件
}

func NewSidebarPanel() *SidebarPanel {
	p := &SidebarPanel{
		BasePanel: NewBasePanel(),
		List: tview.NewList().ShowSecondaryText(false).SetHighlightFullLine(true).
			SetSelectedStyle(
				tcell.Style{}.Background(tcell.ColorBlue),
			),
		NewItem: makeLightTextInput(" + [新建文件] "),
	}

	// 组件
	p.SetTitle("文件")
	p.AddItem(p.List, 0, 1, true).AddItem(p.NewItem, 1, 0, false)
	p.AddTip(strings.Join([]string{
		"新建：N",
		"编辑：E",
		"删除：D",
	}, " ; "), "")

	// 兼容 powerlevel10k
	p.List.SetBorderPadding(0, 0, 1, 1)
	p.NewItem.SetBorderPadding(0, 0, 1, 1)

	// 事件 - list
	p.SetFocusFunc(func() {
		p.SetFocus()
	})

	// SetSelectedFunc
	p.List.SetChangedFunc(func(i int, s1, s2 string, r rune) {
		// view.LoadFile(p.Files[i])
		g.CurrentIndex = i
		view.SetContent(g.GetContent())
	})

	// 事件 - newproject
	p.NewItem.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			p.AddFile()
			// go func() {
			// g.UpdateFile(strings.TrimSpace(p.NewItem.GetText()), "To be edited.")
			// p.LoadFiles()
			// }()
			// p.addFile(strings.TrimSpace(p.NewItem.GetText()))
			// l.AddNewItemImpl(strings.TrimSpace(l.NewItem.GetText()))
			// statusBar.ShowForSeconds("添加完毕...", 3)
		case tcell.KeyEsc:
			p.NewItem.SetText("")
			p.SetFocus()
		}
	})

	return p
}

func (p *SidebarPanel) Setup() {
	go func() {
		g.Setup()
		p.LoadFiles()
		app.Draw()
	}()
}

func (p *SidebarPanel) LoadFiles() {
	p.List.Clear()

	for _, f := range g.Files {
		p.List.AddItem(" - "+f.FileName, "", 0, nil)
	}

	// if len(p.Files) > 0 {
	// 	view.LoadFile(p.Files[0])
	// }

	// app.Draw()
}

func (p *SidebarPanel) AddFile() {
	fileName := strings.TrimSpace(p.NewItem.GetText())

	if utf8.RuneCountInString(fileName) < 2 {
		note.StatusBar.ShowForSeconds("文件名长度最少2个字符", 5)
		return
	}

	g.UpdateFile(fileName, "To be edited.")
	p.LoadFiles()

	curIndex := gs.FindIndex(g.Files, func(item *gist.FileModel, index int) bool {
		return item.FileName == fileName
	})

	g.CurrentIndex = curIndex
	p.List.SetCurrentItem(curIndex)

	p.NewItem.SetText("")
	p.SetFocus()
}

// 处理快捷键
func (p *SidebarPanel) HandleShortcuts(event *tcell.EventKey) *tcell.EventKey {

	switch unicode.ToLower(event.Rune()) {
	// 新建
	case 'n':
		app.SetFocus(p.NewItem)
		return nil
	// 删除
	case 'd':
		// app.SetFocus(l.NewItem)
		// if l.DeleteModelImpl != nil {
		// 	l.DeleteModelImpl(l.Model)
		// }
		file := g.Files[g.CurrentIndex]
		makeConfirm(fmt.Sprintf("确定要删除【%s】吗？", file.FileName), func() {
			g.UpdateFile(file.FileName, nil)
			p.LoadFiles()
		})
		return nil
	}

	// 向左
	if event.Key() == tcell.KeyLeft {
		// l.Prev.SetFocus()
		return nil
	}
	// 向右
	if event.Key() == tcell.KeyRight {
		// l.Next.SetFocus()
		view.SetFocus()
		return nil
	}

	return event
}
