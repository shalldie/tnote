package gist

type GistFile struct {
	FileName  string `json:"filename"`
	Truncated bool   `json:"truncated"` // 是否需要使用 RawUrl 来获取内容
	RawUrl    string `json:"raw_url"`
	Content   string `json:"content"`
}

type GistModel struct {
	Id          string               `json:"id"`
	Url         string               `json:"url"`
	Public      bool                 `json:"public"`
	Description string               `json:"description"`
	Files       map[string]*GistFile `json:"files"`
}

type FetchOptions struct {
	Method  string
	Headers map[string]string
	Query   map[string]string
	Params  map[string]any
}
