package app

// Setup the application configuration
func (a *Application) Setup() {
	a.Config.FromEnv()
	a.Config.FromParameter()
}

// Run the application
func (a *Application) Run() {

}

// NewApplication creates new instance
func NewApplication() *Application {
	return &Application{
		Config: NewConfiguration(),
	}
}
