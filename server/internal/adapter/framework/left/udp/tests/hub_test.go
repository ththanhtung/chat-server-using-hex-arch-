package udp_test

import (
	"lab3/internal/adapter/framework/left/udp"
	"math/rand"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestGetClientByName(t *testing.T) {
	h := udp.NewHub()

	cl := udp.NewClient(uuid.NewString(), "kk", RandomUDPAddress().String())
	cl2 := udp.NewClient(uuid.NewString(), "qq", RandomUDPAddress().String())
	cl3 := udp.NewClient(uuid.NewString(), "oo", RandomUDPAddress().String())
	cl4 := udp.NewClient(uuid.NewString(), "zz", RandomUDPAddress().String())

	h.AddClient(cl)
	h.AddClient(cl2)
	h.AddClient(cl3)
	h.AddClient(cl4)

	user, _ := h.GetClientByName("zz")
	if user.Username != "zz" {
		t.Errorf("fail to get client")
	}
}

func RandomUDPAddress() *net.UDPAddr {
	// Generate a random port number between 1024 and 65535
	port := CreateRandomNumb(1024, 65535)

	// Generate a random IP address
	ip := net.IPv4(byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)))

	// Create the UDP address
	addr := net.UDPAddr{
		IP:   ip,
		Port: port,
	}
	return &addr
}

func CreateRandomNumb(min, max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
    rand := rand.New(seed)

	return rand.Intn(max - min + 1) + min
}