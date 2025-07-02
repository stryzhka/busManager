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
