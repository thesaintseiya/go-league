package main

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanup := createTempFile(t, `[
			{"Name": "Pippin", "Wins": 10},
			{"Name": "Mary", "Wins": 33}]
		`)
		defer cleanup()

		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()
		want := []Player{
			{"Pippin", 10},
			{"Mary", 33},
		}

		assertLeague(t, got, want)

		// read again to make sure read position is getting offset
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanup := createTempFile(t, `[
			{"Name": "Pippin", "Wins": 10},
			{"Name": "Mary", "Wins": 33}]
		`)
		defer cleanup()

		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Mary")
		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanup := createTempFile(t, `[
			{"Name": "Pippin", "Wins": 10},
			{"Name": "Mary", "Wins": 33}]
		`)
		defer cleanup()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pippin")

		got := store.GetPlayerScore("Pippin")
		want := 11

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanup := createTempFile(t, `[
			{"Name": "Pippin", "Wins": 10},
			{"Name": "Mary", "Wins": 33}]
		`)
		defer cleanup()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Sam")

		got := store.GetPlayerScore("Sam")
		want := 1

		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
