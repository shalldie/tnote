package model

var projectPrefix = "project_"

type Project struct {
	*Model
}

func NewProject() *Project {
	p := &Project{
		NewModel(),
	}
	p.ID = projectPrefix + p.ID
	return p
}

func FindProjects(patterns ...string) []*Project {
	patterns = append(patterns, projectPrefix)
	return findModels(NewProject, patterns...)
}
