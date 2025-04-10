package actions

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tacheraSasi/devtime.git/utils"
)

var timeFile = ".devtime"

func Start() {
	now := time.Now().Format(time.RFC3339)
	err := os.WriteFile(timeFile, []byte(now), 0644)
	utils.IsErr("Failed to write the file", err)
	fmt.Println("devtim started tracking at", now)

}
func Check() {
	data, err := os.ReadFile(timeFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No start time found. RUN devtime start to start")
			return
		}
		fmt.Println(err)
		return

	}

	startTimeStr := strings.TrimSpace(string(data))
	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	utils.IsErr("Invalid format", err)

	duration := time.Since(startTime)
	fmt.Println("You have been coding for", duration)

}

func Reset() {
	err := os.Remove(timeFile)
	utils.IsErr("Failed to remove the file", err)
	fmt.Println("Timer is reset")

}
