package main

import (
	"timingsystem/timingserver/websocket"
	"timingsystem/timingserver/grpc"
	"timingsystem/timingserver/hubs"
	"sync"
)


// launch two goroutines
// one for gRPC service
// one for WebSocket service
// serverHub as a broker for communicating between gRPC service and Websocket service
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	serverHub := hubs.NewServerHub()

	go grpc.StartServing(serverHub, &wg)  
	go websocket.StartServing(serverHub, &wg)

	wg.Wait()
}