package model

import (
	"github.com/shalldie/gog/gs"
	"github.com/shalldie/ttm/db"
)

const PROJECT_PREFIX = "project_"

type Project struct {
	*Model
	Name    string
	TaskIds []string
}

func NewProject() *Project {
	p := &Project{
		Model: NewModel(),
	}
	p.ID = PROJECT_PREFIX + p.ID
	return p
}

func FindProjects(patterns ...string) []*Project {
	patterns = append(patterns, PROJECT_PREFIX)
	return findModels(NewProject, patterns...)
}

func DeleteProject(key string) {
	list := FindProjects(key)
	if len(list) <= 0 {
		return
	}
	prj := list[0]

	gs.ForEach(prj.TaskIds, func(s string, i int) {
		DeleteTask(s)
	})

	db.Delete(prj.ID)
}
