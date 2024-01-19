package file_list

import (
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/shalldie/tnote/internal/conf"
	"github.com/shalldie/tnote/internal/gist"
)

type FileListItem struct {
	ID string
	*gist.GistFile
}

func withEllipsis(content string) string {
	ellipsis := "..."

	content = strings.Split(content, "\n")[0]
	runeContent := []rune(content)

	limitWidth := conf.FileListWidth - 2*2
	needEllipsis := lipgloss.Width(content) >= limitWidth

	if needEllipsis {
		for i := len(runeContent); i >= 0; i-- {
			subcontent := string(runeContent[0:i]) + ellipsis
			if lipgloss.Width(subcontent) < limitWidth {
				content = subcontent
				break
			}
		}
	}

	if utf8.RuneCountInString(content) <= 0 {
		return ellipsis
	}
	return content
}

// FIXME: 文字太长被截断的时候，zone不生效
func (item FileListItem) Title() string {
	return zone.Mark(item.ID+"title", withEllipsis(item.GistFile.FileName))
}
func (item FileListItem) Description() string {
	return zone.Mark(item.ID+"des", withEllipsis(item.Content))
}
func (item FileListItem) FilterValue() string { return item.FileName }
