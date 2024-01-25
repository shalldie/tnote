package file_list

import (
	"fmt"
	"unicode/utf8"

	"github.com/shalldie/tnote/internal/app/pkgs/dialog"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/gist"
	"github.com/shalldie/tnote/internal/i18n"
)

func validateFilename(filename string) bool {
	if utf8.RuneCountInString(filename) < 3 {

		go store.Send(store.StatusPayload{
			Message:  i18n.Get(i18nTpl, "new_namevalid"),
			Duration: 3,
		})
		return false
	}
	return true
}

func (m *FileListModel) newFile() {
	store.Send(dialog.DialogPayload{
		Mode:    dialog.ModePrompt,
		Message: i18n.Get(i18nTpl, "new_message"),
		FnOK: func(args ...string) bool {
			filename := args[0]

			valid := validateFilename(filename)
			if !valid {
				return false
			}

			go func() {
				go store.Send(store.StatusPayload{
					Loading: true,
					Message: i18n.Get(i18nTpl, "new_creating"),
				})
				store.Gist.UpdateFile(filename, &gist.UpdateGistPayload{Content: "To be edited."})

				store.Send(store.StatusPayload{
					Loading:  false,
					Message:  fmt.Sprintf(i18n.Get(i18nTpl, "new_done"), filename),
					Duration: 5,
				})
				store.Send(store.CMD_REFRESH_FILES(filename))
				store.Send(store.CMD_UPDATE_FILE(""))
			}()

			return true
		},
	})
}
