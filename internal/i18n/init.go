package i18n

import (
	"os"
	"strings"
)

func init() {
	lan := os.Getenv("TNOTE_LANG")
	if len(lan) <= 0 {
		lan = os.Getenv("LANG")
	}

	lan = strings.ReplaceAll(lan, "-", "_") // en-US, zh-CN => en_US, zh_CN
	if strings.Contains(lan, LANG_ZH_CN) {
		LANG = LANG_ZH_CN
	}
}
