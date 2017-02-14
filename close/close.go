package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()
	port := ":8000"

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		w.Write([]byte("Finished!\n"))
	}))

	srv := &http.Server{Addr: port, Handler: mux}

	go func() {
		log.Printf("Running on port %v", port)
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()

	time.Sleep(5 * time.Second)
	log.Println("Shutting Down")
	srv.Close()
}
