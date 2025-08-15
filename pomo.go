package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pomo timer")

	config := TestConfig()

	fmt.Printf("Config: %+v\n", config)

	maxRounds := config.rounds

	for rounds := 1; rounds <= maxRounds; rounds++ {
		fmt.Printf("Round %d/%d\n", rounds, maxRounds)
		fmt.Printf("Starting work for %d minutes\n", config.workDuration)
		countdownTimer(config.workDuration)
		if rounds%4 == 0 {
			fmt.Printf("Starting long rest for %d minutes\n", config.longRestDuration)
			countdownTimer(config.longRestDuration)
		} else {
			fmt.Printf("Starting rest for %d minutes\n", config.restDuration)
			countdownTimer(config.restDuration)
		}
	}

}
