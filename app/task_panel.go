package app

import (
	"strings"
	"unicode/utf8"

	"github.com/shalldie/gog/gs"
	"github.com/shalldie/ttm/model"
)

type TaskPanel struct {
	*ListPanel[model.Task]
}

func NewTaskPanel() *TaskPanel {
	p := &TaskPanel{
		ListPanel: NewListPanel[model.Task]("任务", "新任务"),
	}

	p.LoadFromDB = func() {
		pattern := strings.Join(projectPanel.Model.TaskIds, "|")
		tasks := []*model.Task{}
		if len(pattern) > 0 {
			tasks = model.FindTasks(pattern)
		}
		p.Items = gs.Sort(tasks, func(item1, item2 *model.Task) bool {
			return item1.CreatedTime < item2.CreatedTime
		})

		p.List.Clear()

		for _, item := range p.Items {
			p.List.AddItem(" - "+item.Name, "", 0, func() {})
		}
	}

	p.AddNewItem = func(text string) {

		if utf8.RuneCountInString(text) < 3 {
			statusBar.ShowForSeconds("任务名长度最少3个字符", 5)
			return
		}

		task := model.NewTask()
		task.Name = text

		projectPanel.Model.TaskIds = append(projectPanel.Model.TaskIds, task.ID)
		projectPanel.SaveModel()
		p.Model = task
		p.SaveModel()

		p.Reset()

		curIndex := gs.FindIndex(p.Items, func(item *model.Task, index int) bool {
			return item.ID == task.ID
		})

		p.List.SetCurrentItem(curIndex)
		p.SetFocus()
	}

	p.OnSelectedChange = func(item *model.Task) {
		detailPanel.Reset()
	}

	return p
}
