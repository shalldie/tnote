package app

type BaseModel struct {
	Width  int
	Height int
	Active bool
}

// 调整尺寸
func (m *BaseModel) Resize(width int, height int) {
	m.Width = width
	m.Height = height
}

// 获取焦点
func (m *BaseModel) Focus() {
	m.Active = true
}

// 失去焦点
func (m *BaseModel) Blur() {
	m.Active = false
}

func newBaseModel() *BaseModel {
	return &BaseModel{}
}
