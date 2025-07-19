package protocol

type Request struct {
	Data              []byte
	responseValidator func([]byte) error
}

func (r Request) ValidateResponse(response []byte) error {
	return r.responseValidator(response)
}

func NewDiscoverRequest() Request {
	return NewAa55Protocol("010200", 0x0182)
}

func NewSearchRequest() Request {
	//goland:noinspection SpellCheckingInspection
	return Request{
		Data: []byte("WIFIKIT-214028-READ"),
		responseValidator: func(bytes []byte) error {
			// TODO: check if there is IP ..
			return nil
		},
	}
}
