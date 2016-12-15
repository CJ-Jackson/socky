package socky

import (
	"fmt"
	"net/http"

	"github.com/CJ-Jackson/socky/src/config"
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/options"
	"github.com/julienschmidt/httprouter"
)

type sockyHttpCommand struct {
	mux     *httprouter.Router
	address string
}

func newSockyHttpCommand() *sockyHttpCommand {
	return &sockyHttpCommand{
		mux:     muxer,
		address: ":8080",
	}
}

func (sC *sockyHttpCommand) CommandConfigure(c *cli.Command) {
	c.SetName("socky:start:server").
		SetDescription("Run Socky Server").
		AddOption("address", "Listening address", options.NewString(&sC.address))
}

func (sC *sockyHttpCommand) CommandExecute() {
	for _, data := range config.GetConfig().SocketList {
		sC.mux.GET(data.Path, getSockyHttp(data.Protocol, data.Address))
	}

	fmt.Printf("Running Socky Server on '%s' (Ctrl + C to exit)...", sC.address)
	fmt.Println()

	http.ListenAndServe(sC.address, sC.mux)
}

func init() {
	cli.RegisterCommand(newSockyHttpCommand())
}
