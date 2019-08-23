package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":50051", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Response) { 
		serveWs(hub, w, r) 
	})
	
	// serveWsl, err := net.Listen("tcp4", *addr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Fatal(http.Serve(l, nil))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}