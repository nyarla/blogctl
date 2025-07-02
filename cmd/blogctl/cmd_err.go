package main

import (
	"fmt"
	"os"
	"strings"
)

func showUnsupportedCommand() {
	fmt.Printf("Unkown command: %s\n\n", strings.Join(os.Args[1:], " "))
	fmt.Print(help)
	os.Exit(1)
}
