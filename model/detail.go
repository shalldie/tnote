package model

import (
	"github.com/shalldie/ttm/db"
)

var detailPrefix = "detail_"

type Detail struct {
	*Model
	Content string
}

func NewDetail() *Detail {
	t := &Detail{
		Model: NewModel(),
	}
	t.ID = detailPrefix + t.ID
	return t
}

func FindDetails(patterns ...string) []*Detail {
	patterns = append(patterns, detailPrefix)
	return findModels(NewDetail, patterns...)
}

func DeleteDetail(key string) {
	db.Delete(key)
}
