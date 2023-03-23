package udp

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/google/uuid"
)

func (udpa *Adapter) readMsg() {
	for {
		msg := make([]byte, 4096)

		n, addr, err := udpa.conn.ReadFromUDP(msg)
		if err != nil {
			log.Println("fail to read messages", err.Error())
		}
		event := DecodeEvent(msg, n, addr.String())
		udpa.EventRouting(event)
		log.Printf("event: %+v", event)
	}
}

func (udpa *Adapter) UserConnected(addr string, username string) {
	userID := uuid.NewString()
	cl := NewClient(userID, username, addr)
	go udpa.Hub.AddClient(cl)
	log.Printf("new user connected: %v", cl.Addr)
}

func (udpa *Adapter) EventRouting(e *Event) {
	if e.Name == REGISTER {
		udpa.UserConnected(e.From, strings.TrimSpace(e.Payload))
		e.Sender = udpa.Hub.GetClientByAddr(e.From)
		log.Print(e.Sender.Username)
	}

	if e.Name == PRIVATE_MSG {
		receiverName := strings.Split(e.Payload, " ")[0][1:]
		r, err := udpa.Hub.GetClientByName(receiverName)
		if err != nil {
			log.Println(err)
		}
		udpa.SendToClient(e.From, r.Addr, e.Payload)
	}

	if e.Name == ALL {
		udpa.SendToAllClients(e.Payload)
	}
}

func (udpa *Adapter) SendToClient(senderAddr, receiverAddr string, msg string) {
	receiver, ok := udpa.Hub.Clients[receiverAddr]
	if !ok {
		log.Println("cannot find receiver")
	}

	sender, ok := udpa.Hub.Clients[senderAddr]
	if !ok {
		log.Println("cannot find sender")
	}

	receiverUdpAddr, err := net.ResolveUDPAddr("udp", receiver.Addr)
	if err != nil {
		log.Println("fail to convert udp address", err.Error())
		return
	}

	byteMsg := fmt.Sprintf("%s to %s: %s", sender.Username, receiver.Username, msg)

	_, err = udpa.conn.WriteToUDP([]byte(byteMsg), receiverUdpAddr)
	if err != nil {
		log.Print("fail to write message:", err.Error())
	}
}

func (udpa *Adapter) SendToAllClients(msg string) {
	for _, cl := range udpa.Hub.Clients {
		udpAddr, err := net.ResolveUDPAddr("udp", cl.Addr)
		if err != nil {
			log.Println("fail to convert udp address", err.Error())
			return
		}
		udpa.conn.WriteToUDP([]byte(msg), udpAddr)
	}
}
