package main

import "fmt"

type Config struct {
	workDuration     int
	restDuration     int
	longRestDuration int
	rounds           int
}

func DefaultConfig() Config {
	return Config{
		workDuration:     25,
		restDuration:     5,
		longRestDuration: 15,
		rounds:           4,
	}
}

func TestConfig() Config {
	return Config{
		workDuration:     1,
		restDuration:     1,
		longRestDuration: 2,
		rounds:           1,
	}
}

func CustomConfig() Config {
	var work, rest, longRest, rounds int
	fmt.Println("Work duration: ")
	fmt.Scanln(&work)
	fmt.Println("Rest duration: ")
	fmt.Scanln(&rest)
	fmt.Println("Long rest duration: ")
	fmt.Scanln(&longRest)
	fmt.Println("Number of rounds: ")
	fmt.Scanln(&rounds)
	return Config{
		workDuration:     work,
		restDuration:     rest,
		longRestDuration: longRest,
		rounds:           rounds,
	}
}
