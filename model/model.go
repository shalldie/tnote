package model

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/google/uuid"
	"github.com/shalldie/ttm/db"
)

type Model struct {
	ID          string
	CreatedTime int64
}

func NewModel() *Model {
	return &Model{
		ID:          uuid.NewString(),
		CreatedTime: time.Now().Unix(),
	}
}

func findModels[T comparable](fac func() T, patterns ...string) []T {
	list := []T{}
	m := db.FindByPattern(patterns...)

	for _, data := range m {
		sender := fac()
		decode := gob.NewDecoder(bytes.NewBuffer(data))
		decode.Decode(sender)
		list = append(list, sender)
	}
	return list
}
