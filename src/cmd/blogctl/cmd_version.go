package main

import (
	"fmt"
	"os"
)

func showVersion() {
	fmt.Printf("blogctl version - %s\n", version)
	os.Exit(0)
}
