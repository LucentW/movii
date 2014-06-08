package main

import (
	"strings"
)

type ActionType int

const (
	ACTION_JOIN   ActionType = iota
	ACTION_MASTER            = iota
	ACTION_LEAVE             = iota
	ACTION_SAY               = iota
	ACTION_PAUSE             = iota
	ACTION_EVENT             = iota
	ACTION_NULL              = iota
)

type Action struct {
	Type ActionType
	Who  string
	What string
}

func ParseLine(line string) Action {
	var out Action
	line = strings.Trim(line, " \r\n")

	// Empty line - SKIP
	if len(line) < 1 {
		out.Type = ACTION_NULL
		return
	}

	// Character: Message - DIALOGUE
	if sep := strings.Index(line, ":"); sep > 0 {
		out.Type = ACTION_SAY
		out.Who = strings.Trim(line[:sep], " ")
		out.What = strings.Trim(line[sep+1:], " \r")
		return out
	}

	// =Character - JOIN MASTER
	if line[0] == '=' {
		out.Type = ACTION_MASTER
		out.Who = lines[1:]
	}

	// +Character - JOIN
	if line[0] == '+' {
		out.Type = ACTION_JOIN
		out.Who = lines[1:]
	}

	// -Character - LEAVE
	if line[0] == '-' {
		out.Type = ACTION_LEAVE
		out.Who = lines[1:]
	}

	// PAUSE
	if line == "PAUSE" {
		out.Type = ACTION_PAUSE
	}
}

func ParseScript(script string) []Action {
	lines := strings.Split(string(content), "\n")
	actions = make([]Action, len(lines))
	for i := range actions {
		actions[i] = ParseLine(lines[i])
	}
}
