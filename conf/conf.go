package conf

import (
	"os"
)

// https://github.com/charmbracelet/lipgloss/issues/40

// LC_CTYPE="en_US.UTF-8" go run ./full-lipgloss

func init() {
	os.Setenv("RUNEWIDTH_EASTASIAN", "0")
}
