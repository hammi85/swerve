package db

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	dbTableName = "Domains"
)

var (
	dbListAllDomains = &dynamodb.ScanInput{
		TableName: aws.String(dbTableName),
	}
	dbTableCreate = &dynamodb.CreateTableInput{
		TableName: aws.String(dbTableName),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("id"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("domain"), KeyType: aws.String("RANGE")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("id"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("domain"), AttributeType: aws.String("S")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}
	dbTableDescribe = &dynamodb.DescribeTableInput{
		TableName: aws.String(dbTableName),
	}
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

	ddb.prepareTable()

	return ddb, nil
}

// prepareTable checks for the main table
func (d *DynamoDB) prepareTable() {
	// get the table description
	_, err := d.Service.DescribeTable(dbTableDescribe)
	if err != nil {
		log.Println("Table 'Domains' didn't exists. Creating ...")
		_, cerr := d.Service.CreateTable(dbTableCreate)
		if cerr != nil {
			log.Fatal(cerr)
		}
		log.Println("Table 'Domains' created")
	}
}

// FetchAll items from domains table
func (d *DynamoDB) FetchAll() error {
	itemList, err := d.Service.Scan(dbListAllDomains)
	log.Printf("FetchAll %#v", itemList)
	if err != nil {
		return fmt.Errorf("Error while fetching domain items %v", err)
	}

	return nil
}

// InsertDomain stores a domain
func (d *DynamoDB) InsertDomain(domain Domain) error {
	mm, err := dynamodbattribute.MarshalMap(domain)

	if err != nil {
		return err
	}

	_, err = d.Service.PutItem(&dynamodb.PutItemInput{
		Item:      mm,
		TableName: aws.String(dbTableName),
	})

	return err
}
