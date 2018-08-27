package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hammi85/swerve/src/api"

	"github.com/hammi85/swerve/src/configuration"
	"github.com/hammi85/swerve/src/db"
	"github.com/hammi85/swerve/src/tls"
)

// Setup the application configuration
func (a *Application) Setup() {
	a.Config.FromEnv()
	a.Config.FromParameter()

	var err error
	a.DynamoDB, err = db.NewDynamoDB(&a.Config.DynamoDB)
	if err != nil {
		log.Fatal(err)
	}
}

// Run the application
func (a *Application) Run() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	// fetch domains from db

	// run the https listener
	httpsServer := tls.NewServer(a.Config.HTTPSListener)
	go func() {
		log.Fatal(httpsServer.Listen())
	}()

	// run the api listener
	apiServer := api.NewServer(a.Config.APIListener, a.DynamoDB)
	go func() {
		log.Fatal(apiServer.Listen())
	}()

	// wait for signals
	<-sigchan
}

// NewApplication creates new instance
func NewApplication() *Application {
	return &Application{
		Config: configuration.NewConfiguration(),
	}
}
