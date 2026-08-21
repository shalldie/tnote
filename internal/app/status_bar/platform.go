package status_bar

import (
	"github.com/shalldie/tnote/v2/internal/app/pkgs/dialog"
	"github.com/shalldie/tnote/v2/internal/app/store"
	"github.com/shalldie/tnote/v2/internal/conf"
	"github.com/shalldie/tnote/v2/internal/i18n"
)

// zone
var PLATFORM_ID = "STATUSBAR_SHOW_PLATFORM"

// 展示 「平台」弹框
func (m *StatusBarModel) showPlatform() {

	store.Send(dialog.Select(
		i18n.Get(i18nTpl, "platform"),
		i18n.Get(i18nTpl, "selectpf")+"\n",
		[]string{conf.PF_GITHUB, conf.PF_GITEE},
		func(choice string) bool {
			if choice == conf.PF_GITHUB && !conf.HasGithub() {
				go store.Send(store.StatusPayload{
					Message:  "No $TNOTE_GIST_TOKEN in $PATH",
					Duration: 3,
				})
				return false
			}

			if choice == conf.PF_GITEE && !conf.HasGitee() {
				go store.Send(store.StatusPayload{
					Message:  "No $TNOTE_GIST_TOKEN_GITEE in $PATH",
					Duration: 3,
				})
				return false
			}

			conf.PF_CURRENT = choice
			store.SafeGo(store.Setup)
			return true
		},
	).WithWidth(50))
}
