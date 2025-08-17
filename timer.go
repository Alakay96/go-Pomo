package main

import (
	"fmt"
	"strings"
	"time"
)

func countdownTimer(countTime int) {
	duration := time.Duration(countTime) * time.Minute
	startTime := time.Now()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		elapsed := time.Since(startTime)
		remaining := duration - elapsed

		if remaining <= 0 {
			fmt.Println("Time is up!")
			return
		}

		fmt.Printf("\rTime remaining: %02d:%02d  ", int(remaining.Minutes()), int(remaining.Seconds())%60)
	}
}

func timer(timerType string, duration int) {
	var userInput string
	var timeMessage string
	var timeDuration int
	switch timerType {
	case "work":
		timeMessage = "work"
		timeDuration = duration
	case "shortRest":
		timeMessage = "rest"
		timeDuration = duration
	case "longRest":
		timeMessage = "long rest"
		timeDuration = duration
	default:
		fmt.Println("Error: invalid timne")
	}

	for userInput != "Y" {
		fmt.Printf("Start %s for %d minutes? [Y/N]:  ", timeMessage, timeDuration)
		fmt.Print()
		fmt.Scanln(&userInput)
		userInput = strings.ToUpper(userInput)

		if userInput != "Y" {
			fmt.Println("Waiting")
		}
	}
	countdownTimer(duration)
}
