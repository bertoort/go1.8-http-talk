package main

import "net/http"

func main() {
	port := ":8000"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// load the main page
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<a href="/images" > Images </a>
		`))

	}))

	http.Handle("/preload", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// server push is available if w implements http.Pusher
		if p, ok := w.(http.Pusher); ok {
			p.Push("/static/gopher1.png", nil)
		}

		// load the main page
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<a href="/images" > Images </a>
		`))

	}))

	http.Handle("/images", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// load the main page
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<img src="/static/gopher2.png" width="300"/>
			<br>
			<img src="/static/gopher1.png" width="300"/>
		`))

	}))

	http.ListenAndServeTLS(port, "cert.pem", "key.pem", nil)

}
