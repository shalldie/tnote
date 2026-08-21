package file_list

import (
	"strings"
	"unicode/utf8"

	"charm.land/lipgloss/v2"
	zone "github.com/lrstanley/bubblezone/v2"
	"github.com/shalldie/tnote/v2/internal/conf"
	"github.com/shalldie/tnote/v2/internal/gist"
)

type FileListItem struct {
	ID string
	*gist.GistFile
}

// 文字太长被截断的时候，zone会不生效
// 这里动态计算下最适合的宽度，填充「…」
func withEllipsis(content string) string {
	ellipsis := "…"

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

func (item FileListItem) Title() string {
	return zone.Mark(item.ID+"title", withEllipsis(item.GistFile.FileName))
}
func (item FileListItem) Description() string {
	return zone.Mark(item.ID+"des", withEllipsis(item.Content))
}
func (item FileListItem) FilterValue() string { return item.FileName }
