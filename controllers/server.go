package controllers

import (
	"io"
	"golang.org/x/net/websocket"
	"fmt"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	fmt.Println(ws.LocalAddr(),ws.RemoteAddr(),ws.PayloadType,ws.Request().Header)
	ws.Write([]byte("连接成功"))
	io.Copy(ws, ws)
}