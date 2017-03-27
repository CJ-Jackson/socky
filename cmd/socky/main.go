package main

import (
	"fmt"
	"os"

	_ "github.com/CJ-Jackson/socky/cmd/socky/internal"
	"github.com/CJ-Jackson/socky/cmd/socky/internal/socky"
)

func main() {
	fmt.Println("Socky:")
	fmt.Println()

	if len(os.Args) < 2 {
		return
	}

	socky.Start(os.Args[1])
}
