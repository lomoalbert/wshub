package main

import (
	"golang.org/x/net/websocket"
	"log"
	"time"
	"fmt"
)

func connect() {
	origin := "http://127.0.0.1/"
	url := "ws://127.0.0.1:8080/ws"
	_, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	for i := 0; i < 100000; i++ {
		go connect()
		fmt.Println(i)
	}
	time.Sleep(time.Second * 100)
}
