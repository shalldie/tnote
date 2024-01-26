package file_list

import (
	"fmt"

	"github.com/shalldie/tnote/internal/app/pkgs/dialog"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/gist"
	"github.com/shalldie/tnote/internal/i18n"
)

func (m *FileListModel) renameFile(filename string) {
	store.Send(dialog.DialogPayload{
		Mode:        dialog.ModePrompt,
		Title:       i18n.Get(i18nTpl, "rename_title"),
		Message:     fmt.Sprintf(i18n.Get(i18nTpl, "rename_message"), filename),
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
					Message: i18n.Get(i18nTpl, "rename_renaming"),
				})

				store.Gist.UpdateFile(filename, &gist.UpdateGistPayload{Filename: newname})

				store.Send(store.StatusPayload{
					Loading:  false,
					Message:  fmt.Sprintf(i18n.Get(i18nTpl, "rename_done"), filename, newname),
					Duration: 5,
				})
				store.Send(store.CMD_REFRESH_FILES(newname))
			}()

			return true
		},
	})

}
