package model

var projectPrefix = "project_"

type Project struct {
	*Model
	TaskIds []string
}

func NewProject() *Project {
	p := &Project{
		Model:   NewModel(),
		TaskIds: []string{},
	}
	p.ID = projectPrefix + p.ID
	return p
}

func FindProjects(patterns ...string) []*Project {
	patterns = append(patterns, projectPrefix)
	return findModels(NewProject, patterns...)
}
