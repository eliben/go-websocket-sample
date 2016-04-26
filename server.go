// Server-side part of the Go websocket sample.
//
// Eli Bendersky [http://eli.thegreenplace.net]
// This code is in the public domain.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/trace"
	"golang.org/x/net/websocket"
)

var (
	port = flag.Int("port", 4050, "The server port")
)

type Event struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// handleWebsocketMessage handles the message e arriving on connection ws.
func handleWebsocketMessage(ws *websocket.Conn, e Event) error {
	// Log the request with net.Trace
	tr := trace.New("websocket.Receive", "receive")
	defer tr.Finish()
	tr.LazyPrintf("Got event %v\n", e)

	// Echo the event back as JSON
	err := websocket.JSON.Send(ws, e)
	if err != nil {
		return fmt.Errorf("Can't send: %s", err.Error())
	}
	return nil
}

// websocketConnection handles a single websocket connection - ws.
func websocketConnection(ws *websocket.Conn) {
	for {
		var event Event
		err := websocket.JSON.Receive(ws, &event)
		if err != nil {
			log.Println("Can't receive:", err.Error())
			break
		} else {
			if err := handleWebsocketMessage(ws, event); err != nil {
				log.Println(err.Error())
				break
			}
		}
	}
}

func main() {
	// Set up websocket server and static file server. In addition, we're using
	// net/trace for debugging - it will be available at /debug/requests.
	http.Handle("/ws", websocket.Handler(websocketConnection))
	http.Handle("/", http.FileServer(http.Dir("static/html")))

	log.Printf("Server listening on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
