package model

var detailPrefix = "detail_"

type Detail struct {
	*Model
	Content string
}

func NewDetail() *Detail {
	t := &Detail{
		Model:   NewModel(),
		Content: "",
	}
	t.ID = detailPrefix + t.ID
	return t
}

func FindDetails(patterns ...string) []*Detail {
	patterns = append(patterns, detailPrefix)
	return findModels(NewDetail, patterns...)
}
