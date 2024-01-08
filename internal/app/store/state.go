// 全局状态
package store

type storeState struct {
	// 输入框是否焦点
	InputFocus bool
}

var State = &storeState{}
