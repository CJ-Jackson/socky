package socky

import "github.com/julienschmidt/httprouter"

var muxer = httprouter.New()

func init() {
	muxer.GET("/", getSockyHttp("", ""))
}