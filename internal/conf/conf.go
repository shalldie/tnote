package conf

import (
	"os"
)

var (
	VERSION = "v1.0.1"
)

// https://github.com/charmbracelet/lipgloss/issues/40

// LC_CTYPE="en_US.UTF-8" go run ./full-lipgloss

func init() {
	os.Setenv("RUNEWIDTH_EASTASIAN", "0")
}
