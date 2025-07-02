package routers

import (
	"busManager/controller"
	"busManager/repository"
	"busManager/service"
	"context"
)

type BusRouter struct {
	ctx           context.Context
	BusController controller.BusController
}

func NewBusRouter() (*BusRouter, error) {
	router := &BusRouter{}
	repo, err := repository.NewSqliteBusRepository("db.db")
	if err != nil {
		return nil, err
	}
	service := service.NewBusService(repo)
	router.BusController = *controller.NewBusController(*service)
	return router, nil
}

func (a *BusRouter) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *BusRouter) GetById(id string) string {
	data := a.BusController.GetById(id)
	return data
}

func (a *BusRouter) GetByNumber(number string) string {
	data := a.BusController.GetByNumber(number)
	return data
}

func (a *BusRouter) GetAll() string {
	data := a.BusController.GetAll()
	return data
}

func (a *BusRouter) Add(busData string) string {
	data := a.BusController.Add(busData)
	return data
}

func (a *BusRouter) DeleteById(id string) string {

	return a.BusController.DeleteById(id)
}

func (a *BusRouter) UpdateById(busData string) string {
	return a.BusController.UpdateById(busData)
}
