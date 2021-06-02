package main

import (
	"log"
	"net/http"
)

// Run the Application -> InMemoryPlayerStore

type InMemoryPlayerStore struct {}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}