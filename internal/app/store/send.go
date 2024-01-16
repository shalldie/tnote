package store

// 通知 - impl
var SendImpl func(cmd any)

// 通知
func Send(cmd any) {
	switch msg := cmd.(type) {

	case StatusPayload:
		State.Status = msg

	case CMD_INVOKE_EDIT:
		State.Editing = bool(msg)

	}

	if SendImpl != nil {
		SendImpl(cmd)
	}
}
