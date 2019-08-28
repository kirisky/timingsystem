package main

import (
	"timingsystem/timingserver/websocket"
	"timingsystem/timingserver/grpc"
	"timingsystem/timingserver/hubs"
	"sync"
)



func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	serverHub := hubs.NewServerHub()

	go grpc.StartServing(serverHub, &wg)  
	go websocket.StartServing(serverHub, &wg)

	wg.Wait()
}