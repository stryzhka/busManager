package repository

import (
	"busManager/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type ListBusRepository struct {
	busList []models.Bus
}

func NewListBusRepository() (ListBusRepository, error) {
	newRepo := &ListBusRepository{busList: make([]models.Bus, 0)}
	return *newRepo, nil
}

func (l *ListBusRepository) GetById(id string) (*models.Bus, error) {
	for _, bus := range l.busList {
		if bus.Id == id {
			return &bus, nil
		}
	}
	return nil, nil
}

func (l *ListBusRepository) GetByNumber(number string) (*models.Bus, error) {
	for _, bus := range l.busList {
		if strings.ToUpper(bus.RegisterNumber) == strings.ToUpper(number) {
			return &bus, nil
		}
	}
	return nil, nil
}

func (l *ListBusRepository) Add(bus *models.Bus) error {
	exist, err := l.GetById(bus.Id)
	if err != nil {
		fmt.Println("Occured error, ", err)
		return err
	}
	if exist != nil {
		return errors.New("Bus already exists")
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	bus.Id = id.String()
	l.busList = append(l.busList, *bus)
	return nil
}

func (l *ListBusRepository) GetAll() []models.Bus {
	return l.busList
}

func (l *ListBusRepository) DeleteById(id string) error {
	for i, bus := range l.busList {
		if bus.Id == id {
			l.busList = append(l.busList[:i], l.busList[i+1:]...)
			return nil
		}
	}
	return errors.New("Bus not found")
}

func (l *ListBusRepository) UpdateById(bus *models.Bus) error {

	for i, tBus := range l.busList {
		if tBus.Id == bus.Id {
			l.busList[i] = *bus
			return nil
		}
	}
	return nil
}
