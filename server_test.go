package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pippin": 20,
			"Mary":   10,
		}}

	server := &PlayerServer{&store}

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request, _ := newGetScoreRequest("404")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound
		assertStatus(t, got, want)
	})

	t.Run("returns Pippin's score", func(t *testing.T) {
		request, _ := newGetScoreRequest("Pippin")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Mary's score", func(t *testing.T) {
		request, _ := newGetScoreRequest("Mary")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
}

func newPostScoreRequest(name string) (*http.Request, error) {
	return http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
}

func assertStatus(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got status %d, want %d", got, want)
	}
}
func assertResponseBody(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{},
	}
	server := &PlayerServer{store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request, _ := newPostScoreRequest("Pippin")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})
}
