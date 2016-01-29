package main

import (
	"golang.org/x/net/websocket"
	"log"
	"time"
	"net"
)

func connect(localip string) {
	origin := "http://192.168.1.150/"
	url := "ws://192.168.1.150:8080/ws"
	//ws, err = websocket.Dial(url, "", origin)
	config,err := websocket.NewConfig(url, origin)
	if err != nil{
		log.Println(err)
		return
	}
	localaddr := &net.TCPAddr{IP:net.ParseIP(localip)}
	remoteaddr := &net.TCPAddr{IP:net.ParseIP("192.168.1.150"),Port:8080}
	client,err := net.DialTCP("tcp4",localaddr,remoteaddr)
	if err != nil{
		log.Println(err)
		return
	}
	_, err = websocket.NewClient(config,client)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(time.Second*100)
	return
}

func main() {
	for i := 0; i < 65535; i++ {
		go connect("192.168.1.101")
		go connect("192.168.1.102")
		go connect("192.168.1.150")
		log.Println(i)
	}
	time.Sleep(time.Second * 100)
}
