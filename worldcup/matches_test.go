package worldcup

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentMatches(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	// Correct response - no errors and a []*Match type returned
	mux.HandleFunc("/matches/current", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("current_matches.json"))

		assert.Equal(t, r.Method, "GET")
	})

	matches, err := client.GetCurrentMatches()

	assert.Nil(t, err)
	assert.NotNil(t, matches)
	assert.IsType(t, []*Match{}, matches)
	for _, match := range matches {
		assert.NotNil(t, match.FifaID)
	}
}

func TestGetCurrentMatchesEmpty(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	// Empty response - no errors and nil returned
	mux.HandleFunc("/matches/current", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[]`)

		assert.Equal(t, r.Method, "GET")
	})

	matches, err := client.GetCurrentMatches()

	assert.Nil(t, err)
	assert.Nil(t, matches)
}

func TestGetTodaysMatchesCorrect(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	// Expected response - no errors and a []*Match type returned
	mux.HandleFunc("/matches/today", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("todays_matches.json"))

		assert.Equal(t, r.Method, "GET")
	})

	matches, err := client.GetTodaysMatches()

	assert.Nil(t, err)
	assert.NotNil(t, matches)
	assert.IsType(t, []*Match{}, matches)
	for _, match := range matches {
		assert.NotNil(t, match.FifaID)
	}
}

func TestGetTodaysMatchesEmpty(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	// Empty response - no errors and nil returned
	mux.HandleFunc("/matches/today", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[]`)

		assert.Equal(t, r.Method, "GET")
	})

	matches, err := client.GetTodaysMatches()

	assert.Nil(t, err)
	assert.Nil(t, matches)
}
