package conf

import (
	"fmt"
	"os"
)

var (
	// 版本
	VERSION = "v1.3.0"
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

	// 啥都没配置
	if !HasGithub() && !HasGitee() {
		fmt.Println("Can't find any token in $PATH")
		os.Exit(1)
	}

}
