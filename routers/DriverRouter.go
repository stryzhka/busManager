package routers

import (
	"busManager/controller"
	"busManager/repository"
	"busManager/service"
	"context"
)

type DriverRouter struct {
	ctx              context.Context
	DriverController controller.DriverController
}

func NewDriverRouter() (*DriverRouter, error) {
	router := &DriverRouter{}
	repo, err := repository.NewSqliteDriverRepository("db.db")
	if err != nil {
		return nil, err
	}
	srv := service.NewDriverService(repo)
	router.DriverController = *controller.NewDriverController(*srv)
	return router, nil
}

func (a *DriverRouter) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *DriverRouter) GetById(id string) string {
	data := a.DriverController.GetById(id)
	return data
}

func (a *DriverRouter) GetByPassportSeries(number string) string {
	data := a.DriverController.GetByPassportSeries(number)
	return data
}

func (a *DriverRouter) GetAll() string {
	data := a.DriverController.GetAll()
	return data
}

func (a *DriverRouter) Add(driverData string) string {
	data := a.DriverController.Add(driverData)
	return data
}

func (a *DriverRouter) DeleteById(id string) string {

	return a.DriverController.DeleteById(id)
}

func (a *DriverRouter) UpdateById(driverData string) string {
	return a.DriverController.UpdateById(driverData)
}
