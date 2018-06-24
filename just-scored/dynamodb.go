package justscored

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// getDynamoDBTable returns initialized dynamo.Table object.
// The table name is taken from DYNAMODB_TABLE environment variable.
func getDynamoDBTable() dynamo.Table {
	awsSession, err := session.NewSessionWithOptions(session.Options{
		// This option makes use of the '~/.aws/config' file
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		log.Print(err)
	}

	db := dynamo.New(awsSession)
	table := db.Table(os.Getenv("DYNAMODB_TABLE"))

	return table
}

// AddGoal adds passed goal to DynamoDB
func AddGoal(goal *Goal) error {
	table := getDynamoDBTable()

	// Check if the goal is already in the DB
	count, err := table.Get("EventID", goal.EventID).Count()
	if err != nil {
		log.Print(err)
	}

	if count != 0 {
		log.Printf("Goal already added: '%+v'", goal)
		return nil
	}

	log.Printf("Adding goal '%+v'", goal)
	err = table.Put(goal).Run()

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// AddGoals adds passed goals to DynamoDB
func AddGoals(goals []*Goal) (addedGoals int) {
	for _, goal := range goals {
		err := AddGoal(goal)

		if err == nil {
			addedGoals++
		}
	}

	return addedGoals
}
