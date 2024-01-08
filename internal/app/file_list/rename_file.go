package file_list

import (
	"fmt"

	"github.com/shalldie/tnote/internal/app/pkgs/dialog"
	"github.com/shalldie/tnote/internal/app/status_bar"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/gist"
)

func (m *FileListModel) renameFile(filename string) {
	store.Send(dialog.DialogPayload{
		Mode:        1,
		Message:     fmt.Sprintf("重命名文件「%v」", filename),
		PromptValue: filename,
		FnOK: func(args ...string) bool {
			newname := args[0]

			valid := validateFilename(newname)
			if !valid {
				return false
			}

			go func() {
				go store.Send(status_bar.StatusPayload{
					Loading: true,
					Message: "重命名中...",
				})

				m.gist.UpdateFile(filename, &gist.UpdateGistPayload{Filename: newname})

				go store.Send(status_bar.StatusPayload{
					Loading: false,
					Message: "",
				})
				go store.Send(store.CMD_REFRESH_FILES(newname))
			}()

			return true
		},
	})

}
