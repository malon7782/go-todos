package main

import (
	"os"
)

var commands = []cmd_struct{
	{"add", Add},
	{"show", Show},
	{"drop", Drop},
	{"clear", Clear},
}
var list = [](*item){}

func main() {
	userArgs := os.Args[1:]
	if len(userArgs) == 0 {
		PrintHelpMsg()
		return
	}
	for _, command := range commands {
		if command.name == userArgs[0] {
			command.executor(userArgs[1:])
			return
		}
	}
	PrintHelpMsg()
}
