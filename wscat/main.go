package main

import (
	"io"
	"log"
	"os"

	"golang.org/x/net/websocket"
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
