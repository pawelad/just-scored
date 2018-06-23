package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pawelad/just-scored/just-scored"
	"github.com/pawelad/just-scored/worldcup"
)

// Response holds Lambda response data
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

// Handler checks for all goals in currently played World Cup match and saves them to DynamoDB
func Handler() (Response, error) {
	client := worldcup.NewClient()
	matches, err := client.GetCurrentMatches()

	if err != nil {
		log.Print(err)
		return Response{
			Message: fmt.Sprintf("Error: %v", err),
			Ok:      false,
		}, nil
	}

	if matches == nil {
		msg := "No matches are being played right now"
		log.Printf(msg)
		return Response{
			Message: msg,
			Ok:      true,
		}, nil
	}

	addedGoals := 0
	for _, match := range matches {
		goals := justscored.GetMatchGoals(match)
		addedGoals += justscored.AddGoals(goals)
		log.Printf("Match %s was successfully parsed", match.FifaID)
	}

	return Response{
		Message: fmt.Sprintf("%d goals were added", addedGoals),
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
