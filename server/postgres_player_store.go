package main

import (
	"database/sql"
	"fmt"

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
	sql := `INSERT INTO score (player_name) VALUES($1)`
	_, err := p.store.Exec(sql, name)
	if err != nil {
		panic(err)
	}
}

func (p *PostgresPlayerStore) GetPlayerScore(name string) int {
	var score int
	sql := `SELECT COUNT(*) FROM score WHERE player_name=$1`
	row := p.store.QueryRow(sql, name)
	err := row.Scan(&score)
	if err != nil {
		panic(err)
	}
	return score
}

func (p *PostgresPlayerStore) GetAllPlayers() map[string]int {
	return nil
}