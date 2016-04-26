package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var (
	debugport = flag.Int("debugport", 44555, "Debugging port for net/trace")
	port      = flag.Int("port", 4050, "The server port")
)

type Event struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func receiveWebsocket(ws *websocket.Conn) {
	for {
		var event Event
		err := websocket.JSON.Receive(ws, &event)
		if err != nil {
			log.Println("Can't receive:", err.Error())
			break
		} else {
			// Echo the event back as JSON
			err := websocket.JSON.Send(ws, event)
			if err != nil {
				fmt.Println("Can't send:", err.Error())
				break
			}
		}
	}
}

func main() {
	http.Handle("/ws", websocket.Handler(receiveWebsocket))
	http.Handle("/", http.FileServer(http.Dir("static/html")))

	log.Printf("Server listening on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
