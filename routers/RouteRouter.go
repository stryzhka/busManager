package routers

import (
	"busManager/controller"
	"busManager/repository"
	"busManager/service"
	"context"
)

type RouteRouter struct {
	ctx             context.Context
	RouteController controller.RouteController
}

func NewRouteRouter() (*RouteRouter, error) {
	router := &RouteRouter{}
	busRepo, err := repository.NewSqliteBusRepository("db.db")
	if err != nil {
		return nil, err
	}
	busStopRepo, err := repository.NewSqliteBusStopRepository("db.db")
	if err != nil {
		return nil, err
	}
	driverRepo, err := repository.NewSqliteDriverRepository("db.db")
	if err != nil {
		return nil, err
	}
	routeRepo, err := repository.NewSqliteRouteRepository("db.db")
	if err != nil {
		return nil, err
	}
	routeService := service.NewRouteService(routeRepo, driverRepo, busRepo, busStopRepo)
	router.RouteController = *controller.NewRouteController(routeService)
	return router, err
}

func (a *RouteRouter) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *RouteRouter) GetById(id string) string {
	data := a.RouteController.GetById(id)
	return data
}

func (a *RouteRouter) GetByNumber(number string) string {
	data := a.RouteController.GetByNumber(number)
	return data
}

func (a *RouteRouter) GetAll() string {
	data := a.RouteController.GetAll()
	return data
}

func (a *RouteRouter) Add(routeData string) string {
	data := a.RouteController.Add(routeData)
	return data
}

func (a *RouteRouter) DeleteById(id string) string {

	return a.RouteController.DeleteById(id)
}

func (a *RouteRouter) UpdateById(routeData string) string {
	return a.RouteController.UpdateById(routeData)
}

func (a *RouteRouter) AssignDriver(routeId, driverId string) string {
	return a.RouteController.AssignDriver(routeId, driverId)
}

func (a *RouteRouter) AssignBusStop(routeId, busStopId string) string {
	return a.RouteController.AssignBusStop(routeId, busStopId)
}

func (a *RouteRouter) AssignBus(routeId, busId string) string {
	return a.RouteController.AssignBus(routeId, busId)
}

func (a *RouteRouter) UnassignDriver(routeId, driverId string) string {
	return a.RouteController.UnassignDriver(routeId, driverId)
}

func (a *RouteRouter) UnassignBusStop(routeId, busStopId string) string {
	return a.RouteController.UnassignBusStop(routeId, busStopId)
}

func (a *RouteRouter) UnassignBus(routeId, busId string) string {
	return a.RouteController.UnassignBus(routeId, busId)
}

func (a *RouteRouter) GetAllDriversById(routeId string) string {
	return a.RouteController.GetAllDriversById(routeId)
}

func (a *RouteRouter) GetAllBusesById(routeId string) string {
	return a.RouteController.GetAllBusesById(routeId)
}

func (a *RouteRouter) GetAllBusStopsById(routeId string) string {
	return a.RouteController.GetAllBusStopsById(routeId)
}
