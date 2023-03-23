package udp

type Event struct {
	SenderName string
	Sender     *Client
	From       string
	Name       string
	Payload    string
}

func NewEvent(name, payload string, from string) *Event {
	return &Event{
		From:       from,
		Name:       name,
		Payload:    payload,
	}
}

func (e *Event) SetName(name string) {
	e.Name = name
}
func (e *Event) SetSenderName(name string) {
	e.SenderName = name
}

func (e *Event) SetPayload(payload string) {
	e.Payload = payload
}

func (e *Event) SetAddr(addr string) {
	e.From = addr
}
