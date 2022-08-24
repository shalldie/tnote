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
		ListPanel: newListPanel[model.Task]("任务", "新任务"),
	}

	p.loadFromDB = func() {
		pattern := strings.Join(projectPanel.model.TaskIds, "|")
		tasks := []*model.Task{}
		if len(pattern) > 0 {
			tasks = model.FindTasks(pattern)
		}
		p.items = gs.Sort(tasks, func(item1, item2 *model.Task) bool {
			return item1.CreatedTime < item2.CreatedTime
		})

		p.list.Clear()

		for _, item := range p.items {
			p.list.AddItem(" - "+item.Name, "", 0, func() {})
		}
	}

	p.addNewItem = func(text string) {

		if utf8.RuneCountInString(text) < 3 {
			statusBar.ShowForSeconds("任务名长度最少3个字符", 5)
			return
		}

		task := model.NewTask()
		task.Name = text

		projectPanel.model.TaskIds = append(projectPanel.model.TaskIds, task.ID)
		projectPanel.SaveModel()
		p.model = task
		p.SaveModel()

		p.reset()

		curIndex := gs.FindIndex(p.items, func(item *model.Task, index int) bool {
			return item.ID == task.ID
		})

		p.list.SetCurrentItem(curIndex)
		p.setFocus()
	}

	p.onSelectedChange = func(item *model.Task) {
		detailPanel.reset()
	}

	return p
}
