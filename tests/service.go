package main

import (
	"net/http"

	"golang.org/x/net/websocket"
	"fmt"
	"io"
)

var totle int
var b []byte
// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	totle ++
	fmt.Println(totle, ws.LocalAddr())
	io.Copy(ws,ws)
}

// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/ws", websocket.Handler(EchoServer))
	err := http.ListenAndServe("192.168.1.150:8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}