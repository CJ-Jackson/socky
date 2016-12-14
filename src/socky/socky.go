package socky

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/websocket"
	"io"
	"net"
	"net/http"
	"strings"
	"log"
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