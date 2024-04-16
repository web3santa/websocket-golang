package main

import (
	"log"
	"net/http"
)

func serverIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "/usr/share/nginx/html/index.html")
}

func main() {

	hub := NewHub()

	go hub.Run()

	http.HandleFunc("/", serverIndex)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
