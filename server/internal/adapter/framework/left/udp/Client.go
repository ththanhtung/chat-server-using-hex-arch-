package udp

type Client struct {
	ID       string
	Addr     string
	Username string
}

func NewClient(id, username string, addr string) *Client {
	return &Client{
		ID:       id,
		Addr:     addr,
		Username: username,
	}
}