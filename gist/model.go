package gist

type FileModel struct {
	FileName string `json:"filename"`
	RawUrl   string `json:"raw_url"`
	Content  string `json:"content"`
}

type GistModel struct {
	Id          string                `json:"id"`
	Url         string                `json:"url"`
	Public      bool                  `json:"public"`
	Description string                `json:"description"`
	Files       map[string]*FileModel `json:"files"`
}

type FetchOptions struct {
	Method  string
	Headers map[string]string
	Query   map[string]string
	Params  map[string]any
}
