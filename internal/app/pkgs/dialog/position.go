package dialog

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
)

// PlaceOverlay 把 fg 叠加渲染到 bg 之上（模态弹框用）。
//
// 思路源自 https://github.com/charmbracelet/lipgloss/pull/102/files
// 用字符串级拼接而非 lipgloss v2 的原生 Layer/Compositor：后者在合成时会
// 剥离 bubblezone 注入的鼠标区域标记，导致对话框内的按钮/输入框点击失效。
//
// 未来计划：待 bubblezone 支持把区域标记写入 lipgloss v2 的 cell-buffer
// （而非当前的零宽 ANSI 转义流）后，本函数即可替换为原生 lipgloss.Compositor，
// 届时可删除本文件。跟踪点：bubblezone 对 cell-buffer / Canvas 的适配。
func PlaceOverlay(x, y int, fg, bg string) string {
	fgLines, fgWidth := getLines(fg)
	bgLines, bgWidth := getLines(bg)
	bgHeight := len(bgLines)
	fgHeight := len(fgLines)

	if fgWidth >= bgWidth && fgHeight >= bgHeight {
		// FIXME: return fg or bg?
		return fg
	}
	// TODO: allow placement outside of the bg box?
	x = clamp(x, 0, bgWidth-fgWidth)
	y = clamp(y, 0, bgHeight-fgHeight)

	var b strings.Builder
	for i, bgLine := range bgLines {
		if i > 0 {
			b.WriteByte('\n')
		}
		if i < y || i >= y+fgHeight {
			b.WriteString(bgLine)
			continue
		}

		pos := 0
		if x > 0 {
			left := ansi.Truncate(bgLine, x, "")
			pos = ansi.StringWidth(left)
			b.WriteString(left)
			if pos < x {
				// 用空格填充左侧留白（原 lipgloss whitespace 渲染器仅在此处
				// 输出纯空格；带样式/宽字符的填充能力从未启用，故直接内联）
				b.WriteString(strings.Repeat(" ", x-pos))
				pos = x
			}
		}

		fgLine := fgLines[i-y]
		b.WriteString(fgLine)
		pos += ansi.StringWidth(fgLine)

		right := ansi.TruncateLeft(bgLine, pos, "")
		bgWidth := ansi.StringWidth(bgLine)
		rightWidth := ansi.StringWidth(right)
		if rightWidth <= bgWidth-pos {
			b.WriteString(strings.Repeat(" ", bgWidth-rightWidth-pos))
		}

		b.WriteString(right)
	}

	return b.String()
}

func clamp(v, lower, upper int) int {
	return max(lower, min(v, upper))
}

// Split a string into lines, additionally returning the size of the widest
// line.
func getLines(s string) (lines []string, widest int) {
	lines = strings.Split(s, "\n")

	for _, l := range lines {
		w := ansi.StringWidth(l)
		if widest < w {
			widest = w
		}
	}

	return lines, widest
}
