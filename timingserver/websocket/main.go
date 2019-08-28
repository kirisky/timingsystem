package websocket

import (
	"flag"
	"net/http"
	"log"
	"sync"

	"timingsystem/timingserver/hubs"
)

var addr = flag.String("addr", ":50052", "websocket service address")

var latestRecords = make([]byte, 0)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println("URL:", r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "websocket/home.html")
}

// StartServing starts websocket serving
func StartServing(serverHub *hubs.ServerHub, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("start websocket service")

	flag.Parse()

	wsHub := newHub(serverHub)
	go wsHub.run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("URL:", r.URL)
		serveWs(wsHub, w, r)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	
}