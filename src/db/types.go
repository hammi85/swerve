package db

import (
	"time"

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

// DomainList db entry
type DomainList struct {
	Domains []Domain `json:"domains"`
}

// Domain db entry
type Domain struct {
	ID           string `json:"id"`
	Name         string `json:"domain"`
	Redirect     string `json:"redirect"`
	Certificate  string `json:"certificate"`
	RedirectCode int    `json:"code"`
	Description  string `json:"description"`
	Created      time.Time
}
