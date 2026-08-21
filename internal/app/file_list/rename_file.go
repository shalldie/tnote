package file_list

import (
	"fmt"

	"github.com/shalldie/tnote/v2/internal/app/pkgs/dialog"
	"github.com/shalldie/tnote/v2/internal/app/store"
	"github.com/shalldie/tnote/v2/internal/gist"
	"github.com/shalldie/tnote/v2/internal/i18n"
)

func (m *FileListModel) renameFile(file *gist.GistFile) {
	filename := file.FileName
	store.Send(dialog.Prompt(
		i18n.Get(i18nTpl, "rename_title"),
		fmt.Sprintf(i18n.Get(i18nTpl, "rename_message"), filename),
		filename,
		func(newname string) bool {
			valid := validateFilename(newname)
			if !valid {
				return false
			}

			store.SafeGo(func() {
				go store.Send(store.StatusPayload{
					Loading: true,
					Message: i18n.Get(i18nTpl, "rename_renaming"),
				})

				store.Gist.UpdateFile(filename, &gist.UpdateGistPayload{
					Filename: newname,
					Content:  file.Content, // gitee 需要这个
				})

				store.Send(store.StatusPayload{
					Loading:  false,
					Message:  fmt.Sprintf(i18n.Get(i18nTpl, "rename_done"), filename, newname),
					Duration: 5,
				})
				store.Send(store.CMD_REFRESH_FILES(newname))
			})

			return true
		},
	))

}
