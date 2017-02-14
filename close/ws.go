package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

func echoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	port := ":8000"

	mux := http.NewServeMux()

	mux.Handle("/echo", websocket.Handler(echoHandler))
	mux.Handle("/", http.FileServer(http.Dir(".")))

	srv := &http.Server{Addr: port, Handler: mux}

	go func() {
		log.Printf("Running on port %v", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Close Server
	time.Sleep(5 * time.Second)
	log.Println("Shutting Down")
	srv.Close()

	// Keep application going
	time.Sleep(10000 * time.Second)
	log.Println("Completely done")
}
