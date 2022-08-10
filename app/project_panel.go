package app

import (
	"strings"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/shalldie/ttm/db"
	"github.com/shalldie/ttm/model"
)

type ProjectPanel struct {
	*tview.Flex
	list          *tview.List
	newProject    *tview.InputField
	projects      []*model.Project
	activeProject *model.Project
}

func NewProjectPanel() *ProjectPanel {

	tcell.Style{}.Underline(true).Background(tcell.ColorBlue)
	p := &ProjectPanel{
		Flex: tview.NewFlex().SetDirection(tview.FlexRow),
		list: tview.NewList().ShowSecondaryText(false).SetHighlightFullLine(true).
			SetSelectedStyle(
				tcell.Style{}.Background(tcell.ColorBlue),
			),
		newProject: makeLightTextInput(" + [新项目] "),
	}

	// 组件
	p.SetBorder(true).SetTitle(" 项目 ")
	p.AddItem(p.list, 0, 1, true).AddItem(p.newProject, 1, 0, false)
	// 兼容 powerlevel10k
	p.list.SetBorderPadding(0, 0, 1, 1)
	p.newProject.SetBorderPadding(0, 0, 1, 1)

	// 事件 - list
	p.SetFocusFunc(func() {
		app.SetFocus(p.list)
	})
	p.list.SetChangedFunc(func(i int, s1, s2 string, r rune) {
		p.activeProject = p.projects[i]
		p.SetTitle(p.activeProject.Name)
	})

	// 事件 - newproject
	p.newProject.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			p.addNewProject()
		case tcell.KeyEsc:
			app.SetFocus(p)
		}
	})

	// 数据
	p.loadFromDB()

	return p
}

// 从 db 更新数据
func (p *ProjectPanel) loadFromDB() {
	p.projects = model.FindProjects()
	p.list.Clear()
	for _, item := range p.projects {
		p.list.AddItem(" - "+item.Name, "", 0, func() {})
	}
}

func (p *ProjectPanel) addNewProject() {

	name := strings.TrimSpace(p.newProject.GetText())

	if len(name) < 3 {
		statusBar.ShowForSeconds("项目名长度最少3个字符", 5)
		return
	}

	prj := model.NewProject()
	prj.Name = name

	db.Save(prj.ID, prj)
	p.loadFromDB()
	p.newProject.SetText("")
	app.SetFocus(p)
}

func (p *ProjectPanel) handleShortcuts(event *tcell.EventKey) *tcell.EventKey {
	switch unicode.ToLower(event.Rune()) {
	case 'j':
		p.list.SetCurrentItem(p.list.GetCurrentItem() + 1)
		return nil
	case 'k':
		p.list.SetCurrentItem(p.list.GetCurrentItem() - 1)
		return nil
	case 'n':
		app.SetFocus(p.newProject)
		return nil
	}

	return event
}
