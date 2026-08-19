package store

import (
	"fmt"

	"github.com/shalldie/tnote/internal/i18n"
)

var i18nTpl = `
network_error:
  en_US: "Network request failed, please check your network or proxy: %v"
  zh_CN: "网络请求失败，请检查网络或代理: %v"
`

// SafeGo starts a background goroutine with a panic guard.
// Any async task that touches the network should be launched with it,
// so a failed request turns into a status-bar message instead of crashing the app.
func SafeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				Send(StatusPayload{Loading: false})
				Send(StatusPayload{
					Message:  fmt.Sprintf(i18n.Get(i18nTpl, "network_error"), r),
					Duration: 5,
				})
			}
		}()
		fn()
	}()
}
