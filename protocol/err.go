package protocol

type PartialResponseErr struct {
	Message string
}

func (e *PartialResponseErr) Error() string {
	return e.Message
}
