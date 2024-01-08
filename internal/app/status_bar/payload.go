package status_bar

var S_ID = 1

type StatusPayload struct {
	Loading  bool
	Message  string
	Duration int
}
