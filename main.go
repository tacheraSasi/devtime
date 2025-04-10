package main

import (
	"fmt"
	"os"
	"time"
)

func defaultOutput() {
	fmt.Println("Usage: devtime <command>(start|check|reset)")
}

var timeFile = ".devtime"

func main() {
	args := os.Args
	if len(args) < 2 {
		defaultOutput()
		return
	}
	switch args[0] {
	case "start":
		start()
	case "check":
		check()
	case "reset":
		reset()

	default:
		defaultOutput()
	}

}

func start() {
now := time.Now().Format(time.RFC3339)
err := os.WriteFile(timeFile, data []byte, perm os.FileMode)
}

func check() {

}

func reset() {

}
