package app

import (
	"unicode/utf8"

	"github.com/shalldie/gog/gs"
	"github.com/shalldie/ttm/db"
	"github.com/shalldie/ttm/model"
)

type ProjectPanel struct {
	*ListPanel[model.Project]
}

func NewProjectPanel() *ProjectPanel {

	p := &ProjectPanel{
		ListPanel: newListPanel[model.Project]("项目", "新项目"),
	}

	p.loadFromDB = func() {
		p.items = gs.Sort(model.FindProjects(), func(item1, item2 *model.Project) bool {
			return item1.CreatedTime < item2.CreatedTime
		})

		p.list.Clear()
		for _, item := range p.items {
			p.list.AddItem(" - "+item.Name, "", 0, func() {})
		}

	}

	p.addNewItem = func(text string) {

		if utf8.RuneCountInString(text) < 3 {
			statusBar.ShowForSeconds("项目名长度最少3个字符", 5)
			return
		}

		prj := model.NewProject()
		prj.Name = text

		db.Save(prj.ID, prj)
		p.reset()

		curIndex := gs.FindIndex(p.items, func(item *model.Project, index int) bool {
			return item.ID == prj.ID
		})

		p.list.SetCurrentItem(curIndex)
		p.setFocus()
	}

	p.onSelectedChange = func(item *model.Project) {
		// p.SetTitle("loading...")
		p.activeItem = item
		taskPanel.reset()
		// taskPanel.setFocus()
	}

	return p
}
