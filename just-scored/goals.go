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
	EventID    int    // e.g. 123 (Hash Key)
	Player     string // e.g. "Bender RODRIGUEZ"
	PlayerTeam string // e.g. "FOO"
	GoalTime   string // e.g. "90'+2'"
	IsPenalty  bool

	MatchID    string // e.g. "123456"
	MatchScore string // e.g. "FOO 0 - 1 BAR"

	CreatedAt          time.Time
	Processed          bool
	NotificationSentAt interface{} // time.Time if sent, nil otherwise
}

// SetValue updates goal DynamoDB field with passed value
func (goal Goal) SetValue(field string, value interface{}) error {
	table := getDynamoDBTable()

	err := table.Update("EventID", goal.EventID).
		Set(field, value).
		Run()

	return err
}

// ToSlackMessage formats goal data to a Slack compatible string format
func (goal Goal) ToSlackMessage() string {
	// TODO: Take penalties and own goals into consideration
	message := ":soccer: *%v* (%v) just scored for *%v*\n\n"
	message += ":joystick: %v"
	message = fmt.Sprintf(message, goal.Player, goal.PlayerTeam, goal.MatchScore)

	return message
}

// GetMatchGoals parses passed worldcup.Match and returns a list of its goals
func GetMatchGoals(match *worldcup.Match) (goals []*Goal) {
	log.Printf("Parsing match '%+v'", match)

	goalDefaults := &Goal{
		MatchID: match.FifaID,
		MatchScore: fmt.Sprintf(
			"%s %d - %d %s",
			match.AwayTeam.Code, match.AwayTeam.Goals,
			match.HomeTeam.Goals, match.HomeTeam.Code,
		),

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
