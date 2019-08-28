package hubs

// ServerHub as a message pipe that sends information from 
// grpc server to  websocket server
type ServerHub struct {
	MessagePipe chan []byte
}

// NewServerHub creates an instance of ServerHub
func NewServerHub() *ServerHub {
	return &ServerHub {
		MessagePipe: make(chan []byte),
	}
}