package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/shalldie/gog/gs"
	"github.com/shalldie/ttm/app"
	"github.com/shalldie/ttm/db"
	"github.com/shalldie/ttm/model"
)

func main() {
	main1()
	// main3()
}

func main1() {
	app.Setup()
}

func main2() {
	keys := db.FindKeysLike("project_")
	gs.ForEach(keys, func(key string, index int) {
		p := model.NewProject()
		db.Get(key, p)
		fmt.Println(key, p.ID)
	})
}

func main3() {
	// m := db.FindByPattern("project_")

	data := db.Get("project_b1be4454-0b33-4533-821b-e5f5b1fc0b0f", nil)
	prj := model.NewProject()

	decode := gob.NewDecoder(bytes.NewBuffer(data))
	decode.Decode(prj)

	fmt.Println(prj.ID)

	// for key, data := range m {
	// 	prj := model.NewProject()
	// 	decode := gob.NewDecoder(bytes.NewBuffer(data))
	// 	decode.Decode(prj)

	// 	fmt.Println(key, prj.ID)

	// }
}
