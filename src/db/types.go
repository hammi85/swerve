package db

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDB model
type DynamoDB struct {
	Session *session.Session
	Service *dynamodb.DynamoDB
}

// DynamoConnection model
type DynamoConnection struct {
	Endpoint  string
	User      string
	Password  string
	TableName string
	Region    string
}
