package udp_test

import (
	"lab3/internal/adapter/framework/left/udp"
	"testing"
)

func TestDecodeEvent(t *testing.T) {
	// testing register event
	rawEvent := []byte("@register testing event decoder")
	randomUDPAddress := RandomUDPAddress().String()
	e := udp.DecodeEvent(rawEvent, 31, randomUDPAddress)

	if e.From != randomUDPAddress {
		t.Errorf("fail to convert UDP Address")
	}

	if e.Name != "@register" {
		t.Errorf("fail to convert event name")
	}

	if e.Payload != "testing event decoder" {
		t.Errorf("fail to convert payload")
	}

	// testing private message event
	rawPrivateMsgEvent := []byte("@ll testing event decoder")
	randomUDPAddress2 := RandomUDPAddress().String()
	pme := udp.DecodeEvent(rawPrivateMsgEvent, 25, randomUDPAddress2)

	if pme.From != randomUDPAddress2 {
		t.Errorf("fail to convert UDP Address")
	}

	if pme.Name != "@private_msg" {
		t.Errorf("fail to convert event name")
	}

	if pme.Payload != "@ll testing event decoder" {
		t.Errorf("fail to convert payload")
	}
}