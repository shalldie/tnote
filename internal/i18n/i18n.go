package i18n

import (
	"gopkg.in/yaml.v3"
)

var (
	// current lang
	LANG = "en_US"

	// en_US
	LANG_EN_US = LANG

	// zh_CN
	LANG_ZH_CN = "zh_CN"
)

/*
从 yaml 模板提取内容

example:

var i18nTpl = `

	en-US:
		hello: hello
	zh-CN:
		hello: 你好

`

	word := i18n.Get(i18nTpl, "hello")
*/
func Get(tpl string, key string) string {
	// if key == "editor_placeholder" {
	// 	println("editor_placeholder")
	// }
	dic := map[string](map[string]string){}
	err := yaml.Unmarshal([]byte(tpl), dic)
	if err != nil {
		panic(err)
	}
	return dic[LANG][key]
}
