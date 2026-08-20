package dialog

// DialogMode 对话框形态。用独立类型而非裸 int，避免误传任意整数。
type DialogMode int

const (
	ModeAlert   DialogMode = iota // alert 模式：仅 OK 按钮
	ModeConfirm                   // confirm 模式：取消 + OK 按钮
	ModePrompt                    // prompt 模式：多一个文本输入框
)

type DialogPayload struct {
	Mode        DialogMode
	Title       string
	Message     string
	PromptValue string
	SelectList  []string
	Width       int // 宽度，默认 46

	// fnOK 为确认回调，不可导出：外部只能通过下方语义化构造器设置，
	// 由构造器把类型明确的回调适配成内部统一的变长签名，返回 true 关闭对话框。
	fnOK func(args ...string) bool
}

// WithWidth 返回设置了宽度的副本，便于在构造器后链式指定宽度。
func (p DialogPayload) WithWidth(w int) DialogPayload {
	p.Width = w
	return p
}

// Alert 仅有 OK 按钮的提示框。
func Alert(title, message string) DialogPayload {
	return DialogPayload{Mode: ModeAlert, Title: title, Message: message}
}

// Confirm 带「取消 / 确定」两个按钮的确认框。onOK 返回 true 关闭对话框。
func Confirm(title, message string, onOK func() bool) DialogPayload {
	return DialogPayload{
		Mode:    ModeConfirm,
		Title:   title,
		Message: message,
		fnOK:    func(args ...string) bool { return onOK() },
	}
}

// Prompt 带文本输入框的对话框。onOK 收到输入值，返回 true 关闭对话框（校验失败返回 false 可保持打开）。
func Prompt(title, message, initValue string, onOK func(input string) bool) DialogPayload {
	return DialogPayload{
		Mode:        ModePrompt,
		Title:       title,
		Message:     message,
		PromptValue: initValue,
		fnOK:        func(args ...string) bool { return onOK(args[0]) },
	}
}

// Select 在提示框中额外渲染可选列表。onOK 收到选中项，返回 true 关闭对话框。
func Select(title, message string, options []string, onOK func(choice string) bool) DialogPayload {
	return DialogPayload{
		Mode:       ModeAlert,
		Title:      title,
		Message:    message,
		SelectList: options,
		fnOK:       func(args ...string) bool { return onOK(args[0]) },
	}
}
