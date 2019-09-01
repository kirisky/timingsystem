package websocket

import (
	"log"
	"time"
	"net/http"

	"github.com/gorilla/websocket"
	// "timingsystem/timingserver/hubs"
)

const (
	// Time allowed to write a message to the peer
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maxium message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

// WSClient is a middleman between the websocket connection and the hub
type WSClient struct {
	wsHub *Hub

	// The websocket connection
	conn *websocket.Conn

	// Buffered channel of outbound messages
	send chan []byte
}

// write message to websocket client
func (w *WSClient) writeMessage() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		w.conn.Close()
	}()

	for {
		select {
		case message, ok := <-w.send:
			w.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				w.conn.WriteMessage(websocket.CloseMessage, []byte{})
			}

			nWriter, err := w.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return 
			}
			nWriter.Write(message)

			n := len(w.send)

			for i := 0; i < n; i++ {
				nWriter.Write(newline)
				nWriter.Write(<-w.send)
			}

			if err := nWriter.Close(); err != nil {
				return
			}

		case <- ticker.C:
			w.conn.SetWriteDeadline(time.Now().Add(writeWait))
			log.Println("ticker c")
			if err := w.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}

}

// read message from websocket client
func (w *WSClient) readMessage() {
	defer func() {
		log.Println("Unregister client")
		w.wsHub.unregister <- w
		w.conn.Close()
	}()

	w.conn.SetReadLimit(maxMessageSize)
	w.conn.SetReadDeadline(time.Now().Add(pongWait))
	w.conn.SetPongHandler(func (string) error {
		w.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := w.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		if string(message) == "reconnect" && len(w.wsHub.latestRecords) > 0 {
			log.Println("Client reconnected! Resend data to the client.")
			w.send <- w.wsHub.latestRecords
		}

		if string(message) == "resetrecords" {
			log.Println("Reset the records.")
			w.wsHub.resetLatestRecords <- true
		}


	}

	log.Println("finish the message reading")
}

// func serveWs(wsHub *Hub, w http.ResponseWriter, r *http.Request, serverHub *hubs.ServerHub ) {
func serveWs(wsHub *Hub, w http.ResponseWriter, r *http.Request) {
	log.Println("Client connected!")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	wsClient := &WSClient{ wsHub: wsHub, conn: conn, send: make(chan []byte, 256), }
	wsClient.wsHub.register <- wsClient

	go wsClient.writeMessage()
	go wsClient.readMessage()

}
