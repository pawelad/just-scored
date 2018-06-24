// Program goal-notifier is a AWS Lambda function that's invoked via DynamoDB stream events,
// processes the added / updated goal and sends Slack notifications to configred Slack webhook URLs.
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

// UnmarshalStreamImage is a helper function that unmarshalls a map of DynamoDB
// fields (i.e. record.Change.NewImage) onto passed struct pointer.
// It's needed because dynamodbattribute.UnmarshalMap takes dynamodb.AttributeValue
// and record.Change.NewImage returns events.DynamoDBAttributeValue.
//
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

// SendGoalNotification sends a scored goal notification to passed Slack
// webhook URL and sets it's DB processed value.
func SendGoalNotification(url string, goal *justscored.Goal) error {
	slackWebhook := slack.NewWebHook(url)

	webhookPayload := slack.WebHookPostPayload{
		Text:      goal.ToSlackMessage(),
		IconEmoji: ":soccer:",
	}

	err := slackWebhook.PostMessage(&webhookPayload)
	if err != nil {
		log.Print(err)
		return err
	}

	// Update goal processed status
	err = goal.SetDBValue("Processed", true)
	if err != nil {
		log.Print(err)
	}

	return nil
}

// Handler is the AWS Lambda entry point
func Handler(ctx context.Context, event events.DynamoDBEvent) {
	var goals []*justscored.Goal

	// SLACK_WEBHOOK_URLS is a comma separated list of webhook URLs
	slackWebhookURLs := strings.Split(os.Getenv("SLACK_WEBHOOK_URLS"), ",")

	// Convert the DynamoDB event to a list of goals
	for _, record := range event.Records {
		log.Printf("Processing request data for event ID '%s' (%s).", record.EventID, record.EventName)

		// We only care about inserts (and *maybe* updates)
		if record.EventName == string(events.DynamoDBOperationTypeRemove) {
			return
		}

		var goal justscored.Goal
		err := UnmarshalStreamImage(record.Change.NewImage, &goal)
		if err != nil {
			log.Print(err)
		}

		// Only add non-processed goals
		if goal.EventID != 0 && goal.Processed == false {
			goals = append(goals, &goal)
		}
	}

	// Send goals notifications to all configured Slack webhook URLs
	for _, url := range slackWebhookURLs {
		log.Printf("Slack Webhook URL: '%s'", url)

		for _, goal := range goals {
			log.Printf("Sending notification for goal '%+v'", goal)
			// TODO: Use goroutines and channels
			SendGoalNotification(url, goal)
		}
	}
}

func main() {
	lambda.Start(Handler)
}
