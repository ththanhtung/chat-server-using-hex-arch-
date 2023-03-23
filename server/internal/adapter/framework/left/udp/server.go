package udp

import (
	"lab3/internal/ports"
	"log"
	"net"
	"strconv"
)

type Adapter struct {
	api  ports.APIPort
	Hub  *Hub
	conn *net.UDPConn
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api:  api,
		Hub:  NewHub(),
	}
}

func (udpa *Adapter) Listen(port int) {
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+strconv.FormatInt(int64(port), 10))

	if err != nil {
		log.Fatalf("fail to resolve udp addr %v", err.Error())
	}

	conn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatal("fail to start server", err.Error())
	}

	udpa.conn = conn

	udpa.readMsg()

	log.Println("server is up on port 3000")
}
