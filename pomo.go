package main

import (
	"fmt"
	"strings"
)

func main() {

	var startInput string
	var restInput string
	fmt.Println("Pomo timer")

	config := TestConfig()

	fmt.Printf("Config: %+v\n", config)

	maxRounds := config.rounds

	for rounds := 1; rounds <= maxRounds; rounds++ {
		fmt.Printf("Round %d/%d\n", rounds, maxRounds)

		for startInput != "Y" {
			fmt.Printf("Start work timer for %d minutes? [Y/N]:  ", config.workDuration)
			fmt.Print()
			fmt.Scanln(&startInput)
			startInput = strings.ToUpper(startInput)

			if startInput != "Y" {
				fmt.Println("Waiting")
			}
		}
		countdownTimer(config.workDuration)

		for restInput != "Y" {
			fmt.Printf("Start rest timer now? [Y/N]:  ")
			fmt.Scanln(&restInput)
			restInput = strings.ToUpper(restInput)

			if restInput != "Y" {
				fmt.Println("Waiting")
			}
		}
		if rounds%4 == 0 {
			fmt.Printf("Starting long rest for %d minutes\n", config.longRestDuration)
			countdownTimer(config.longRestDuration)
		} else {
			fmt.Printf("Starting rest for %d minutes\n", config.restDuration)
			countdownTimer(config.restDuration)
		}
	}
}
