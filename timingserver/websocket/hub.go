package websocket

import (
	"timingsystem/timingserver/hubs"
)

// Hub maintains the set of active websocket clients and broadcasts messages to the clients
type Hub struct {
	// Registered clients.
	clients map[*WSClient]bool

	// Outbound messages to the clients
	broadcast chan []byte

	// Register requirets from clients
	register chan *WSClient

	// Unregister requests from clients
	unregister chan *WSClient

	// latest records
	latestRecords []byte
}

// register/unregister client
// listen messages from gRPC service by channel and send the message to websocket clients
func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				close(client.send)
				delete(h.clients, client)
			}

		case message := <-h.broadcast:
			h.latestRecords = message
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}


	}


}

// create a instance of Hub
func newHub(serverHub *hubs.ServerHub) *Hub {
	return &Hub {
		broadcast: serverHub.MessagePipe,
		register: make(chan *WSClient),
		unregister: make(chan *WSClient),
		clients: make(map[*WSClient]bool),
	}
}