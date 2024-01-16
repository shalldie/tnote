package file_list

import (
	"fmt"

	"github.com/shalldie/tnote/internal/app/pkgs/dialog"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/gist"
)

func (m *FileListModel) renameFile(filename string) {
	store.Send(dialog.DialogPayload{
		Mode:        dialog.ModePrompt,
		Message:     fmt.Sprintf("重命名文件「%v」", filename),
		PromptValue: filename,
		FnOK: func(args ...string) bool {
			newname := args[0]

			valid := validateFilename(newname)
			if !valid {
				return false
			}

			go func() {
				go store.Send(store.StatusPayload{
					Loading: true,
					Message: "重命名中...",
				})

				store.Gist.UpdateFile(filename, &gist.UpdateGistPayload{Filename: newname})

				go store.Send(store.StatusPayload{
					Loading:  false,
					Message:  fmt.Sprintf("「%v」->「%v」完成重命名", filename, newname),
					Duration: 5,
				})
				go store.Send(store.CMD_REFRESH_FILES(newname))
			}()

			return true
		},
	})

}
