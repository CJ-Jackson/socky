package socky

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/CJ-Jackson/socky/cmd/socky/internal/config"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/websocket"
)

func reformat(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}

func isWebSocketReq(req *http.Request) bool {
	return "websocket" == reformat(req.Header.Get("Upgrade")) && "upgrade" == reformat(req.Header.Get("Connection"))
}

func getSocket(protocol, address string) websocket.Handler {
	return func(in *websocket.Conn) {
		out, err := net.Dial(protocol, address)
		if nil != err {
			log.Panic(err)
		}
		defer out.Close()

		go io.Copy(in, out)
		io.Copy(out, in)
	}
}

func getSockyHttp(protocol, address string) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, param httprouter.Params) {
		if isWebSocketReq(req) && "" != protocol && "" != address {
			getSocket(protocol, address).ServeHTTP(res, req)
			return
		}

		fmt.Fprintln(res, "Hello World")
	}
}

func Start(address string) {
	for _, data := range config.GetConfig().SocketList {
		muxer.GET(data.Path, getSockyHttp(data.Protocol, data.Address))
	}

	fmt.Printf("Running Socky Server on '%s' (Ctrl + C to exit)...", address)
	fmt.Println()

	http.ListenAndServe(address, muxer)
}
