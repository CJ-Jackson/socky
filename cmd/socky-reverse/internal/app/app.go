package app

import (
	"io"
	"log"
	"net"

	"github.com/CJ-Jackson/socky/cmd/socky-reverse/internal/config"
	"golang.org/x/net/websocket"
)

func getPortForwardClosure(value config.SocketValue) func() {
	return func() {
		listener, err := net.Listen(value.ListenType, value.ListenAddress)
		if nil != err {
			log.Fatal(err)
		}
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			go doPortForward(value, conn)
		}
	}
}

func doPortForward(value config.SocketValue, conn net.Conn) {
	defer conn.Close()

	ws, err := websocket.Dial(value.OriginUrl, value.OriginProtocol, value.OriginOrigin)
	if nil != err {
		log.Println(err)
		return
	}
	defer ws.Close()

	go io.Copy(conn, ws)
	io.Copy(ws, conn)
}

func Start() {
	fns := []func(){}
	for _, data := range config.GetConfig().SocketList {
		fns = append(fns, getPortForwardClosure(data))
	}

	first, rest := fns[0], fns[1:]
	for _, fn := range rest {
		go fn()
	}
	first()
}
