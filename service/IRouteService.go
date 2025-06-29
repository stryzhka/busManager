package service

import "busManager/models"

type IRouteService interface {
	GetById(id string) (*models.Route, error)
	GetByNumber(number string) (*models.Route, error)
	Add(route *models.Route) error
	DeleteById(id string) error
	GetAll() ([]models.Route, error)
	UpdateById(route *models.Route) error
	AssignDriver(routeId, driverId string) error
	AssignBusStop(routeId, busStopId string) error
	AssignBus(routeId, busId string) error
	// TODO: getall for all models, unassign
}
