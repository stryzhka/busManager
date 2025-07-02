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
