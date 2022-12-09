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
		ListPanel: NewListPanel[model.Project]("项目", "新项目"),
	}

	p.LoadFromDBImpl = func() {
		p.Items = gs.Sort(model.FindProjects(), func(item1, item2 *model.Project) bool {
			return item1.CreatedTime < item2.CreatedTime
		})

		p.List.Clear()
		for _, item := range p.Items {
			p.List.AddItem(" - "+item.Name, "", 0, func() {})
		}

	}

	p.AddNewItemImpl = func(text string) {

		if utf8.RuneCountInString(text) < 2 {
			statusBar.ShowForSeconds("项目名长度最少2个字符", 5)
			return
		}

		prj := model.NewProject()
		prj.Name = text

		db.Save(prj.ID, prj)
		p.Reset()

		curIndex := gs.FindIndex(p.Items, func(item *model.Project, index int) bool {
			return item.ID == prj.ID
		})

		p.List.SetCurrentItem(curIndex)
		p.SetFocus()
	}

	p.OnSelectedChangeImpl = func(item *model.Project) {
		// p.SetTitle("loading...")
		// taskPanel.setFocus()
		p.Model = item
		taskPanel.Reset()
	}

	p.DeleteModelImpl = func(item *model.Project) {
		model.DeleteProject(item.ID)
		p.Reset()
	}

	return p
}
