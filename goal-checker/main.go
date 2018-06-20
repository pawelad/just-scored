package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
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
	match, err := client.GetCurrentMatch()

	if err != nil {
		return Response{
			Message: fmt.Sprintf("Error: %v", err),
			Ok:      false,
		}, nil
	}

	if match == nil {
		msg := "No match is being played right now"
		log.Printf(msg)

		return Response{
			Message: msg,
			Ok:      true,
		}, nil
	}

	// TODO: Return number of added goals?
	ParseMatchEvents(match)

	return Response{
		Message: fmt.Sprintf("Match %s was successfully parsed", match.FifaID),
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
