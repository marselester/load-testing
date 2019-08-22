// Program server is an HTTP server API that shows books with random latencies up to 200ms.
package main

import (
	"math/rand"
	"net/http"
	"time"
)

// Max API response latency in ms added to emulate random response time.
const maxLatency = 200

func main() {
	// Initialize the default source of uniformly-distributed pseudo-random ints
	// to add latencies to API responses.
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/v1/books", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(rand.Intn(maxLatency)) * time.Millisecond)
		w.Write([]byte(`{"books": [{"name": "The Adventures of Tom Sawyer"}]}`))
	})
	http.ListenAndServe(":8000", nil)
}
