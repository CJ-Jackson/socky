package main

import (
	"fmt"

	_ "github.com/CJ-Jackson/socky/cmd/socky/internal"
	"github.com/cjtoolkit/cli"
)

func main() {
	fmt.Println("Socky:")
	fmt.Println()

	cli.Run()
}
