package main

import (
	"github.com/thoj/go-ircevent"
)

const COMMANDER = "Hamcha"

func Execute(actions []Action) {
	// Get the character list
	chars, master := CharacterList(actions)

	// Create and connect all the characters
	var actors = make(map[string]*irc.Connection)
	for i := range chars {
		conn := irc.IRC(chars[i], chars[i])
		conn.UseTLS = false
		conn.Connect("localhost:6667")
		conn.Join("#trials")
		actors[chars[i]] = conn
	}
	actors[master].AddCallback("PRIVMSG", commands)
	actors[master].Loop()
}

func commands(e *irc.Event) {
	if e.Nick != COMMANDER {
		return
	}
}
