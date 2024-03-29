// 全局状态
package store

import (
	"sync"

	"github.com/shalldie/tnote/internal/gist"
)

var Gist *gist.Gist

type storeState struct {
	// 状态
	Status StatusPayload

	// 输入框是否焦点
	InputFocus bool

	// 编辑中
	Editing bool

	// 对话框模式
	DialogMode bool

	// 当前文件
	file *gist.GistFile
}

var fileLock sync.Mutex

func (s *storeState) GetFile() *gist.GistFile {
	fileLock.Lock()
	defer fileLock.Unlock()
	return s.file
}

func (s *storeState) SetFile(file *gist.GistFile) {
	fileLock.Lock()
	defer fileLock.Unlock()
	s.file = file
}

func Setup() {
	if Gist == nil {
		Gist = gist.NewGist()
	}
	Send(StatusPayload{
		Loading: true,
		Message: "loading...",
	})
	Gist.Setup()
	Send(StatusPayload{Loading: false})
	Send(CMD_REFRESH_FILES(""))
	Send(CMD_UPDATE_FILE(""))
}

var State = &storeState{
	Status: StatusPayload{},
}
