// 全局状态
package store

import "github.com/shalldie/tnote/internal/gist"

var Gist *gist.Gist

type storeState struct {
	// gist 模块
	// Gist *gist.Gist

	// 状态
	Status StatusPayload

	// 输入框是否焦点
	InputFocus bool
}

func Setup(token string) {
	Gist = gist.NewGist(token)
	Gist.Setup()
}

var State = &storeState{
	Status: StatusPayload{},
}
