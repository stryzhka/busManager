package controller

import (
	"busManager/customErrors"
	"busManager/models"
	"busManager/service"
	"encoding/json"
	"errors"
	"strings"
)

type DriverController struct {
	ds service.IDriverService
}

func NewDriverController(ds service.DriverService) *DriverController {
	return &DriverController{ds}
}

func (dc DriverController) GetById(id string) string {
	if strings.TrimSpace(id) == "" {
		return customErrors.NewJsonError(errors.New("ID cant be null"))
	}
	data, err := dc.ds.GetById(id)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	return string(jsonData)
}

func (dc DriverController) GetByPassportSeries(series string) string {
	if strings.TrimSpace(series) == "" {
		return customErrors.NewJsonError(errors.New("PassportSeries cant be null"))
	}
	data, err := dc.ds.GetByPassportSeries(series)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	return string(jsonData)
}

func (dc DriverController) GetAll() string {
	data := dc.ds.GetAll()
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	return string(jsonData)
}

func (dc DriverController) Add(driverData string) string {
	byteDriver := []byte(driverData)
	var driver models.Driver
	err := json.Unmarshal(byteDriver, &driver)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	err = dc.ds.Add(&driver)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	return driverData
}

func (dc DriverController) DeleteById(id string) string {
	if strings.TrimSpace(id) == "" {
		return customErrors.NewJsonError(errors.New("ID cant be null"))
	}
	err := dc.ds.DeleteById(id)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	return ""
}

func (dc DriverController) UpdateById(driverData string) string {
	byteDriver := []byte(driverData)
	var driver models.Driver
	err := json.Unmarshal(byteDriver, &driver)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	err = dc.ds.UpdateById(&driver)
	if err != nil {
		return customErrors.NewJsonError(err)
	}
	return driverData
}
