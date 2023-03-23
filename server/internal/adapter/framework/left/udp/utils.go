package udp

import (
	"strings"
)

func DecodeEvent(rawMsg []byte, n int, from string) *Event {
	event := &Event{
		From: from,
		Name: "",
		Payload: "",
	}

	msg := strings.TrimSpace(string(rawMsg[:n]))
	eventName := strings.Split(msg, " ")[0]

	if eventName == REGISTER {
		username := strings.TrimPrefix(msg, REGISTER+" ")
		event.SetSenderName(username)
		event.SetName(REGISTER)
		event.SetPayload(username) 
	} else if eventName == ALL {
		event.SetName(ALL) 
		event.SetPayload(strings.TrimPrefix(msg, ALL+" "))
	} else {
		event.SetName(PRIVATE_MSG) 
		event.SetPayload(msg)
	}
	
	return event
}
