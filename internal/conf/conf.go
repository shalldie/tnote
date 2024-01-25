package conf

import (
	"os"
)

var (
	// 版本
	VERSION = "v1.2.0"
	// filelist 宽度
	FileListWidth = 42
)

// https://github.com/charmbracelet/lipgloss/issues/40

// LC_CTYPE="en_US.UTF-8" go run ./full-lipgloss

func init() {
	os.Setenv("RUNEWIDTH_EASTASIAN", "0")
}
