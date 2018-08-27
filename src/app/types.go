package app

// Application model
type Application struct {
	Config Configuration
}

type DynamoConnection struct {
	Host      string
	User      string
	Password  string
	TableName string
	Region    string
}

type Configuration struct {
	HTTPListener  string
	HTTPSListener string
	APIListener   string
	DynamoDB      DynamoConnection
}
