// 全局状态
package store

import "github.com/shalldie/tnote/internal/gist"

var Gist *gist.Gist

type storeState struct {
	// 状态
	Status StatusPayload

	// 输入框是否焦点
	InputFocus bool

	// 编辑中
	Editing bool
}

func Setup(token string) {
	Gist = gist.NewGist(token).Setup()
}

var State = &storeState{
	Status: StatusPayload{},
}
