package main

import (
	"./irc"
)

const (
	SERVER    = "localhost:6667"
	COMMANDER = "Hamcha"
	CHANNEL   = "#trials"
)

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
		sinfo.Channels = []string{}
		sinfo.Perform = []string{}
		conn.ServerInfo = sinfo
		conn.Sid = chars[i]
		conn.ServerName = "Actor"
		err, ch := conn.Connect(SERVER)
		if err != nil {
			panic(err)
		}
		actors[chars[i]] = conn
		chans[chars[i]] = ch
	}

	for {
		message := <-chans[master]
		if message.Message.Source.Nickname == COMMANDER {
			if message.Message.Text == "play" {
				go play(actors, actions)
			}
		}
	}
}

func play(actors map[string]*irc.Client, actions []Action) {
	for _, act := range actions {
		if act.Type == ACTION_MASTER && act.Type == ACTION_JOIN {
			actors[act.Who].Send(irc.Message{
				Command: "JOIN",
				Target:  CHANNEL,
			})
		}
	}
}
