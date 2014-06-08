package main

import (
	"fmt"
	"os"
)

func main() {
	// Check for script file
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: %s SCRIPT", os.Args[0])
	}
	os.exis
	fmt.Println(os.Args)
}
