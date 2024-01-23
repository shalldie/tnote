package file_list

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/shalldie/tnote/internal/i18n"
)

func newListKeyMap() list.KeyMap {
	return list.KeyMap{
		// Browsing.
		CursorUp: key.NewBinding(
			key.WithKeys("up", "k"),
			// key.WithHelp("↑/k", "上"),
			key.WithHelp("↑/k", i18n.Get(i18nTpl, "key_up")),
		),
		CursorDown: key.NewBinding(
			key.WithKeys("down", "j"),
			// key.WithHelp("↓/j", "下"),
			key.WithHelp("↓/j", i18n.Get(i18nTpl, "key_down")),
		),
		PrevPage: key.NewBinding(
			// key.WithKeys("left", "h", "pgup", "b", "u"),
			key.WithKeys("h", "pgup"),
			// key.WithHelp("←/h/pgup", "prev page"),
			// key.WithHelp("h/pgup", "上一页"),
			key.WithHelp("h/pgup", i18n.Get(i18nTpl, "key_pgup")),
		),
		NextPage: key.NewBinding(
			// key.WithKeys("right", "l", "pgdown", "f", "d"),
			key.WithKeys("l", "pgdown"),
			// key.WithHelp("→/l/pgdn", "next page"),
			// key.WithHelp("l/pgdn", "下一页"),
			key.WithHelp("l/pgdn", i18n.Get(i18nTpl, "key_pgdown")),
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
			// key.WithHelp("/", "过滤"),
			key.WithHelp("/", i18n.Get(i18nTpl, "key_filter")),
		),
		// ClearFilter: key.NewBinding(
		// 	key.WithKeys("esc"),
		// 	// key.WithHelp("esc", "clear filter"),
		// 	key.WithHelp("esc", "清空过滤条件"),
		// ),

		// Filtering.
		CancelWhileFiltering: key.NewBinding(
			key.WithKeys("esc"),
			// key.WithHelp("esc", "cancel"),
			// key.WithHelp("esc", "取消"),
			key.WithHelp("esc", i18n.Get(i18nTpl, "key_cancel")),
		),
		AcceptWhileFiltering: key.NewBinding(
			key.WithKeys("enter", "tab", "shift+tab", "ctrl+k", "up", "ctrl+j", "down"),
			// key.WithHelp("enter", "apply filter"),
			// key.WithHelp("enter", "应用过滤"),
			key.WithHelp("enter", i18n.Get(i18nTpl, "key_filter_apply")),
		),

		// Toggle help.
		ShowFullHelp: key.NewBinding(
			key.WithKeys("?"),
			// key.WithHelp("?", "更多"),
			key.WithHelp("?", i18n.Get(i18nTpl, "key_filter_more")),
		),
		CloseFullHelp: key.NewBinding(
			key.WithKeys("?"),
			// key.WithHelp("?", "折叠"),
			key.WithHelp("?", i18n.Get(i18nTpl, "key_filter_less")),
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
			// key.WithHelp("n", "新建"),
			key.WithHelp("n", i18n.Get(i18nTpl, "key_new")),
		),
		key.NewBinding(
			key.WithKeys("r"),
			// key.WithHelp("r", "重命名"),
			key.WithHelp("r", i18n.Get(i18nTpl, "key_rename")),
		),
		key.NewBinding(
			key.WithKeys("e"),
			// key.WithHelp("e", "编辑"),
			key.WithHelp("e", i18n.Get(i18nTpl, "key_edit")),
		),
		key.NewBinding(
			key.WithKeys("d"),
			// key.WithHelp("d", "删除"),
			key.WithHelp("d", i18n.Get(i18nTpl, "key_del")),
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
