package configuration

// DynamoConnection model
type DynamoConnection struct {
	Endpoint  string
	User      string
	Password  string
	TableName string
	Region    string
}

// Configuration model
type Configuration struct {
	HTTPListener  string
	HTTPSListener string
	APIListener   string
	DynamoDB      DynamoConnection
}
