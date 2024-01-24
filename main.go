package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
	winners []string
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.winners = append(s.winners, name)
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5001", server))
}
