package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

var (
	serverport = flag.Int("serverport", 4050, "The server port")
)

type Event struct {
	X int `json:"x"`
	Y int `json:"y"`
}

var base_addr = "ws://localhost"
var origin = "http://localhost"

func checkWsEcho() {
	addr := fmt.Sprintf("%s:%d/wsecho", base_addr, *serverport)
	conn, err := websocket.Dial(addr, "", origin)
	if err != nil {
		log.Fatal("websocket.Dial error", err)
	}
	e := Event{
		X: 42,
		Y: 123456,
	}
	err = websocket.JSON.Send(conn, e)
	if err != nil {
		log.Fatal("websocket.JSON.Send error", err)
	}

	var reply Event
	err = websocket.JSON.Receive(conn, &reply)
	if err != nil {
		log.Fatal("websocket.JSON.Receive error", err)
	}

	if reply != e {
		log.Fatalf("reply != e: %s != %s", reply, e)
	}
	if err = conn.Close(); err != nil {
		log.Fatal("conn.Close error", err)
	}
}

func main() {
	checkWsEcho()
	//checkWsTime()
}
