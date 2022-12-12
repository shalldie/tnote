package model

import (
	"github.com/shalldie/ttm/db"
)

const DETAIL_PREFIX = "detail_"

type Detail struct {
	*Model
	Content string
}

func NewDetail() *Detail {
	t := &Detail{
		Model: NewModel(),
	}
	t.ID = DETAIL_PREFIX + t.ID
	return t
}

func FindDetails(patterns ...string) []*Detail {
	patterns = append(patterns, DETAIL_PREFIX)
	return findModels(NewDetail, patterns...)
}

func DeleteDetail(key string) {
	db.Delete(key)
}
