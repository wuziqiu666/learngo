package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowSever := makeDelayedServer(20 * time.Millisecond)
		fastSever := makeDelayedServer(0 * time.Millisecond)
		defer slowSever.Close()
		defer fastSever.Close()

		slowURL := slowSever.URL
		fastURL := fastSever.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("returns an error if a server donesn't reponse with in 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expect an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
