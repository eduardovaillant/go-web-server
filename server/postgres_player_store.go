package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresPlayerStore() *PostgresPlayerStore {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "players"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }

  err = db.Ping()
  if err != nil {
    panic(err)
  }
	return &PostgresPlayerStore{db}
}

type PostgresPlayerStore struct {
	store *sql.DB
}

func (p *PostgresPlayerStore) RecordWin(name string) {
	query := `INSERT INTO score (player_name) VALUES($1)`
	_, err := p.store.Exec(query, name)
	if err != nil {
		panic(err)
	}
}

func (p *PostgresPlayerStore) GetPlayerScore(name string) int {
	var score int
	query := `SELECT COUNT(*) FROM score WHERE player_name=$1`
	row := p.store.QueryRow(query, name)
	err := row.Scan(&score)
	if err != nil {
		panic(err)
	}
	return score
}

func (p *PostgresPlayerStore) GetAllPlayersScores() map[string]int {
	var result = make(map[string]int)
	query := `SELECT player_name, COUNT(player_name) FROM score GROUP BY player_name`
	rows, err := p.store.Query(query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var playerName string
		var score int
		if err := rows.Scan(&playerName, &score); err != nil {
			log.Fatal(err)
		}
		result[playerName] = score
	}

	return result
}