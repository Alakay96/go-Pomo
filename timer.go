package main

import (
	"fmt"
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
