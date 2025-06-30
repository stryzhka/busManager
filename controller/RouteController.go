package controller

import (
	"busManager/models"
	"busManager/responses"
	"busManager/service"
	"encoding/json"
	"errors"
	"strings"
)

type RouteController struct {
	rs service.IRouteService
}

func NewRouteController(rs service.IRouteService) *RouteController {
	return &RouteController{rs}
}

func (rc RouteController) GetById(id string) string {
	if strings.TrimSpace(id) == "" {
		return responses.NewJsonError(errors.New("ID cant be null"))
	}
	data, err := rc.rs.GetById(id)
	if err != nil {
		return responses.NewJsonError(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return responses.NewJsonError(err)
	}
	return string(jsonData)
}

func (rc RouteController) GetByNumber(number string) string {
	if strings.TrimSpace(number) == "" {
		return responses.NewJsonError(errors.New("Number cant be null"))
	}
	data, err := rc.rs.GetByNumber(number)
	if err != nil {
		return responses.NewJsonError(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return responses.NewJsonError(err)
	}
	return string(jsonData)
}

func (rc RouteController) GetAll() string {
	data, err := rc.rs.GetAll()
	if err != nil {
		return responses.NewJsonError(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return responses.NewJsonError(err)
	}
	return string(jsonData)
}

func (rc RouteController) Add(routeData string) string {
	byteRoute := []byte(routeData)
	var route models.Route
	err := json.Unmarshal(byteRoute, &route)
	if err != nil {
		return responses.NewJsonError(err)
	}
	err = rc.rs.Add(&route)
	if err != nil {
		return responses.NewJsonError(err)
	}
	return routeData
}

func (rc RouteController) DeleteById(id string) string {
	if strings.TrimSpace(id) == "" {
		return responses.NewJsonError(errors.New("ID cant be null"))
	}
	err := rc.rs.DeleteById(id)
	if err != nil {
		return responses.NewJsonError(err)
	}
	return ""
}

func (rc RouteController) UpdateById(routeData string) string {
	byteRoute := []byte(routeData)
	var route models.Route
	err := json.Unmarshal(byteRoute, &route)
	if err != nil {
		return responses.NewJsonError(err)
	}
	err = rc.rs.UpdateById(&route)
	if err != nil {
		return responses.NewJsonError(err)
	}
	return routeData
}

func (rc RouteController) AssignDriver(routeId, driverId string) string {
	if strings.TrimSpace(routeId) == "" {
		return responses.NewJsonError(errors.New("Route ID cant be null"))
	}
	if strings.TrimSpace(driverId) == "" {
		return responses.NewJsonError(errors.New("Driver ID cant be null"))
	}
	err := rc.rs.AssignDriver(routeId, driverId)
	if err != nil {
		return responses.NewJsonError(err)
	}
	return responses.NewSuccessResponse(`Assigned driver successfully`)
}

func (rc RouteController) AssignBusStop(routeId, busStopId string) string {
	if strings.TrimSpace(routeId) == "" {
		return responses.NewJsonError(errors.New("Route ID cant be null"))
	}
	if strings.TrimSpace(busStopId) == "" {
		return responses.NewJsonError(errors.New("Bus stop ID cant be null"))
	}
	err := rc.rs.AssignBusStop(routeId, busStopId)
	if err != nil {
		return responses.NewJsonError(err)
	}
	return responses.NewSuccessResponse(`Assigned bus stop successfully`)
}

func (rc RouteController) AssignBus(routeId, busId string) string {
	if strings.TrimSpace(routeId) == "" {
		return responses.NewJsonError(errors.New("Route ID cant be null"))
	}
	if strings.TrimSpace(busId) == "" {
		return responses.NewJsonError(errors.New("Bus ID cant be null"))
	}
	err := rc.rs.AssignBus(routeId, busId)
	if err != nil {
		return responses.NewJsonError(err)
	}
	return responses.NewSuccessResponse(`Assigned bus successfully`)
}

func (rc RouteController) UnassignDriver(routeId, driverId string) string {
	if strings.TrimSpace(routeId) == "" {
		return responses.NewJsonError(errors.New("Route ID cant be null"))
	}
	if strings.TrimSpace(driverId) == "" {
		return responses.NewJsonError(errors.New("Driver ID cant be null"))
	}
	err := rc.rs.UnassignDriver(routeId, driverId)
	if err != nil {
		return responses.NewJsonError(err)
	}
	return responses.NewSuccessResponse(`Unassigned driver successfully`)
}

func (rc RouteController) UnassignBusStop(routeId, busStopId string) string {
	if strings.TrimSpace(routeId) == "" {
		return responses.NewJsonError(errors.New("Route ID cant be null"))
	}
	if strings.TrimSpace(busStopId) == "" {
		return responses.NewJsonError(errors.New("Bus stop ID cant be null"))
	}
	err := rc.rs.UnassignBusStop(routeId, busStopId)
	if err != nil {
		return responses.NewJsonError(err)
	}
	return responses.NewSuccessResponse(`Unassigned bus stop successfully`)
}

func (rc RouteController) UnassignBus(routeId, busId string) string {
	if strings.TrimSpace(routeId) == "" {
		return responses.NewJsonError(errors.New("Route ID cant be null"))
	}
	if strings.TrimSpace(busId) == "" {
		return responses.NewJsonError(errors.New("Bus ID cant be null"))
	}
	err := rc.rs.UnassignBus(routeId, busId)
	if err != nil {
		return responses.NewJsonError(err)
	}
	return responses.NewSuccessResponse(`Unassigned bus successfully`)
}

func (rc RouteController) GetAllDriversById(routeId string) string {
	if strings.TrimSpace(routeId) == "" {
		return responses.NewJsonError(errors.New("Route ID cant be null"))
	}
	data, err := rc.rs.GetAllDriversById(routeId)
	if err != nil {
		return responses.NewJsonError(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return responses.NewJsonError(err)
	}
	return string(jsonData)
}

func (rc RouteController) GetAllBusesById(routeId string) string {
	if strings.TrimSpace(routeId) == "" {
		return responses.NewJsonError(errors.New("Route ID cant be null"))
	}
	data, err := rc.rs.GetAllBusesById(routeId)
	if err != nil {
		return responses.NewJsonError(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return responses.NewJsonError(err)
	}
	return string(jsonData)
}

func (rc RouteController) GetAllBusStopsById(routeId string) string {
	if strings.TrimSpace(routeId) == "" {
		return responses.NewJsonError(errors.New("Route ID cant be null"))
	}
	data, err := rc.rs.GetAllBusStopsById(routeId)
	if err != nil {
		return responses.NewJsonError(err)
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return responses.NewJsonError(err)
	}
	return string(jsonData)
}
