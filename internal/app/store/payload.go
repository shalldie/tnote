package store

// status 状态
type StatusPayload struct {
	Loading  bool
	Message  string
	Duration int
}
