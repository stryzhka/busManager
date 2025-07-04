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

func (a *BusStopRouter) GetById(id string) string {
	data := a.BusStopController.GetById(id)
	return data
}

func (a *BusStopRouter) GetByName(name string) string {
	data := a.BusStopController.GetByName(name)
	return data
}

func (a *BusStopRouter) GetAll() string {
	data := a.BusStopController.GetAll()
	return data
}

func (a *BusStopRouter) Add(busStopData string) string {
	data := a.BusStopController.Add(busStopData)
	return data
}

func (a *BusStopRouter) DeleteById(id string) string {

	return a.BusStopController.DeleteById(id)
}

func (a *BusStopRouter) UpdateById(busStopData string) string {
	return a.BusStopController.UpdateById(busStopData)
}
