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
	EventID    int    // Hash key
	Player     string // e.g. "Bender RODRIGUEZ"
	PlayerTeam string // e.g. "FOO"
	GoalTime   string // e.g. "90'+2'"
	IsPenalty  bool

	MatchID string
	Match   string // e.g. "FOO - BAR"
	Score   string // e.g. "0 - 0"

	CreatedAt          time.Time   // Range key
	Processed          bool        // Whether the notification was sent
	NotificationSentAt interface{} // time.Time if sent, nil otherwise
}

// GetMatchGoals parses passed worldcup.Match and returns a list of its goals
func GetMatchGoals(match *worldcup.Match) (goals []*Goal) {
	log.Printf("Parsing match '%+v'", match)

	goalDefaults := &Goal{
		MatchID: match.FifaID,
		Match:   fmt.Sprintf("%s - %s", match.AwayTeam.Code, match.HomeTeam.Code),
		Score:   fmt.Sprintf("%d - %d", match.AwayTeam.Goals, match.HomeTeam.Goals),

		CreatedAt:          time.Now().UTC(),
		Processed:          false,
		NotificationSentAt: nil,
	}

	teamEvents := map[string]*[]worldcup.TeamEvent{
		match.AwayTeam.Code: &match.AwayTeamEvents,
		match.HomeTeam.Code: &match.HomeTeamEvents,
	}

	for team, events := range teamEvents {
		for _, event := range *events {
			if event.IsGoal() {
				goal := *goalDefaults

				goal.EventID = event.ID
				goal.Player = event.Player
				goal.PlayerTeam = team
				goal.GoalTime = event.Time
				goal.IsPenalty = event.IsPenalty()

				goals = append(goals, &goal)
			}
		}
	}

	return goals
}
