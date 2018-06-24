package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/bluele/slack"
	"github.com/pawelad/just-scored/just-scored"
)

// UnmarshalStreamImage converts events.DynamoDBAttributeValue to struct
// Taken from: https://stackoverflow.com/a/50017398/3023841
func UnmarshalStreamImage(attribute map[string]events.DynamoDBAttributeValue, out interface{}) error {
	dbAttrMap := make(map[string]*dynamodb.AttributeValue)

	for k, v := range attribute {
		var dbAttr dynamodb.AttributeValue

		bytes, marshalErr := v.MarshalJSON()
		if marshalErr != nil {
			return marshalErr
		}

		json.Unmarshal(bytes, &dbAttr)
		dbAttrMap[k] = &dbAttr
	}

	return dynamodbattribute.UnmarshalMap(dbAttrMap, out)

}

// SendGoalNotification sends a scored goal notification to passed Slack webhook URL
func SendGoalNotification(url string, goal *justscored.Goal) (bool, error) {
	slackWebhook := slack.NewWebHook(url)

	webhookPayload := slack.WebHookPostPayload{
		Text:      goal.ToSlackMessage(),
		IconEmoji: ":soccer:",
	}

	err := slackWebhook.PostMessage(&webhookPayload)
	if err != nil {
		log.Print(err)
		return false, err
	}

	// Update goal processed status
	err = goal.SetValue("Processed", true)
	if err != nil {
		log.Print(err)
	}

	return true, nil
}

// Handler is the AWS Lambda entry point.
// It's invoked when a goal is added to a DynamoDB.
func Handler(ctx context.Context, event events.DynamoDBEvent) {
	var goals []*justscored.Goal

	// SLACK_WEBHOOK_URLS is a comma separated list of webhook URLs
	slackWebhookURLs := strings.Split(os.Getenv("SLACK_WEBHOOK_URLS"), ",")

	// Convert DynamoDB event to a list of goals
	for _, record := range event.Records {
		log.Printf("Processing request data for event ID %s, type %s.", record.EventID, record.EventName)

		goal := justscored.Goal{}
		err := UnmarshalStreamImage(record.Change.NewImage, &goal)
		if err != nil {
			log.Print(err)
		}

		// Only add non-processed goals
		if goal.Processed == false {
			goals = append(goals, &goal)
		}
	}

	// Send all goal notifications to all provided Slack webhook URLs
	for _, url := range slackWebhookURLs {
		for _, goal := range goals {
			log.Printf("Sending goal %d Slack notification to %v", goal.EventID, url)
			// TODO: Use goroutines and channels
			SendGoalNotification(url, goal)
		}
	}
}

func main() {
	lambda.Start(Handler)
}
