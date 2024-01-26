package file_list

import (
	"fmt"

	"github.com/shalldie/tnote/internal/app/pkgs/dialog"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/i18n"
)

func (m *FileListModel) delFile(filename string) {
	store.Send(dialog.DialogPayload{
		Mode:    dialog.ModeConfirm,
		Title:   i18n.Get(i18nTpl, "del_title"),
		Message: fmt.Sprintf(i18n.Get(i18nTpl, "del_confirm"), filename),
		FnOK: func(args ...string) bool {
			go func() {
				go store.Send(store.StatusPayload{
					Loading: true,
					Message: i18n.Get(i18nTpl, "del_deleting"),
				})

				store.Gist.UpdateFile(filename, nil)

				store.Send(store.StatusPayload{
					Loading:  false,
					Message:  fmt.Sprintf(i18n.Get(i18nTpl, "del_done"), filename),
					Duration: 5,
				})
				store.Send(store.CMD_REFRESH_FILES(""))
				store.Send(store.CMD_UPDATE_FILE(""))
			}()

			return true
		},
	})

}
