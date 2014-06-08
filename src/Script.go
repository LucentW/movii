package main

import (
	"fmt"
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
		return out
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
		out.Who = line[1:]
		return out
	}

	// +Character - JOIN
	if line[0] == '+' {
		out.Type = ACTION_JOIN
		out.Who = line[1:]
		return out
	}

	// -Character - LEAVE
	if line[0] == '-' {
		out.Type = ACTION_LEAVE
		out.Who = line[1:]
		return out
	}

	// PAUSE
	if line == "PAUSE" {
		out.Type = ACTION_PAUSE
		return out
	}

	// ** EVENT **
	if line[0] == '*' {
		out.Type = ACTION_EVENT
		out.What = line
		return out
	}

	fmt.Println("Unparsable line: " + line)
	out.Type = ACTION_NULL
	return out
}

func ParseScript(script string) []Action {
	lines := strings.Split(string(script), "\n")
	actions := make([]Action, len(lines))
	for i := range actions {
		actions[i] = ParseLine(lines[i])
	}
	return actions
}

func CharacterList(actions []Action) []string {
	chars := make(map[string]bool)

	for _, v := range actions {
		if v.Type == ACTION_JOIN || v.Type == ACTION_MASTER {
			chars[v.Who] = true
		}
	}

	out := make([]string, len(chars))
	i := 0
	for k, _ := range chars {
		out[i] = k
		i++
	}
	return out
}
