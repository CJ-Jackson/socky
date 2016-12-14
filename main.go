package main

import (
	"fmt"
	_ "github.com/CJ-Jackson/socky/src"
	"github.com/cjtoolkit/cli"
)

func main() {
	fmt.Println("Socky:")
	fmt.Println()

	cli.Run()
}
