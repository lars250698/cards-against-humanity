package websocket

type Hub struct {
	clients map[*Client]bool
	broadcast chan Broadcast
	register chan *Client
	unregister chan *Client
}

type Broadcast struct {
	message []byte
	targets []*Client
}

func newHub () *Hub {
	return &Hub {
		broadcast: make(chan Broadcast),
		register: make(chan *Client),
		unregister: make(chan *Client),
		clients: make(map[*Client]bool),
	}
}

func (h *Hub) run () {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <- h.broadcast:
			for _, client := range message.targets {
				select {
				case client.send <- message.message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
