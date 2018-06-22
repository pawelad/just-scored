package justscored

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// getDynamoDBTable returns initialized dynamo.Table object
// The table name is taken from DYNAMODB_TABLE environment variable
func getDynamoDBTable() dynamo.Table {
	db := dynamo.New(session.New())
	table := db.Table(os.Getenv("DYNAMODB_TABLE"))

	return table
}

// addGoal adds passed goal to DynamoDB
func addGoal(goal *Goal) (added bool, err error) {
	table := getDynamoDBTable()

	// Check if the goal is already in the DB
	count, err := table.Get("EventID", goal.EventID).Count()
	if err != nil {
		log.Print(err)
	}

	if count != 0 {
		log.Printf("Goal already added: '%+v'", goal)
		return false, nil
	}

	log.Printf("Adding goal '%+v'", goal)
	err = table.Put(goal).Run()

	if err != nil {
		log.Print(err)
		return false, err
	}

	return true, nil
}
