package controller

import (
	"busManager/models"
	"busManager/responses"
	"busManager/service"
	"encoding/json"
	"errors"
	"strings"
)

type BusStopController struct {
	bss service.IBusStopService
}

func NewBusStopController(bss service.IBusStopService) *BusStopController {
	return &BusStopController{bss}
}

func (bsc BusStopController) GetById(id string) string {
	if strings.TrimSpace(id) == "" {
		return responses.NewJsonError(errors.New("ID cant be null"))
	}
	data, err := bsc.bss.GetById(id)
	if err != nil {
		return responses.NewJsonError(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return responses.NewJsonError(err)
	}
	return string(jsonData)
}

func (bsc BusStopController) GetByName(name string) string {
	if strings.TrimSpace(name) == "" {
		return responses.NewJsonError(errors.New("Name cant be null"))
	}
	data, err := bsc.bss.GetByName(name)
	if err != nil {
		return responses.NewJsonError(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return responses.NewJsonError(err)
	}
	return string(jsonData)
}

func (bsc BusStopController) GetAll() string {
	data, err := bsc.bss.GetAll()
	if err != nil {
		return responses.NewJsonError(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return responses.NewJsonError(err)
	}
	return string(jsonData)
}

func (bsc BusStopController) Add(busStopData string) string {
	byteBusStop := []byte(busStopData)
	var busStop models.BusStop
	err := json.Unmarshal(byteBusStop, &busStop)
	if err != nil {
		return responses.NewJsonError(err)
	}
	err = bsc.bss.Add(&busStop)
	if err != nil {
		return responses.NewJsonError(err)
	}
	return busStopData
}

func (bsc BusStopController) DeleteById(id string) string {
	if strings.TrimSpace(id) == "" {
		return responses.NewJsonError(errors.New("ID cant be null"))
	}
	err := bsc.bss.DeleteById(id)
	if err != nil {
		return responses.NewJsonError(err)
	}
	return ""
}

func (bsc BusStopController) UpdateById(busStopData string) string {
	byteBusStop := []byte(busStopData)
	var busStop models.BusStop
	err := json.Unmarshal(byteBusStop, &busStop)
	if err != nil {
		return responses.NewJsonError(err)
	}
	err = bsc.bss.UpdateById(&busStop)
	if err != nil {
		return responses.NewJsonError(err)
	}
	return busStopData
}
