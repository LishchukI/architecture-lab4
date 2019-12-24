package main

import (
	"architecture-lab4/engine"
	"bufio"
	"strings"
	"os"
)

func parse(commandLine string) engine.Command {
	text := strings.Fields(commandLine)
	if text[0] == "split" {
		if len(text) == 3 {
			str := text[1]
			sep := text[2]
				if (len(sep) == 1) {
						return &engine.PrintcCommand{Str: str, Sep: sep}
				} else {
						return &engine.PrintCommand{Arg: "SYNTAX ERROR: 'sep' need >1 length!"}
				}
		} else {
			return &engine.PrintCommand{Arg: "SYNTAX ERROR: need 3 arguments! 'split <str> <sep>'"}
		}
	} else if text[0] == "print" {
		if len(text) == 2 {
			return &engine.PrintCommand{Arg: text[1]}
		} else {
			return &engine.PrintCommand{Arg: "SYNTAX ERROR: need 2 arguments! 'print <str>'"}
		}
	} else {
		return &engine.PrintCommand{Arg: "SYNTAX ERROR: need 'split' or 'print'!"}
	}
}

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open("./example.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine)
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}
