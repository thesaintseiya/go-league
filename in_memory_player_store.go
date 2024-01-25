package main

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store  map[string]int
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.store[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.store[name]++
}

func (s *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range s.store {
		league = append(league, Player{Name: name, Wins: wins})
	}
	return league
}
