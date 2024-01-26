package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) Find(name string) *Player {
	for i, player := range l {
		if player.Name == name {
			return &l[i]
		}
	}
	return nil
}

func CreateLeague(db io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(db).Decode(&league)
	if err != nil {
		err = fmt.Errorf("error parsing league: %v", err)
	}
	return league, err
}
