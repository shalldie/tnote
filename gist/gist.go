package gist

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/shalldie/gog/gs"
)

type H map[string]any

type Gist struct {
	TOKEN        string
	Model        *GistModel
	CurrentIndex int
	Files        []*GistFile
}

func NewGist(token string) *Gist {
	return &Gist{
		TOKEN: token,
	}
}

func (g *Gist) Setup() {
	// 1. 找到 gist
	list := g.FetchGists(1, 100)

	item, err := gs.Find(list, func(item *GistModel, i int) bool {
		return item.Description == SPECIAL_DESCRIPTION
	})
	// 2. 第一页没找到，再去第二页找
	if err != nil {
		list = g.FetchGists(2, 100)

		item, err = gs.Find(list, func(item *GistModel, i int) bool {
			return item.Description == SPECIAL_DESCRIPTION
		})
	}

	// 3. 如果没找到，去创建 gist
	if err != nil {
		g.Model = g.CreateGist("newfile.md", "welcome to use tnote >_<#@!")
	} else {
		g.Model = item
	}

	// 4. 有 gist id 后，update 获取所有内容
	g.Update()
}

func (g *Gist) getHeaders() map[string]string {
	return map[string]string{
		"Accept":               "application/vnd.github+json",
		"Authorization":        fmt.Sprintf("bearer %v", g.TOKEN),
		"X-GitHub-Api-Version": "2022-11-28",
	}
}

// func (g *Gist) Query(content string) string {

// 	client := &http.Client{}

// 	paramsObj := make(map[string]string)
// 	paramsObj["query"] = content
// 	paramsBytes, _ := json.Marshal(paramsObj)

// 	req, _ := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewReader(paramsBytes))
// 	req.Header.Add("Authorization", fmt.Sprintf("bearer %v", g.TOKEN))

// 	res, _ := client.Do(req)
// 	body, _ := io.ReadAll(res.Body)

// 	return string(body)
// }

// 获取 gists 列表
func (g *Gist) FetchGists(page int, perPage int) []*GistModel {

	body := fetch("https://api.github.com/gists", &FetchOptions{
		Method: "GET",
		Query: map[string]string{
			"page":     strconv.Itoa(page),
			"per_page": strconv.Itoa(perPage),
		},
		Headers: g.getHeaders(),
	})

	var gistList []*GistModel
	err := json.Unmarshal(body, &gistList)

	if err != nil {
		panic(err)
	}

	return gistList
}

// 获取文件内容
func (g *Gist) FetchFile(fileName string) string {
	targetFile := g.Model.Files[fileName]
	if targetFile == nil {
		return fmt.Sprintf("Error! File %s not found.", fileName)
	}
	fileUrl := targetFile.RawUrl

	body := fetch(fileUrl, &FetchOptions{
		Method:  "GET",
		Headers: g.getHeaders(),
	})

	return string(body)
}

// 创建 gist
func (g *Gist) CreateGist(fileName string, content string) *GistModel {
	body := fetch("https://api.github.com/gists", &FetchOptions{
		Method: "POST",
		Params: H{
			"title":       SPECIAL_DESCRIPTION,
			"description": SPECIAL_DESCRIPTION,
			"files": H{
				fileName: H{
					"content": content,
				},
			},
			"public": false,
		},
		Headers: g.getHeaders(),
	})

	var model *GistModel

	err := json.Unmarshal(body, model)

	if err != nil {
		panic(err)
	}

	return model
}

// 更新文件，https://docs.github.com/zh/rest/gists/gists?apiVersion=2022-11-28#update-a-gist
func (g *Gist) UpdateFile(fileName string, payload *UpdateGistPayload) {

	body := fetch("https://api.github.com/gists/"+g.Model.Id, &FetchOptions{
		Method:  "PATCH",
		Headers: g.getHeaders(),
		Params: H{
			"files": H{
				fileName: payload,
			},
		},
	})

	model := &GistModel{}
	err := json.Unmarshal(body, model)

	if err != nil {
		panic(err)
	}

	// 如果有文件，表示返回内容正常
	// 进行更新
	if len(model.Files) > 0 {
		g.Model = model
		g.updateFiles()
		return
	}

	// 无文件，返回内容异常
	// 全量更新
	g.Update()
}

func (g *Gist) Update() {
	body := fetch("https://api.github.com/gists/"+g.Model.Id, &FetchOptions{
		Method:  "GET",
		Headers: g.getHeaders(),
	})

	model := &GistModel{}
	err := json.Unmarshal(body, model)

	if err != nil {
		panic(err)
	}

	g.Model = model
	g.updateFiles()
}

func (g *Gist) updateFiles() {
	files := make([]*GistFile, 0)
	for fileName := range g.Model.Files {
		files = append(files, g.Model.Files[fileName])
	}
	fileNames := gs.Map(files, func(f *GistFile, _ int) string {
		return f.FileName
	})
	sort.Strings(fileNames)
	files = gs.Sort(files, func(f1 *GistFile, f2 *GistFile) bool {
		return gs.IndexOf(fileNames, f1.FileName) < gs.IndexOf(fileNames, f2.FileName)
	})
	g.Files = files
}

func (g *Gist) GetFile() *GistFile {
	return g.Files[g.CurrentIndex]
}
