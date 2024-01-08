package dialog

var (
	ModeConfirm = 0 // confirm 模式
	ModePrompt  = 1 // prompt 模式
)

type DialogPayload struct {
	Mode        int // 0 - confirm 模式 1 - prompt 模式
	Message     string
	PromptValue string
	FnOK        func(args ...string) bool
}

func NewDialogPayload() *DialogPayload {
	return &DialogPayload{
		Mode: ModeConfirm,
	}
}
