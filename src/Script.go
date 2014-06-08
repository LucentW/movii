package main

import (
	"strings"
)

type ActionType int

const (
	ACTION_ENTER ActionType = iota
	ACTION_LEFT             = iota
	ACTION_SAY              = iota
	ACTION_PAUSE            = iota
	ACTION_EVENT            = iota
	ACTION_NULL             = iota
)

type Action struct {
	Type ActionType
	Who  string
	What string
}

func ParseLine(line string) Action {
	var out Action
	line = strings.Trim(line, " \r\n")

	if len(line) < 1 {
		out.Type = ACTION_NULL
		return
	}

	if sep := strings.Index(line, ":"); sep > 0 {
		out.Type = ACTION_SAY
		out.Who = strings.Trim(line[:sep], " ")
		out.What = strings.Trim(line[sep+1:], " \r")
		return out
	}
}

func ParseScript(script string) []Action {
	lines := strings.Split(string(content), "\n")
	actions = make([]Action, len(lines))
	for i := range actions {
		actions[i] = ParseLine(lines[i])
	}
}
