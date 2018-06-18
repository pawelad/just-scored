package worldcup

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentMatchCorrect(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	// Correct response - no errors and a *Match type returned
	mux.HandleFunc("/matches/current", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, fixture("current_match.json"))

		assert.Equal(t, r.Method, "GET")
	})

	match, err := client.GetCurrentMatch()

	assert.Nil(t, err)
	assert.NotNil(t, match)
	assert.IsType(t, &Match{}, match)
	assert.NotNil(t, match.FifaID)
}

func TestGetCurrentMatchEmpty(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	// Empty response - no errors and nil returned
	mux.HandleFunc("/matches/current", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[]`)

		assert.Equal(t, r.Method, "GET")
	})

	match, err := client.GetCurrentMatch()

	assert.Nil(t, err)
	assert.Nil(t, match)
}

func TestGetCurrentMatchMultiple(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	// Multiple items - an error and nil returned
	mux.HandleFunc("/matches/current", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"fifa_id":1},{"fifa_id":2}]`)

		assert.Equal(t, r.Method, "GET")
	})

	match, err := client.GetCurrentMatch()

	assert.NotNil(t, err)
	assert.Nil(t, match)
}

func TestGetTodaysMatchesCorrect(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	// Correct response - no errors and a []*Match type returned
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
	assert.NotNil(t, matches[0].FifaID)
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
