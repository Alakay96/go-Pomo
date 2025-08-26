package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	work := intCheck("Work duration:	")
	rest := intCheck("Rest duration:	")
	longRest := intCheck("Long rest duration:	")
	rounds := intCheck("Number of rounds:	")
	return Config{
		workDuration:     work,
		restDuration:     rest,
		longRestDuration: longRest,
		rounds:           rounds,
	}
}

func intCheck(input string) int {
	var value int
	for {
		fmt.Println(input)
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}
		line = strings.TrimSpace(line)
		value, err = strconv.Atoi(line)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}
		return value
	}
}
