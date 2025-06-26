package controller

import (
	"busManager/customErrors"
	"busManager/models"
	"busManager/service"
	"encoding/json"
	"errors"
	"strings"
)

type BusController struct {
	bs service.IBusService
}

func NewBusController(bs service.BusService) *BusController {
	return &BusController{bs}
}

func (bc BusController) GetBusById(id string) string {
	if strings.TrimSpace(id) == "" {
		return customErrors.NewJsonError(errors.New("Id cant be null"))
	}
	data, err := bc.bs.GetById(id)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	return string(jsonData)
}

func (bc BusController) GetBusByNumber(number string) string {
	if strings.TrimSpace(number) == "" {
		return customErrors.NewJsonError(errors.New("Number cant be null"))
	}
	data, err := bc.bs.GetByNumber(number)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	return string(jsonData)
}

func (bc BusController) GetAllBuses() string {
	data := bc.bs.GetAll()
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	return string(jsonData)
}

func (bc BusController) Add(busData string) string {
	byteBus := []byte(busData)
	var bus models.Bus
	err := json.Unmarshal(byteBus, &bus)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	err = bc.bs.Add(&bus)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	return busData
}

func (bc BusController) DeleteById(id string) string {
	if strings.TrimSpace(id) == "" {
		return customErrors.NewJsonError(errors.New("Id cant be null"))
	}
	err := bc.bs.DeleteById(id)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	return ""
}

func (bc BusController) UpdateById(busData string) string {
	byteBus := []byte(busData)
	var bus models.Bus
	err := json.Unmarshal(byteBus, &bus)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	err = bc.bs.UpdateById(&bus)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	return busData
}
