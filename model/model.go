package model

import (
	"github.com/google/uuid"
	"github.com/shalldie/gog/gs"
	"github.com/shalldie/ttm/db"
)

type Model struct {
	ID   string
	Name string
}

func NewModel() *Model {
	return &Model{
		ID: uuid.NewString(),
	}
}

func findModels[T comparable](fac func() T, patterns ...string) []T {
	keys := db.FindKeysLike(patterns...)
	return gs.Map(keys, func(key string, index int) T {
		t := fac()
		db.Get(key, t)
		return t
	})
}
