package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"time"
)


func main() {
	origin := "http://localhost/"
	url := "ws://localhost:8080/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	for i:=0;i<10;i++{
		if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
			log.Fatal(err)
		}
		var msg = make([]byte, 512)
		var n int
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received: %s.\n", msg[:n])
		time.Sleep(time.Second*2)
	}
	ws.Close()
}


