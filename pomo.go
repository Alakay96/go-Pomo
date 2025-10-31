package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/glebarez/go-sqlite"
)

type SessionStats struct {
	totalWorkSessions int
	// totalWorkTime     int
	totalRestSessions int
	// totalRestTime     int
}

func main() {

	db, err := sql.Open("sqlite", "./pomo.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	fmt.Println("Connected to DB")

	_, err = CreateTable(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Table created")

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

	statsID, err := Insert(db, &currentSession)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Stats ID: %d\n", statsID)
}

func CreateTable(db *sql.DB) (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS stats (
	id INTEGER PRIMARY KEY,
	rest INTEGER NOT NULL,
	work INTEGER NOT NULL
	);`

	return db.Exec(sql)
}

func Insert(db *sql.DB, s *SessionStats) (int64, error) {
	sql := `INSERT INTO stats (work, rest)
	VALUES (?, ?)`

	result, err := db.Exec(sql, s.totalWorkSessions, s.totalRestSessions)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
