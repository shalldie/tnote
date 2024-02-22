package dialog

var (
	ModeAlert   = 0 // alert 模式
	ModeConfirm = 1 // confirm 模式
	ModePrompt  = 2 // prompt 模式
)

type DialogPayload struct {
	Mode        int // 0 - alert模式，1 - confirm模式，2 - prompt模式
	Title       string
	Message     string
	PromptValue string
	SelectList  []string
	FnOK        func(args ...string) bool
	Width       int // 宽度，默认 42
}
