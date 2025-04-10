package main

import (
	"fmt"
	"os"

	"github.com/tacheraSasi/devtime.git/actions"
)

func defaultOutput() {
	fmt.Println("Usage: devtime <command> (start|check|reset)")
}

func main() {
	args := os.Args
	if len(args) < 2 {
		defaultOutput()
		return
	}

	switch args[1] {
	case "start":
		actions.Start()
	case "check":
		actions.Check()
	case "reset":
		actions.Reset()
	default:
		defaultOutput()
	}

}
