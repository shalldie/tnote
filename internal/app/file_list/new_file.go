package file_list

import (
	"unicode/utf8"

	"github.com/shalldie/tnote/internal/app/pkgs/dialog"
	"github.com/shalldie/tnote/internal/app/status_bar"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/gist"
)

func (m *FileListModel) newFile() {
	store.Send(dialog.DialogPayload{
		Mode:    1,
		Message: "新建文件，请输入文件名",
		FnOK: func(args ...string) bool {
			filename := args[0]

			if utf8.RuneCountInString(filename) < 3 {

				go store.Send(status_bar.StatusPayload{
					Message:  "文件名长度需要大于3",
					Duration: 3,
				})
				return false
			}

			go func() {
				go store.Send(status_bar.StatusPayload{
					Loading: true,
					Message: "新建中...",
				})
				m.gist.UpdateFile(filename, &gist.UpdateGistPayload{Content: "To be edited."})
				store.Send(store.CMD_REFRESH_FILES(""))
				go store.Send(status_bar.StatusPayload{
					Loading: false,
					Message: "",
				})
			}()

			return true
		},
	})
}
