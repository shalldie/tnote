package status_bar

import (
	"github.com/shalldie/tnote/internal/app/pkgs/dialog"
	"github.com/shalldie/tnote/internal/app/store"
	"github.com/shalldie/tnote/internal/conf"
)

// zone
var PLATFORM_ID = "STATUSBAR_SHOW_PLATFORM"

// 展示 「平台」弹框
func (m *StatusBarModel) showPlatform() {

	store.Send(dialog.DialogPayload{
		Mode:       dialog.ModeAlert,
		Title:      "平台",
		Message:    "选择要使用的平台\n",
		SelectList: []string{conf.PF_GITHUB, conf.PF_GITEE},
		Width:      50,
		FnOK: func(args ...string) bool {
			if args[0] == conf.PF_GITHUB && !conf.HasGithub() {
				go store.Send(store.StatusPayload{
					Message:  "",
					Duration: 3,
				})
				return false
			}

			if args[0] == conf.PF_GITEE && !conf.HasGitee() {
				go store.Send(store.StatusPayload{
					Message:  "",
					Duration: 3,
				})
				return false
			}

			conf.PF_CURRENT = args[0]
			go store.Setup()
			return true
		},
	})
}
