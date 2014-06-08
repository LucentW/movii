package main

import (
	"fmt"
	//"github.com/thoj/go-ircevent"
)

func Execute(actions []Action) {
	chars := CharacterList(actions)
	fmt.Println(chars)
}
