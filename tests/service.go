package main

import (
	"net/http"

	"golang.org/x/net/websocket"
	"fmt"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	fmt.Println(ws.Request().Header)
}

// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/ws", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}