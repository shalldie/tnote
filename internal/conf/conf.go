package conf

import (
	"fmt"
	"os"
)

var (
	// 版本
	VERSION = "v1.2.1"
	// gist token
	TNOTE_GIST_TOKEN = ""
	// gist token gitee
	TNOTE_GIST_TOKEN_GITEE = ""
	// filelist 宽度
	FileListWidth = 44
)

// https://github.com/charmbracelet/lipgloss/issues/40

// LC_CTYPE="en_US.UTF-8" go run ./full-lipgloss

func init() {
	os.Setenv("RUNEWIDTH_EASTASIAN", "0")

	TNOTE_GIST_TOKEN = os.Getenv("TNOTE_GIST_TOKEN")
	TNOTE_GIST_TOKEN_GITEE = os.Getenv("TNOTE_GIST_TOKEN_GITEE")

	if TNOTE_GIST_TOKEN == "" && TNOTE_GIST_TOKEN_GITEE == "" {
		fmt.Println("Can't find any $TNOTE_GIST_TOKEN in $PATH")
		os.Exit(1)
	}

	if TNOTE_GIST_TOKEN_GITEE != "" {
		ENV_CURRENT = ENV_GITEE
	}

}
