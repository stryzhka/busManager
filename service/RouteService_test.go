package service

import (
	"busManager/models"
	"errors"
	"github.com/google/uuid"
	"testing"
	"time"
)

type MockRouteRepository struct {
	getByIdResp      *models.Route
	getByIdErr       error
	getByNumberResp  *models.Route
	getByNumberErr   error
	addErr           error
	getAllResp       []models.Route
	deleteByIdErr    error
	updateByIdErr    error
	assignDriverErr  error
	assignBusStopErr error
	assignBusErr     error
}

func (m *MockRouteRepository) GetById(id string) (*models.Route, error) {
	return m.getByIdResp, m.getByIdErr
}

func (m *MockRouteRepository) GetByNumber(number string) (*models.Route, error) {
	return m.getByNumberResp, m.getByNumberErr
}

func (m *MockRouteRepository) Add(route *models.Route) error {
	return m.addErr
}

func (m *MockRouteRepository) GetAll() ([]models.Route, error) {
	return m.getAllResp, nil
}

func (m *MockRouteRepository) DeleteById(id string) error {
	return m.deleteByIdErr
}

func (m *MockRouteRepository) UpdateById(route *models.Route) error {
	return m.updateByIdErr
}

func (m *MockRouteRepository) AssignDriver(routeId, driverId string) error {
	return m.assignDriverErr
}

func (m *MockRouteRepository) AssignBusStop(routeId, busStopId string) error {
	return m.assignBusStopErr
}

func (m *MockRouteRepository) AssignBus(routeId, busId string) error {
	return m.assignBusErr
}

type MockBusRepository struct {
	getByIdResp     *models.Bus
	getByIdErr      error
	getByNumberResp *models.Bus
	getByNumberErr  error
	addErr          error
	getAllResp      []models.Bus
	deleteByIdErr   error
	updateByIdErr   error
}

func (m *MockBusRepository) GetById(id string) (*models.Bus, error) {
	return m.getByIdResp, m.getByIdErr
}

func (m *MockBusRepository) GetByNumber(number string) (*models.Bus, error) {
	return m.getByNumberResp, m.getByNumberErr
}

func (m *MockBusRepository) Add(bus *models.Bus) error {
	return m.addErr
}

func (m *MockBusRepository) DeleteById(id string) error {
	return m.deleteByIdErr
}

func (m *MockBusRepository) GetAll() ([]models.Bus, error) {
	return m.getAllResp, nil
}

func (m *MockBusRepository) UpdateById(bus *models.Bus) error {
	return m.updateByIdErr
}

