package http_server

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"
)

type StubPlayerStore struct {
  scores   map[string]int
  winCalls []string
}

func (s StubPlayerStore) GetPlayerScore(name string) int {
  score := s.scores[name]
  return score
}

func (s *StubPlayerStore) RecordWin(name string) {
  s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
  store := StubPlayerStore{
    map[string]int{
      "Pepper": 20,
      "Floyd":  10,
    },
    nil,
  }
  server := &PlayerServer{&store}

  t.Run("returns Pepper's score", func(t *testing.T) {
    req := newGetScoreRequest("Pepper")
    res := httptest.NewRecorder()

    server.ServeHTTP(res, req)

    assertStatus(t, res.Code, http.StatusOK)
    assertResponseBody(t, res.Body.String(), "20")
  })

  t.Run("returns Floyd's score", func(t *testing.T) {
    req := newGetScoreRequest("Floyd")
    res := httptest.NewRecorder()

    server.ServeHTTP(res, req)

    assertStatus(t, res.Code, http.StatusOK)
    assertResponseBody(t, res.Body.String(), "10")
  })

  t.Run("returns 404 on missing players", func(t *testing.T) {
    req := newGetScoreRequest("Apollo")
    res := httptest.NewRecorder()

    server.ServeHTTP(res, req)

    assertStatus(t, res.Code, http.StatusNotFound)
  })
}

func TestStoreWins(t *testing.T) {
  store := StubPlayerStore{
    map[string]int{},
    nil,
  }
  server := &PlayerServer{&store}

  t.Run("it returns accepted on POST", func(t *testing.T) {
    player := "Pepper"
    req := newPostWinRequest(player)
    res := httptest.NewRecorder()

    server.ServeHTTP(res, req)

    assertStatus(t, res.Code, http.StatusAccepted)

    if len(store.winCalls) != 1 {
      t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
    }

    if store.winCalls[0] != player {
      t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
    }
  })
}

func newPostWinRequest(name string) *http.Request {
  req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
  return req
}

func assertStatus(t testing.TB, got, want int) {
  t.Helper()
  if got != want {
    t.Errorf("did not get correct status, got %d, want %d", got, want)
  }
}

func newGetScoreRequest(name string) *http.Request {
  req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
  return req
}

func assertResponseBody(t testing.TB, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("response body is wrong, got %q want %q", got, want)
  }
}
