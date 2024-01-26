package main

import (
	"log"
	"net/http"

	"github.com/thesaintseiya/league"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	defer close()

	if err != nil {
		log.Fatal(err)
	}

	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5001", server))
}
