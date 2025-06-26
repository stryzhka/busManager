package repository

import "busManager/models"

type BusRepository interface {
	GetById(id string) (*models.Bus, error)
	GetByNumber(number string) (*models.Bus, error)
	Add(bus *models.Bus) error
	DeleteById(id string) error
	GetAll() []models.Bus
	UpdateById(bus *models.Bus) error
}
