package model

import "github.com/shalldie/ttm/db"

var taskPrefix = "task_"

type Task struct {
	*Model
	Name     string
	DetailId string
}

func NewTask() *Task {
	t := &Task{
		Model: NewModel(),
	}
	t.ID = taskPrefix + t.ID
	return t
}

func FindTasks(patterns ...string) []*Task {
	patterns = append(patterns, taskPrefix)
	return findModels(NewTask, patterns...)
}

func DeleteTask(key string) {
	list := FindTasks(key)
	if len(list) <= 0 {
		return
	}

	task := list[0]

	DeleteDetail(task.DetailId)
	db.Delete(task.ID)
}
