package main

import (
	"net/http"
	"time"

	sse "github.com/r3labs/sse/v2"
)

func main() {
	server := sse.New()
	server.CreateStream("messages")

	mux := http.NewServeMux()
	mux.HandleFunc("/events/", func(w http.ResponseWriter, r *http.Request) {
		go func () {
			nowPlus10Sec := time.Now().Add(10 * time.Second)
			for time.Now().Before(nowPlus10Sec) {
				time.Sleep(time.Second)
				server.Publish("messages", &sse.Event{
					Data: []byte("ping"),
				})
			}
		}()

		server.ServeHTTP(w, r)
	})

	http.ListenAndServe(":8080", mux)
}
