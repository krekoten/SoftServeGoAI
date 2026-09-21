package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Fprintln(os.Stderr, "Usage: greeting_ai <name>")
		os.Exit(1)
	}

	fmt.Printf("Hello %s!\n", os.Args[1])
}
