package controllers

import (
	"golang.org/x/net/websocket"
	"github.com/astaxie/beego"
)

var MsgChan []chan string

// Echo the data received on the WebSocket.
func GroupChatServer(ws *websocket.Conn) {
	ws.Write([]byte("连接成功"))
	busch := make(chan string,1)
	MsgChan = append(MsgChan,busch)
	mych := make(chan string,1)
	go waitio(ws, mych)
	for {
		select {
		case msg := <-busch:
			ws.Write([]byte(msg))
		case msg, ok := <-mych:
			if ok {
				forward(msg)
			}else {
				removech(busch)
				close(busch)
				return
			}
		default:
		}
	}
}

func waitio(conn *websocket.Conn, mych chan string) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			close(mych)
			return
		}
		mych <- string(buffer[:n])
	}
}

func forward(msg string) {
	for _, ch := range MsgChan {
		ch <- msg
	}
}

func removech(ch chan string) {
	beego.Debug("removech...")
	var msgChan []chan string
	for _, msgchan := range MsgChan {
		if ch != msgchan {
			msgChan = append(msgChan, msgchan)
		}
	}
	MsgChan = msgChan
}