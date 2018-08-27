package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hammi85/swerve/src/tls"
)

// Setup the application configuration
func (a *Application) Setup() {
	a.Config.FromEnv()
	a.Config.FromParameter()
}

// Run the application
func (a *Application) Run() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	// run the https listener
	httpsServer := tls.NewTLSServer(a.Config.HTTPSListener)
	go func() {
		log.Fatal(httpsServer.Listen())
	}()

	<-sigchan
}

// NewApplication creates new instance
func NewApplication() *Application {
	return &Application{
		Config: NewConfiguration(),
	}
}
