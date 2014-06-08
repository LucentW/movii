package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Check for script file
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: %s SCRIPT", os.Args[0])
	}

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	actions := ParseScript(string(content))
	for i := range actions {
		fmt.Printf("%+v\r\n", actions[i])
	}
}
