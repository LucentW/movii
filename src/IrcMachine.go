package main

import (
	"./irc"
	"fmt"
	"time"
)

const (
	SERVER    = "localhost:6667"
	COMMANDER = "Hamcha"
	CHANNEL   = "#trials"
	PAUSE     = time.Second * 3
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
		sinfo.Perform = []string{"PRIVMSG NickServ :identify PASSWORDHERE\r\n"}
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
				go play(actors, actions, master)
			}
		} else {
			fmt.Println(message)
		}
	}
}

func play(actors map[string]*irc.Client, actions []Action, master string) {
	var lastPerson string
	for _, act := range actions {
		switch act.Type {
		case ACTION_MASTER, ACTION_JOIN:
			actors[act.Who].Send(irc.Message{
				Command: "JOIN",
				Target:  CHANNEL,
			})
			time.Sleep(PAUSE / 4)
		case ACTION_LEAVE:
			actors[act.Who].Send(irc.Message{
				Command: "PART",
				Target:  CHANNEL,
				Text:    "Leaving..",
			})
			time.Sleep(PAUSE / 4)
		case ACTION_SAY:
			if act.What[0] == '*' {
				act.What = "\002"+act.What
			}
			if act.What[0] == '>' {
				act.What = "\0033"+act.What[1:]
				time.Sleep(time.Second)
			}
			actors[act.Who].Send(irc.Message{
				Command: "PRIVMSG",
				Target:  CHANNEL,
				Text:    act.What,
			})
			if act.What[len(act.What)-2:] == "--" {
				continue
			}
			if act.Who == lastPerson {
				time.Sleep(PAUSE - time.Second)
			} else {
				time.Sleep(PAUSE)
			}
			lastPerson = act.Who
		case ACTION_PAUSE:
			actors[master].Send(irc.Message{
				Command: "PRIVMSG",
				Target:  CHANNEL,
				Text:    "\001ACTION ** Court is now in recess for 5 minutes **\001",
			})
			time.Sleep(time.Minute * 4)
			actors[master].Send(irc.Message{
				Command: "PRIVMSG",
				Target:  CHANNEL,
				Text:    "\001ACTION ** Court will reconvene in a minute **\001",
			})
			time.Sleep(time.Minute)
		case ACTION_EVENT:
			actors[master].Send(irc.Message{
				Command: "PRIVMSG",
				Target:  CHANNEL,
				Text:    "\001ACTION " + act.What + "\001",
			})
			time.Sleep(PAUSE / 2)
		case ACTION_NULL:
			continue
		}
	}
}
