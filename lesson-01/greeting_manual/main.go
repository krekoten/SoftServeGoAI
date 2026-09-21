package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Tell your name :)")
		return
	}

	name := os.Args[1]
	fmt.Printf("Hello %s!\n", name)
}
