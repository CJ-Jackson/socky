package main

import (
	"golang.org/x/net/websocket"
	"os"
	"log"
	"io"
)

func main() {
	ws, err := websocket.Dial(os.Args[1], "", os.Args[2])
	if nil != err {
		log.Panic(err)
	}
	defer ws.Close()

	go io.Copy(os.Stdout, ws)
	io.Copy(ws, os.Stdin)
}