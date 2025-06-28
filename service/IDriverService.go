package service

import "busManager/models"

type IDriverService interface {
	GetById(id string) (*models.Driver, error)
	GetByPassportSeries(series string) (*models.Driver, error)
	Add(driver *models.Driver) error
	DeleteById(id string) error
	GetAll() []models.Driver
	UpdateById(driver *models.Driver) error
}
