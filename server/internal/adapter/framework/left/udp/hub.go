package udp

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

type Hub struct {
	Clients map[string]*Client
	mutex   sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Clients: make(map[string]*Client),
	}
}

func (h *Hub) AddClient(cl *Client) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if _, ok := h.Clients[cl.Addr]; ok {
		log.Print("client already exist")
		return
	}

	h.Clients[cl.Addr] = cl
}

func (h *Hub) RemoveClient(cl *Client) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	if _, ok := h.Clients[cl.Addr]; !ok {
		log.Println("user do not exist")
		return
	}
	delete(h.Clients, cl.Addr)
}

func (h *Hub) GetClientByName(username string) (*Client, error) {
	for _, cl := range h.Clients {
		if strings.EqualFold(cl.Username, username) {
			return cl, nil
		}
	}

	return NewClient("default_id", "default_username", ""), fmt.Errorf("%v do not existed", username)
}

func (h *Hub) GetClientByAddr(addr string) *Client {
	if _, ok := h.Clients[addr]; ok {
		return h.Clients[addr]
	}
	return NewClient("default_id", "default_username", "")
}
