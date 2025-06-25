package repository

import "busManager/models"

type BusRepository interface {
	FindById(id string) (models.Bus, error)
	FindByNumber(number string) (models.Bus, error)
	Add(bus *models.Bus) error
	DeleteById(id string) error
	GetAll() []models.Bus
}
