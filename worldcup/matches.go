package worldcup

import (
	"log"
	"time"
)

// Team represents a single team resource received from http://worldcup.sfg.io/
type Team struct {
	Code    string `json:"code"`
	Country string `json:"country"`
	Goals   int    `json:"goals"`
}

// TeamStatistics represents a single match team statistics resource received from http://worldcup.sfg.io/
type TeamStatistics struct {
	AttemptsOnGoal  int         `json:"attempts_on_goal"`
	BallPossession  int         `json:"ball_possession"`
	BallsRecovered  int         `json:"balls_recovered"`
	Blocked         int         `json:"blocked"`
	Clearances      int         `json:"clearances"`
	Corners         int         `json:"corners"`
	Country         string      `json:"country"`
	DistanceCovered int         `json:"distance_covered"`
	FoulsCommitted  interface{} `json:"fouls_committed"`
	NumPasses       int         `json:"num_passes"`
	Offsides        int         `json:"offsides"`
	OffTarget       int         `json:"off_target"`
	OnTarget        int         `json:"on_target"`
	PassAccuracy    int         `json:"pass_accuracy"`
	PassesCompleted int         `json:"passes_completed"`
	RedCards        int         `json:"red_cards"`
	Tackles         int         `json:"tackles"`
	Woodwork        int         `json:"woodwork"`
	YellowCards     int         `json:"yellow_cards"`
}

// Match represents a single match resource received from http://worldcup.sfg.io/
type Match struct {
	AwayTeam           Team           `json:"away_team"`
	AwayTeamCountry    string         `json:"away_team_country"`
	AwayTeamEvents     []TeamEvent    `json:"away_team_events"`
	AwayTeamStatistics TeamStatistics `json:"away_team_statistics"`
	Datetime           time.Time      `json:"datetime"`
	FifaID             string         `json:"fifa_id"`
	HomeTeam           Team           `json:"home_team"`
	HomeTeamCountry    string         `json:"home_team_country"`
	HomeTeamEvents     []TeamEvent    `json:"home_team_events"`
	HomeTeamStatistics TeamStatistics `json:"home_team_statistics"`
	LastEventUpdateAt  interface{}    `json:"last_event_update_at"`
	LastScoreUpdateAt  interface{}    `json:"last_score_update_at"`
	Location           string         `json:"location"`
	Status             string         `json:"status"`
	Time               string         `json:"time"`
	Venue              string         `json:"venue"`
	Winner             string         `json:"winner"`
	WinnerCode         string         `json:"winner_code"`
}

// getMultipleMatches is a helper method that processes an API endpoint
// that returns a list of matches
func (c *Client) getMultipleMatches(endpoint string) ([]*Match, error) {
	request, err := c.NewRequest("GET", endpoint)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var matches []*Match
	_, err = c.Do(request, &matches)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	if len(matches) == 0 {
		// No matches are played right now
		return nil, nil
	}

	return matches, nil
}

// GetCurrentMatches returns the currently played matches,
// if any are played at the moment, otherwise nil
func (c *Client) GetCurrentMatches() ([]*Match, error) {
	return c.getMultipleMatches("matches/current")
}

// GetTodaysMatches returns all today's matches,
// if any are happening, otherwise nil
func (c *Client) GetTodaysMatches() ([]*Match, error) {
	return c.getMultipleMatches("matches/today")
}
