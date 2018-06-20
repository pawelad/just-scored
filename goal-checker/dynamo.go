package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/pawelad/just-scored/worldcup"
)

// Goal holds all relevant World Cup goal information,
// based on worldcup.Match and worldcup.TeamEvent structure
type Goal struct {
	EventID    int    // e.g. 1
	Player     string // e.g. "Bender RODRIGUEZ"
	PlayerTeam string // e.g. "FOO"
	GoalTime   string // e.g. "90'+2'"

	MatchID string // e.g. "1"
	Match   string // e.g. "FOO - BAR"
	Score   string // e.g. "0 - 0"

	CreatedAt          time.Time
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

		CreatedAt:          time.Now(),
		Processed:          false,
		NotificationSentAt: nil,
	}

	for _, event := range match.AwayTeamEvents {
		if event.TypeOfEvent == "goal" {
			goal.EventID = event.ID
			goal.Player = event.Player
			goal.PlayerTeam = match.AwayTeam.Code
			goal.GoalTime = event.Time

			addGoal(&goal)
		}
	}

	for _, event := range match.HomeTeamEvents {
		if event.TypeOfEvent == "goal" {
			goal.EventID = event.ID
			goal.Player = event.Player
			goal.PlayerTeam = match.HomeTeam.Code
			goal.GoalTime = event.Time

			addGoal(&goal)
		}
	}
}

// addGoal adds passed goal to DynamoDB
func addGoal(goal *Goal) {
	db := dynamo.New(session.New())
	table := db.Table(os.Getenv("DYNAMODB_TABLE"))

	log.Printf("Adding goal '%+v'", goal)

	// TODO: Don't update existing goals
	err := table.Put(goal).Run()

	if err != nil {
		// TODO: Proper error handling
		log.Print(err)
	}
}
