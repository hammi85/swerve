package app

// Application model
type Application struct {
	Config *Configuration
}

// DynamoConnection model
type DynamoConnection struct {
	Host      string
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
