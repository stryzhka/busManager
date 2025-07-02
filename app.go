package main

import (
	"busManager/controller"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
	//busRepo     repository.IBusStopRepository
	//driverRepo  repository.IDriverRepository
	//busStopRepo repository.IBusStopRepository
	//
	//busService     service.IBusService
	//driverService  service.IDriverService
	//busStopService service.IBusStopService

	BusController     controller.BusController
	DriverController  controller.DriverController
	BusStopController controller.BusStopController
	RouteController   controller.RouteController
}

// NewApp creates a new App application struct
func NewApp() (*App, error) {
	app := &App{}

	return app, nil
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
