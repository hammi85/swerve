package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// NewDynamoDB creates a new instance
func NewDynamoDB(c *DynamoConnection) (*DynamoDB, error) {
	ddb := &DynamoDB{}

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(c.Region),
		Endpoint: aws.String(c.Endpoint),
	})

	if err != nil {
		return nil, err
	}

	ddb.Session = sess
	ddb.Service = dynamodb.New(sess)
	return ddb, nil
}
