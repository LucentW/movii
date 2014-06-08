package main

import (
	"./irc"
	"fmt"
)

const COMMANDER = "Hamcha"

func Execute(actions []Action) {
	// Get the character list
	chars, master := CharacterList(actions)

	// Create and connect all the characters
	var actors = make(map[string]*irc.Client)
	var chans = make(map[string]chan irc.ClientMessage)
	for i := range chars {
		conn := new(irc.Client)
		var sinfo irc.Server
		sinfo.Username = chars[i]
		sinfo.Nickname = chars[i]
		sinfo.Altnick = chars[i] + "`"
		sinfo.Realname = chars[i]
		sinfo.Channels = []string{"#trials"}
		conn.ServerInfo = sinfo
		conn.Sid = chars[i]
		conn.ServerName = "Actor"
		err, ch := conn.Connect("localhost:6667")
		if err != nil {
			panic(err)
		}
		actors[chars[i]] = conn
		chans[chars[i]] = ch
	}

	for {
		message := <-chans[master]
		fmt.Printf("%v+\r\n", message)
	}
}
