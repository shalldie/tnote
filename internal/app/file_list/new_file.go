package file_list

import (
	"fmt"
	"unicode/utf8"

	"github.com/shalldie/tnote/internal/app/pkgs/dialog"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/gist"
)

func validateFilename(filename string) bool {
	if utf8.RuneCountInString(filename) < 3 {

		go store.Send(store.StatusPayload{
			Message:  "文件名长度需要大于3",
			Duration: 3,
		})
		return false
	}
	return true
}

func (m *FileListModel) newFile() {
	store.Send(dialog.DialogPayload{
		Mode:    dialog.ModePrompt,
		Message: "新建文件，请输入文件名",
		FnOK: func(args ...string) bool {
			filename := args[0]

			valid := validateFilename(filename)
			if !valid {
				return false
			}

			go func() {
				go store.Send(store.StatusPayload{
					Loading: true,
					Message: "新建中...",
				})
				store.Gist.UpdateFile(filename, &gist.UpdateGistPayload{Content: "To be edited."})
				store.Send(store.CMD_REFRESH_FILES(filename))
				store.Send(store.CMD_UPDATE_FILE(""))
				go store.Send(store.StatusPayload{
					Loading:  false,
					Message:  fmt.Sprintf("「%v」完成新建", filename),
					Duration: 5,
				})
			}()

			return true
		},
	})
}
