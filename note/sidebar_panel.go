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
		"重命名：R",
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
		if note.Gist.Model == nil {
			return
		}
		note.Gist.CurrentIndex = i
		note.View.LoadFile(note.Gist.GetFile())
	})

	// 事件 - newproject
	p.NewItem.SetFocusFunc(func() {
		note.StatusBar.ShowMessage("新建文件中...")
	})
	p.NewItem.SetBlurFunc(func() {
		note.StatusBar.ShowMessage("")
	})
	p.NewItem.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			p.AddFile(p.NewItem.GetText())
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
		note.StatusBar.ShowMessage("加载中...")
		note.Gist.Setup()
		p.LoadFiles()
		note.App.Draw()
		note.StatusBar.ShowMessage("")
	}()
}

func (p *SidebarPanel) LoadFiles() {
	p.List.Clear()

	for _, f := range note.Gist.Files {
		p.List.AddItem(" - "+f.FileName, "", 0, nil)
	}
}

func (p *SidebarPanel) AddFile(fileName string) {
	fileName = strings.TrimSpace(fileName)

	if utf8.RuneCountInString(fileName) < 2 {
		note.StatusBar.ShowForSeconds("文件名长度最少2个字符", 3)
		return
	}

	note.Gist.UpdateFile(fileName, "To be edited.")
	p.LoadFiles()

	curIndex := gs.FindIndex(note.Gist.Files, func(item *gist.GistFile, index int) bool {
		return item.FileName == fileName
	})

	note.Gist.CurrentIndex = curIndex
	p.List.SetCurrentItem(curIndex)

	p.NewItem.SetText("")
	p.SetFocus()
}

// 处理快捷键
func (p *SidebarPanel) HandleShortcuts(event *tcell.EventKey) *tcell.EventKey {

	switch unicode.ToLower(event.Rune()) {
	// 新建
	case 'n':
		note.App.SetFocus(p.NewItem)
		return nil
	// 删除
	case 'd':
		file := note.Gist.Files[note.Gist.CurrentIndex]
		note.Modal.Confirm(fmt.Sprintf("确定要删除【%s】吗？", file.FileName), func() {
			go func() {
				note.StatusBar.ShowMessage("删除中...")
				note.Gist.UpdateFile(file.FileName, nil)
				p.LoadFiles()
				note.StatusBar.ShowForSeconds("操作成功", 3)
			}()
		})
		return nil
	}

	// 向左
	if event.Key() == tcell.KeyLeft {
		return nil
	}
	// 向右
	if event.Key() == tcell.KeyRight {
		note.View.SetFocus()
		return nil
	}

	return event
}
