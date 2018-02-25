package bootstrap

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	// "github.com/kataras/iris/sessions"
	// "github.com/kataras/iris/websocket"
	"github.com/telingadigital/goblog/config"
)

type Configurator func(*App)

type App struct {
	*iris.Application
	Config *config.Config
}

// New returns a new App
func New(cfgs ...Configurator) *App {
	appConfig := LoadConfig()
	app := &App{
		Config: appConfig,
		Application: iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(app)
	}
	return app
}

// Load config
func LoadConfig() *config.Config{
	return config.SetConfig()
}

// Configure accepts configurations and runs them inside App's context
func (app *App) Configure(cfgs ...Configurator) {
	for _, cfg := range cfgs {
		cfg(app)
	}
}

// Prepares the application
// returns itself
func (app *App) NewApp() *App {
	// middleware
	app.Use(recover.New())
	app.Use(logger.New())
	return app
}

// Listen starts the http server with the specified "addr"
func (app *App) Listen(cfgs ...iris.Configurator) {
	port := app.Config.Port
	app.Run(iris.Addr(port), cfgs...)
}