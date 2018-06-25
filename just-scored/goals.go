package justscored

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/structs"
	"github.com/pawelad/just-scored/worldcup"
)

const (
	Away = "away"
	Home = "home"
)

// Goal holds all relevant World Cup goal information,
// based on worldcup.Match and worldcup.TeamEvent structure
type Goal struct {
	EventID int // e.g. 123 (Hash Key)

	Player     string // e.g. "Bender RODRIGUEZ"
	PlayerTeam string // "home" or "away"
	GoalTime   string // e.g. "90'+2'"
	GoalType   string // "goal", "goal-penalty" or "goal-own"

	MatchID  string
	AwayTeam map[string]interface{}
	HomeTeam map[string]interface{}

	CreatedAt          time.Time
	Processed          bool
	NotificationSentAt interface{} // time.Time if sent, nil otherwise
}

// GetTeamInfo returns scorer team info
func (goal Goal) GetTeamInfo(team string) *map[string]interface{} {
	switch team {
	case Away:
		return &goal.AwayTeam
	case Home:
		return &goal.HomeTeam
	default:
		return nil
	}
}

// SetDBValue updates goal DynamoDB field with passed value
func (goal Goal) SetDBValue(field string, value interface{}) error {
	// TODO: Process multiple fields
	table := getDynamoDBTable()

	err := table.Update("EventID", goal.EventID).
		Set(field, value).
		Run()

	if err != nil {
		log.Print(err)
	}

	return err
}

// ToSlackMessage formats goal data to a Slack compatible string format
func (goal Goal) ToSlackMessage() string {
	var goalMessage, scoreMessage string
	var team, scorerTeam, oppositeTeam *map[string]interface{}

	// Find out scorer team
	switch goal.PlayerTeam {
	case Away:
		scorerTeam = goal.GetTeamInfo(Away)
		oppositeTeam = goal.GetTeamInfo(Home)
	case Home:
		scorerTeam = goal.GetTeamInfo(Home)
		oppositeTeam = goal.GetTeamInfo(Away)
	default:
		log.Panicf("goal.PlayerTeam must be either 'home' or 'away', got '%s'", goal.PlayerTeam)
	}

	// Format goal message
	switch goal.GoalType {
	case worldcup.RegularGoal:
		team = scorerTeam
		goalMessage = "*%s* %s just scored for *%s*"
	case worldcup.PenaltyGoal:
		team = scorerTeam
		goalMessage = "*%s* %s (P) just scored for *%s*"
	case worldcup.OwnGoal:
		// Scored for the opposite team
		team = oppositeTeam
		goalMessage = "*%s* %s (OG) just scored for *%s*"
	}
	goalMessage = fmt.Sprintf(goalMessage, goal.Player, goal.GoalTime, (*team)["Country"])

	// Format match score message
	scoreMessage = fmt.Sprintf(
		"%s %d - %d %s",
		goal.AwayTeam["Code"], int(goal.AwayTeam["Goals"].(float64)),
		int(goal.HomeTeam["Goals"].(float64)), goal.HomeTeam["Code"],
	)

	return fmt.Sprintf(":soccer: %s\n\n:loudspeaker: %s", goalMessage, scoreMessage)
}

// GetMatchGoals parses passed worldcup.Match and returns a list of its goals
func GetMatchGoals(match *worldcup.Match) (goals []*Goal) {
	log.Printf("Parsing match '%+v'", match)

	goalDefaults := &Goal{
		MatchID:  match.FifaID,
		AwayTeam: structs.Map(match.AwayTeam),
		HomeTeam: structs.Map(match.HomeTeam),

		CreatedAt:          time.Now().UTC(),
		Processed:          false,
		NotificationSentAt: nil,
	}

	teamEvents := map[string]*[]worldcup.TeamEvent{
		Away: &match.AwayTeamEvents,
		Home: &match.HomeTeamEvents,
	}

	for team, events := range teamEvents {
		for _, event := range *events {
			if event.IsGoal() {
				goal := *goalDefaults

				goal.EventID = event.ID
				goal.Player = event.Player
				goal.PlayerTeam = team
				goal.GoalTime = event.Time
				goal.GoalType = event.TypeOfEvent

				goals = append(goals, &goal)
			}
		}
	}

	return goals
}
