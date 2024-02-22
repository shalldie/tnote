package store

type CMD_APP_FOCUS int

// 更新全部文件
type CMD_REFRESH_FILES string

// 更新当前文件
type CMD_UPDATE_FILE string

// 选择文件
type CMD_SELECT_FILE string

// 触发编辑模式
type CMD_INVOKE_EDIT bool

// 展示平台选择
type CMD_SHOW_PLATFORM bool
