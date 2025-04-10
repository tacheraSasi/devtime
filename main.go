package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func defaultOutput() {
	fmt.Println("Usage: devtime <command>(start|check|reset)")
}

var timeFile = ".devtime"

func main() {
	args := os.Args

	// fmt.Println(args)
	if len(args) < 2 {
		defaultOutput()
		return
	}
	switch args[1] {
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
	err := os.WriteFile(timeFile, []byte(now), 0644)
	if err != nil {
		fmt.Errorf("Geez Something went wrong(Failed to write the file)", err)
		return
	}
	fmt.Println("devtime started the tracking")
}

func check() {
	data, err := os.ReadFile(timeFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No start time found. Run devtime start first before proceeding")
			return
		}

		fmt.Errorf("devtime failed to retrieve the data", err)
		return
	}

	startTimeStr := strings.TrimSpace(string(data))
	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		fmt.Println("Invalid time format in the data stored")
		return
	}

	duration := time.Since(startTime)
	fmt.Println("You have been coding for", duration)

}

func reset() {
	err := os.Remove(timeFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("not timer to reset")
			defaultOutput()
			return
		}
		fmt.Println("Failed to reset", err)
		return
	}

	fmt.Println("Timer is reset")

}
