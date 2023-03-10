package http_server

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestRecordingWinsRetrievingThem(t *testing.T) {
  store := NewInMemoryPlayerStore()
  server := PlayerServer{store}
  player := "Pepper"

  server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
  server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
  server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

  res := httptest.NewRecorder()
  server.ServeHTTP(res, newGetScoreRequest(player))
  assertStatus(t, res.Code, http.StatusOK)

  assertResponseBody(t, res.Body.String(), "3")
}
