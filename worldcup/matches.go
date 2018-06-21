package worldcup

import (
	"errors"
	"time"
)

// A Team represents a single team resource received from http://worldcup.sfg.io/
type Team struct {
	Code    string `json:"code"`
	Country string `json:"country"`
	Goals   int    `json:"goals"`
}

// A Match represents a single match resource received from http://worldcup.sfg.io/
type Match struct {
	AwayTeam          Team        `json:"away_team"`
	AwayTeamEvents    []TeamEvent `json:"away_team_events"`
	Datetime          time.Time   `json:"datetime"`
	FifaID            string      `json:"fifa_id"`
	HomeTeam          Team        `json:"home_team"`
	HomeTeamEvents    []TeamEvent `json:"home_team_events"`
	LastEventUpdateAt interface{} `json:"last_event_update_at"`
	LastScoreUpdateAt interface{} `json:"last_score_update_at"`
	Location          string      `json:"location"`
	Status            string      `json:"status"`
	Time              string      `json:"time"`
	Venue             string      `json:"venue"`
	Winner            string      `json:"winner"`
	WinnerCode        string      `json:"winner_code"`
}

// GetCurrentMatch returns the current match, if a match is happening, otherwise nil
func (c *Client) GetCurrentMatch() (*Match, error) {
	request, err := c.NewRequest("GET", "matches/current")
	if err != nil {
		return nil, err
	}

	// The API returns a one-element list
	var matches []*Match
	_, err = c.Do(request, &matches)

	if err != nil {
		return nil, err
	}

	if len(matches) == 0 {
		// No matches are played right now
		return nil, nil
	} else if len(matches) > 1 {
		// TODO: A custom error?
		err := errors.New("the API returned more then 1 current matches")
		return nil, err
	}

	return matches[0], err
}

// GetTodaysMatches returns today's matches, if any are happening, otherwise nil
func (c *Client) GetTodaysMatches() ([]*Match, error) {
	request, err := c.NewRequest("GET", "matches/today")
	if err != nil {
		return nil, err
	}

	var matches []*Match
	_, err = c.Do(request, &matches)

	if len(matches) == 0 {
		// No matches are played right now
		return nil, nil
	}

	return matches, err
}
