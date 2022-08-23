package app

import (
	"strings"
	"unicode/utf8"

	"github.com/shalldie/gog/gs"
	"github.com/shalldie/ttm/db"
	"github.com/shalldie/ttm/model"
)

type TaskPanel struct {
	*ListPanel[model.Task]
}

func NewTaskPanel() *TaskPanel {
	t := &TaskPanel{
		ListPanel: newListPanel[model.Task]("任务", "新任务"),
	}

	t.loadFromDB = func() {
		pattern := strings.Join(projectPanel.model.TaskIds, "|")
		tasks := []*model.Task{}
		if len(pattern) > 0 {
			tasks = model.FindTasks(pattern)
		}
		t.items = gs.Sort(tasks, func(item1, item2 *model.Task) bool {
			return item1.CreatedTime < item2.CreatedTime
		})
		t.list.Clear()
		for _, item := range t.items {
			t.list.AddItem(" - "+item.Name, "", 0, func() {})
		}
	}

	t.addNewItem = func(text string) {

		if utf8.RuneCountInString(text) < 3 {
			statusBar.ShowForSeconds("任务名长度最少3个字符", 5)
			return
		}

		task := model.NewTask()
		task.Name = text

		projectPanel.model.TaskIds = append(projectPanel.model.TaskIds, task.ID)
		db.Save(task.ID, task)
		db.Save(projectPanel.model.ID, projectPanel.model)

		t.reset()

		curIndex := gs.FindIndex(t.items, func(item *model.Task, index int) bool {
			return item.ID == task.ID
		})

		t.list.SetCurrentItem(curIndex)
		t.setFocus()
	}

	return t
}
