package repository

import "busManager/models"

type IBusStopRepository interface {
	GetById(id string) (*models.BusStop, error)
	GetByName(name string) (*models.BusStop, error)
	Add(stop *models.BusStop) error
	DeleteById(id string) error
	GetAll() ([]models.BusStop, error)
	GetAllByRouteId() ([]models.BusStop, error)
	UpdateById(stop *models.BusStop) error
}
