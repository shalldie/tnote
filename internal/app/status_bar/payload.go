package status_bar

var ID = 1

type StatusPayload struct {
	Loading  bool
	Message  string
	Duration int
}
