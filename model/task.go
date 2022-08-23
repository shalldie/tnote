package model

var taskPrefix = "task_"

type Task struct {
	*Model
	Name     string
	DetailId string
}

func NewTask() *Task {
	t := &Task{
		Model:    NewModel(),
		Name:     "",
		DetailId: "",
	}
	t.ID = taskPrefix + t.ID
	return t
}

func FindTasks(patterns ...string) []*Task {
	patterns = append(patterns, taskPrefix)
	return findModels(NewTask, patterns...)
}
