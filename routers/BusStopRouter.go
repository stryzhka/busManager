package routers

import (
	"busManager/controller"
	"busManager/repository"
	"busManager/service"
	"context"
)

type BusStopRouter struct {
	ctx               context.Context
	BusStopController controller.BusStopController
}

func NewBusStopRouter() (*BusStopRouter, error) {
	router := &BusStopRouter{}
	repo, err := repository.NewSqliteBusStopRepository("db.db")
	if err != nil {
		return nil, err
	}
	srv := service.NewBusStopService(repo)
	router.BusStopController = *controller.NewBusStopController(*srv)
	return router, nil
}

func (a *BusStopRouter) Startup(ctx context.Context) {
	a.ctx = ctx
}
