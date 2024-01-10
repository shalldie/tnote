package file_list

import (
	"fmt"

	"github.com/shalldie/tnote/internal/app/pkgs/dialog"
	"github.com/shalldie/tnote/internal/app/store"
)

func (m *FileListModel) delFile(filename string) {
	store.Send(dialog.DialogPayload{
		Mode:    0,
		Message: fmt.Sprintf("确定要删除文件「%v」吗？", filename),
		FnOK: func(args ...string) bool {
			go func() {
				go store.Send(store.StatusPayload{
					Loading: true,
					Message: "删除中...",
				})
				// note.StatusBar.ShowMessage("删除中...")
				store.Gist.UpdateFile(filename, nil)
				// p.LoadFiles()
				// note.StatusBar.ShowForSeconds("操作成功", 3)
				go store.Send(store.StatusPayload{
					Loading:  false,
					Message:  fmt.Sprintf("「%v」完成删除", filename),
					Duration: 5,
				})
				go store.Send(store.CMD_REFRESH_FILES(""))
			}()

			return true
		},
	})

}