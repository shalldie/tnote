package status_bar

import (
	"fmt"
	"strings"

	"github.com/shalldie/tnote/internal/app/pkgs/dialog"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/conf"
	"github.com/shalldie/tnote/internal/utils"
)

// zone
var ABOUT_ID = "STATUSBAR_SHOW_ABOUT"

// 展示 「关于」信息
func (m *StatusBarModel) showAbout() {

	content := fmt.Sprintf(`
# tnote

Note in terminal. 终端运行的记事本。

> Version `+"`%v`"+`
> [Github](https://github.com/shalldie/tnote)
				`, conf.VERSION)

	message := utils.RenderMarkdown(strings.TrimSpace(content), 50)

	store.Send(dialog.DialogPayload{
		Mode:    dialog.ModeAlert,
		Message: message,
		Width:   50,
	})
}
