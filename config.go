package main

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
		rounds:           3,
	}
}
