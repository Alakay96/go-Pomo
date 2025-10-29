package main

import (
	"fmt"
	"strings"
)

type SessionStats struct {
	totalWorkSessions int
	// totalWorkTime     int
	totalRestSessions int
	// totalRestTime     int
}

func main() {
	currentSession := SessionStats{0, 0}

	var configSelect string
	var configMode func() Config

	fmt.Println("Pomo timer")

	fmt.Println("Which config to use:")
	fmt.Println("D - default")
	fmt.Println("C - custom")
	fmt.Println("T - test")
	fmt.Scanln(&configSelect)
	configSelect = strings.ToUpper(configSelect)

	switch configSelect {
	case "T":
		configMode = TestConfig
	case "D":
		configMode = DefaultConfig
	case "C":
		configMode = CustomConfig
	}

	config := configMode()

	fmt.Printf("Config: %+v\n", config)

	maxRounds := config.rounds

	for rounds := 1; rounds <= maxRounds; rounds++ {
		fmt.Printf("Round %d/%d\n", rounds, maxRounds)

		timer("work", config.workDuration)
		currentSession.totalWorkSessions++

		if rounds%4 == 0 {
			timer("longRest", config.longRestDuration)
			currentSession.totalRestSessions++
		} else {
			timer("shortRest", config.restDuration)
			currentSession.totalRestSessions++
		}
	}

	fmt.Println("All rounds finished")
	fmt.Printf(`Session Stats
	Total work sessions: %d
	Total rest sessions: %d`,
		currentSession.totalWorkSessions, currentSession.totalRestSessions)
}
