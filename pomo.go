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

		timer("work", config.workDuration)

		if rounds%4 == 0 {
			timer("longRest", config.longRestDuration)
		} else {
			timer("shortRest", config.restDuration)
		}
	}

	fmt.Println("All rounds finished")
}
