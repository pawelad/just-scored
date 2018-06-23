package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pawelad/just-scored/just-scored"
	"github.com/pawelad/just-scored/worldcup"
)

// Handler is the AWS Lambda entry point.
// It checks for all goals in currently played World Cup match and saves them to DynamoDB
func Handler() (string, error) {
	client := worldcup.NewClient()
	matches, err := client.GetCurrentMatches()

	if err != nil {
		log.Print(err)

		return fmt.Sprintf("Error: %v", err), err
	}

	if matches == nil {
		msg := "No matches are being played right now"
		log.Printf(msg)

		return msg, nil
	}

	addedGoals := 0
	for _, match := range matches {
		goals := justscored.GetMatchGoals(match)
		addedGoals += justscored.AddGoals(goals)
		log.Printf("Match %s was successfully parsed", match.FifaID)
	}

	return fmt.Sprintf("%d goals were added", addedGoals), nil
}

func main() {
	lambda.Start(Handler)
}
