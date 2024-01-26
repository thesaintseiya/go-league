package poker_test

import (
	"strings"
	"testing"

	poker "github.com/thesaintseiya/league"
)

func TestCLI(t *testing.T) {
	t.Run("record pippin win from stdin", func(t *testing.T) {
		in := strings.NewReader("Pippin wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Pippin")
	})

	t.Run("record mary win from stdin", func(t *testing.T) {
		in := strings.NewReader("Mary wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Mary")
	})
}
