package file_list

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

func newListKeyMap() list.KeyMap {
	return list.KeyMap{
		// Browsing.
		CursorUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "上"),
		),
		CursorDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "下"),
		),
		PrevPage: key.NewBinding(
			// key.WithKeys("left", "h", "pgup", "b", "u"),
			key.WithKeys("h", "pgup"),
			// key.WithHelp("←/h/pgup", "prev page"),
			key.WithHelp("h/pgup", "上一页"),
		),
		NextPage: key.NewBinding(
			// key.WithKeys("right", "l", "pgdown", "f", "d"),
			key.WithKeys("l", "pgdown"),
			// key.WithHelp("→/l/pgdn", "next page"),
			key.WithHelp("l/pgdn", "下一页"),
		),
		// GoToStart: key.NewBinding(
		// 	key.WithKeys("home", "g"),
		// 	key.WithHelp("g/home", "开头"),
		// ),
		// GoToEnd: key.NewBinding(
		// 	key.WithKeys("end"),
		// 	key.WithHelp("end", "尾部"),
		// ),
		Filter: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "过滤"),
		),
		ClearFilter: key.NewBinding(
			key.WithKeys("esc"),
			// key.WithHelp("esc", "clear filter"),
			key.WithHelp("esc", "清空过滤条件"),
		),

		// Filtering.
		CancelWhileFiltering: key.NewBinding(
			key.WithKeys("esc"),
			// key.WithHelp("esc", "cancel"),
			key.WithHelp("esc", "取消"),
		),
		AcceptWhileFiltering: key.NewBinding(
			key.WithKeys("enter", "tab", "shift+tab", "ctrl+k", "up", "ctrl+j", "down"),
			// key.WithHelp("enter", "apply filter"),
			key.WithHelp("enter", "应用过滤"),
		),

		// Toggle help.
		ShowFullHelp: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "更多"),
		),
		CloseFullHelp: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "折叠"),
		),

		// Quitting.
		// Quit: key.NewBinding(
		// 	key.WithKeys("q", "esc"),
		// 	key.WithHelp("q", "quit"),
		// ),
		// ForceQuit: key.NewBinding(key.WithKeys("ctrl+c")),
	}
}

func additionalKeyMap() []key.Binding {
	return []key.Binding{
		key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "新建"),
		),
		key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "重命名"),
		),
		key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "编辑"),
		),
		key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "删除"),
		),
		// openDirectoryKey,
		// createFileKey,
		// createDirectoryKey,
		// deleteItemKey,
		// copyItemKey,
		// zipItemKey,
		// unzipItemKey,
		// toggleHiddenKey,
		// homeShortcutKey,
		// copyToClipboardKey,
		// escapeKey,
		// renameItemKey,
		// openInEditorKey,
		// submitInputKey,
		// moveItemKey,
	}
}
