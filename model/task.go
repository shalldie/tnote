package model

var taskPrefix = "task_"

type Task struct {
	*Model
}

func NewTask() *Task {
	t := &Task{
		NewModel(),
	}
	t.ID = taskPrefix + t.ID
	return t
}

func FindTasks(patterns ...string) []*Task {
	patterns = append(patterns, taskPrefix)
	return findModels(NewTask, patterns...)
}
