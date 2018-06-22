package justscored

import (
	"fmt"
	"log"
	"time"

	"github.com/pawelad/just-scored/worldcup"
)

// Goal holds all relevant World Cup goal information,
// based on worldcup.Match and worldcup.TeamEvent structure
type Goal struct {
	EventID   int       // Hash key
	CreatedAt time.Time // Range key

	Player     string // e.g. "Bender RODRIGUEZ"
	PlayerTeam string // e.g. "FOO"
	GoalTime   string // e.g. "90'+2'"
	IsPenalty  bool

	MatchID string
	Match   string // e.g. "FOO - BAR"
	Score   string // e.g. "0 - 0"

	Processed          bool
	NotificationSentAt interface{} // time.Time if sent, nil otherwise
}

// ParseMatchEvents parses passed worldcup.Match and adds all goals to DynamoDB
func ParseMatchEvents(match *worldcup.Match) {
	log.Printf("Parsing match '%+v'", match)

	goal := Goal{
		MatchID: match.FifaID,
		Match:   fmt.Sprintf("%s - %s", match.AwayTeam.Code, match.HomeTeam.Code),
		Score:   fmt.Sprintf("%d - %d", match.AwayTeam.Goals, match.HomeTeam.Goals),

		CreatedAt:          time.Now().UTC(),
		Processed:          false,
		NotificationSentAt: nil,
	}

	for _, event := range match.AwayTeamEvents {
		if event.IsGoal() {
			goal.EventID = event.ID
			goal.Player = event.Player
			goal.PlayerTeam = match.AwayTeam.Code
			goal.GoalTime = event.Time
			goal.IsPenalty = event.IsPenalty()

			addGoal(&goal)
		}
	}

	for _, event := range match.HomeTeamEvents {
		if event.IsGoal() {
			goal.EventID = event.ID
			goal.Player = event.Player
			goal.PlayerTeam = match.HomeTeam.Code
			goal.GoalTime = event.Time
			goal.IsPenalty = event.IsPenalty()

			addGoal(&goal)
		}
	}
}