func TestRouteService_GetById(t *testing.T) {
	route := &models.Route{ID: uuid.New().String(), Number: "101"}

	t.Run("Success", func(t *testing.T) {
		mockRepo := &MockRouteRepository{getByIdResp: route}
		service := NewRouteService(mockRepo, nil, nil, nil)

		result, err := service.GetById(route.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.ID != route.ID {
			t.Errorf("Expected route with ID %s, got %v", route.ID, result)
		}
	})

	t.Run("Not found", func(t *testing.T) {
		mockRepo := &MockRouteRepository{getByIdErr: errors.New("Route not found")}
		service := NewRouteService(mockRepo, nil, nil, nil)

		_, err := service.GetById(uuid.New().String())
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})
}

func TestRouteService_GetByNumber(t *testing.T) {
	route := &models.Route{ID: uuid.New().String(), Number: "101"}

	t.Run("Success", func(t *testing.T) {
		mockRepo := &MockRouteRepository{getByNumberResp: route}
		service := NewRouteService(mockRepo, nil, nil, nil)

		result, err := service.GetByNumber("101")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.Number != route.Number {
			t.Errorf("Expected route with number %s, got %v", route.Number, result)
		}
	})

	t.Run("Not found", func(t *testing.T) {
		mockRepo := &MockRouteRepository{getByNumberErr: errors.New("Route not found")}
		service := NewRouteService(mockRepo, nil, nil, nil)

		_, err := service.GetByNumber("999")
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})
}

func TestRouteService_Add(t *testing.T) {
	route := &models.Route{Number: "101"}

	t.Run("Success", func(t *testing.T) {
		mockRepo := &MockRouteRepository{}
		service := NewRouteService(mockRepo, nil, nil, nil)

		err := service.Add(route)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Add with error from repo", func(t *testing.T) {
		mockRepo := &MockRouteRepository{addErr: errors.New("Database error")}
		service := NewRouteService(mockRepo, nil, nil, nil)

		err := service.Add(route)
		if err == nil || err.Error() != "Database error" {
			t.Errorf("Expected 'Database error', got %v", err)
		}
	})
}

func TestRouteService_GetAll(t *testing.T) {
	route1 := models.Route{ID: uuid.New().String(), Number: "101"}
	route2 := models.Route{ID: uuid.New().String(), Number: "102"}

	t.Run("Get all routes", func(t *testing.T) {
		mockRepo := &MockRouteRepository{getAllResp: []models.Route{route1, route2}}
		service := NewRouteService(mockRepo, nil, nil, nil)

		routes := service.GetAll()
		if len(routes) != 2 {
			t.Errorf("Expected 2 routes, got %d", len(routes))
		}
	})

	t.Run("Get all from empty repo", func(t *testing.T) {
		mockRepo := &MockRouteRepository{getAllResp: []models.Route{}}
		service := NewRouteService(mockRepo, nil, nil, nil)

		routes := service.GetAll()
		if len(routes) != 0 {
			t.Errorf("Expected 0 routes, got %d", len(routes))
		}
	})
}

func TestRouteService_DeleteById(t *testing.T) {
	t.Run("Delete existing route", func(t *testing.T) {
		mockRepo := &MockRouteRepository{}
		service := NewRouteService(mockRepo, nil, nil, nil)

		err := service.DeleteById(uuid.New().String())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Delete with error from repo", func(t *testing.T) {
		mockRepo := &MockRouteRepository{deleteByIdErr: errors.New("Database error")}
		service := NewRouteService(mockRepo, nil, nil, nil)

		err := service.DeleteById(uuid.New().String())
		if err == nil || err.Error() != "Database error" {
			t.Errorf("Expected 'Database error', got %v", err)
		}
	})
}

func TestRouteService_UpdateById(t *testing.T) {
	route := &models.Route{ID: uuid.New().String(), Number: "101"}

	t.Run("Update existing route", func(t *testing.T) {
		mockRepo := &MockRouteRepository{}
		service := NewRouteService(mockRepo, nil, nil, nil)

		err := service.UpdateById(route)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Update with error from repo", func(t *testing.T) {
		mockRepo := &MockRouteRepository{updateByIdErr: errors.New("Database error")}
		service := NewRouteService(mockRepo, nil, nil, nil)

		err := service.UpdateById(route)
		if err == nil || err.Error() != "Database error" {
			t.Errorf("Expected 'Database error', got %v", err)
		}
	})
}

func TestRouteService_AssignDriver(t *testing.T) {
	route := &models.Route{ID: uuid.New().String(), Number: "101"}
	driver := &models.Driver{ID: uuid.New().String(), Name: "John"}

	t.Run("Success", func(t *testing.T) {
		mockRouteRepo := &MockRouteRepository{getByIdResp: route}
		mockDriverRepo := &MockDriverRepository{getByIdResp: driver}
		service := NewRouteService(mockRouteRepo, mockDriverRepo, nil, nil)

		err := service.AssignDriver(route.ID, driver.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Route not found", func(t *testing.T) {
		mockRouteRepo := &MockRouteRepository{getByIdErr: errors.New("Route not found")}
		mockDriverRepo := &MockDriverRepository{getByIdResp: driver}
		service := NewRouteService(mockRouteRepo, mockDriverRepo, nil, nil)

		err := service.AssignDriver(uuid.New().String(), driver.ID)
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})

	t.Run("Driver not found", func(t *testing.T) {
		mockRouteRepo := &MockRouteRepository{getByIdResp: route}
		mockDriverRepo := &MockDriverRepository{getByIdErr: errors.New("Driver not found")}
		service := NewRouteService(mockRouteRepo, mockDriverRepo, nil, nil)

		err := service.AssignDriver(route.ID, uuid.New().String())
		if err == nil || err.Error() != "Driver not found" {
			t.Errorf("Expected 'Driver not found' error, got %v", err)
		}
	})

	t.Run("Assign driver with repo error", func(t *testing.T) {
		mockRouteRepo := &MockRouteRepository{getByIdResp: route, assignDriverErr: errors.New("Database error")}
		mockDriverRepo := &MockDriverRepository{getByIdResp: driver}
		service := NewRouteService(mockRouteRepo, mockDriverRepo, nil, nil)

		err := service.AssignDriver(route.ID, driver.ID)
		if err == nil || err.Error() != "Database error" {
			t.Errorf("Expected 'Database error', got %v", err)
		}
	})
}

func TestRouteService_AssignBusStop(t *testing.T) {
	route := &models.Route{ID: uuid.New().String(), Number: "101"}
	busStop := &models.BusStop{ID: uuid.New().String(), Name: "Stop A"}

	t.Run("Success", func(t *testing.T) {
		mockRouteRepo := &MockRouteRepository{getByIdResp: route}
		mockBusStopRepo := &MockBusStopRepository{getByIdResp: busStop}
		service := NewRouteService(mockRouteRepo, nil, nil, mockBusStopRepo)

		err := service.AssignBusStop(route.ID, busStop.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Route not found", func(t *testing.T) {
		mockRouteRepo := &MockRouteRepository{getByIdErr: errors.New("Route not found")}
		mockBusStopRepo := &MockBusStopRepository{getByIdResp: busStop}
		service := NewRouteService(mockRouteRepo, nil, nil, mockBusStopRepo)

		err := service.AssignBusStop(uuid.New().String(), busStop.ID)
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})

	t.Run("Bus stop not found", func(t *testing.T) {
		mockRouteRepo := &MockRouteRepository{getByIdResp: route}
		mockBusStopRepo := &MockBusStopRepository{getByIdErr: errors.New("Bus stop not found")}
		service := NewRouteService(mockRouteRepo, nil, nil, mockBusStopRepo)

		err := service.AssignBusStop(route.ID, uuid.New().String())
		if err == nil || err.Error() != "Bus stop not found" {
			t.Errorf("Expected 'Bus stop not found' error, got %v", err)
		}
	})

	t.Run("Assign bus stop with repo error", func(t *testing.T) {
		mockRouteRepo := &MockRouteRepository{getByIdResp: route, assignBusStopErr: errors.New("Database error")}
		mockBusStopRepo := &MockBusStopRepository{getByIdResp: busStop}
		service := NewRouteService(mockRouteRepo, nil, nil, mockBusStopRepo)

		err := service.AssignBusStop(route.ID, busStop.ID)
		if err == nil || err.Error() != "Database error" {
			t.Errorf("Expected 'Database error', got %v", err)
		}
	})
}

func TestRouteService_AssignBus(t *testing.T) {
	route := &models.Route{ID: uuid.New().String(), Number: "101"}
	bus := &models.Bus{
		ID:             "2",
		RegisterNumber: "TEST666",
		Brand:          "Sca2nia MANDEC",
		BusModel:       "66622",
		AssemblyDate:   time.Now(),
		LastRepairDate: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		mockRouteRepo := &MockRouteRepository{getByIdResp: route}
		mockBusRepo := &MockBusRepository{getByIdResp: bus}
		service := NewRouteService(mockRouteRepo, nil, mockBusRepo, nil)

		err := service.AssignBus(route.ID, bus.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("Route not found", func(t *testing.T) {
		mockRouteRepo := &MockRouteRepository{getByIdErr: errors.New("Route not found")}
		mockBusRepo := &MockBusRepository{getByIdResp: bus}
		service := NewRouteService(mockRouteRepo, nil, mockBusRepo, nil)

		err := service.AssignBus(uuid.New().String(), bus.ID)
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})

	t.Run("Bus not found", func(t *testing.T) {
		mockRouteRepo := &MockRouteRepository{getByIdResp: route}
		mockBusRepo := &MockBusRepository{getByIdErr: errors.New("Bus not found")}
		service := NewRouteService(mockRouteRepo, nil, mockBusRepo, nil)

		err := service.AssignBus(route.ID, uuid.New().String())
		if err == nil || err.Error() != "Bus not found" {
			t.Errorf("Expected 'Bus not found' error, got %v", err)
		}
	})

	t.Run("Assign bus with repo error", func(t *testing.T) {
		mockRouteRepo := &MockRouteRepository{getByIdResp: route, assignBusErr: errors.New("Database error")}
		mockBusRepo := &MockBusRepository{getByIdResp: bus}
		service := NewRouteService(mockRouteRepo, nil, mockBusRepo, nil)

		err := service.AssignBus(route.ID, bus.ID)
		if err == nil || err.Error() != "Database error" {
			t.Errorf("Expected 'Database error', got %v", err)
		}
	})
}
