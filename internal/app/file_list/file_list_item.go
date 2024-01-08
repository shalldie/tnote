package file_list

import "github.com/shalldie/tnote/internal/gist"

type FileListItem struct {
	gistfile *gist.GistFile
}

func (item FileListItem) Title() string       { return item.gistfile.FileName }
func (item FileListItem) Description() string { return item.gistfile.Content }
func (item FileListItem) FilterValue() string { return item.Title() }
