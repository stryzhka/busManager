package service

import (
	"busManager/models"
	"busManager/repository"
	"errors"
)

type BusService struct {
	repo repository.BusRepository
}

func NewBusService(r repository.BusRepository) *BusService {
	b := &BusService{r}
	return b
}

func (bs *BusService) FindById(id string) (*models.Bus, error) {

	bus, err := bs.repo.FindById(id)
	if bus == nil {
		return bus, errors.New("Bus not found")
	}
	return nil, err
}

func (bs *BusService) FindByNumber(number string) (*models.Bus, error) {
	bus, err := bs.repo.FindByNumber(number)
	if bus == nil {
		return bus, errors.New("Bus not found")
	}
	return nil, err
}

func (bs *BusService) Add(bus *models.Bus) error {
	err := bs.repo.Add(bus)
	return err
}

func (bs *BusService) GetAll() []models.Bus {
	return bs.repo.GetAll()
}

func (bs *BusService) DeleteById(id string) error {
	err := bs.repo.DeleteById(id)
	return err
}
