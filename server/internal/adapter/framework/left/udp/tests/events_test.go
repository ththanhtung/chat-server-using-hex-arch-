package udp_test

import (
	"lab3/internal/adapter/framework/left/udp"
	"testing"
)

func TestEvents(t *testing.T){
	e := udp.NewEvent("@register", "testing", RandomUDPAddress().String())

	e.SetName("@all")
	if e.Name != "@all"{
		t.Errorf("fail to change event name")
	}
	
	e.SetPayload("testing2")
	if e.Payload != "testing2" {
		t.Errorf("fail to change payload")
	}

	udpAddr := RandomUDPAddress().String()
	e.SetAddr(udpAddr)
	if e.From != udpAddr {
		t.Errorf("fail to change UDP address")
	}
}