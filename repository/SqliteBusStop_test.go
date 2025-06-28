package repository

import (
	"busManager/models"
	"database/sql"
	"github.com/google/uuid"
	"testing"
)

func setupTestDBBusStop(t *testing.T) (*SqliteBusStopRepository, func()) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE bus_stops (
			id TEXT PRIMARY KEY,
			route_id TEXT NOT NULL,
			lat REAL NOT NULL,
			long REAL NOT NULL,
			"order" INTEGER NOT NULL,
			name TEXT NOT NULL
		)
	`)
	if err != nil {
		t.Fatalf("Failed to create bus_stops table: %v", err)
	}

	repo := &SqliteBusStopRepository{db: db}
	return repo, func() { db.Close() }
}

func TestSqliteBusStopRepository_GetById(t *testing.T) {
	repo, cleanup := setupTestDBBusStop(t)
	defer cleanup()

	busStop := &models.BusStop{
		ID:      uuid.New().String(),
		RouteId: "route1",
		Lat:     55.7558,
		Long:    37.6173,
		Order:   1,
		Name:    "Stop A",
	}

	_, err := repo.db.Exec(`
		INSERT INTO bus_stops (id, route_id, lat, long, "order", name)
		VALUES (?, ?, ?, ?, ?, ?)`,
		busStop.ID, busStop.RouteId, busStop.Lat, busStop.Long, busStop.Order, busStop.Name)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get existing bus stop by ID", func(t *testing.T) {
		result, err := repo.GetById(busStop.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.ID != busStop.ID {
			t.Errorf("Expected bus stop with ID %s, got %v", busStop.ID, result)
		}
		if result != nil && result.Lat != busStop.Lat {
			t.Errorf("Expected Lat %v, got %v", busStop.Lat, result.Lat)
		}
	})

	t.Run("Get non-existent bus stop by ID", func(t *testing.T) {
		_, err := repo.GetById(uuid.New().String())
		if err == nil || err.Error() != "Bus stop not found" {
			t.Errorf("Expected 'Bus stop not found' error, got %v", err)
		}
	})
}

func TestSqliteBusStopRepository_GetByName(t *testing.T) {
	repo, cleanup := setupTestDBBusStop(t)
	defer cleanup()

	busStop := &models.BusStop{
		ID:      uuid.New().String(),
		RouteId: "route1",
		Lat:     55.7558,
		Long:    37.6173,
		Order:   1,
		Name:    "Stop A",
	}

	_, err := repo.db.Exec(`
		INSERT INTO bus_stops (id, route_id, lat, long, "order", name)
		VALUES (?, ?, ?, ?, ?, ?)`,
		busStop.ID, busStop.RouteId, busStop.Lat, busStop.Long, busStop.Order, busStop.Name)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get existing bus stop by name", func(t *testing.T) {
		result, err := repo.GetByName(busStop.Name)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.Name != busStop.Name {
			t.Errorf("Expected bus stop with name %s, got %v", busStop.Name, result)
		}
		if result != nil && result.Lat != busStop.Lat {
			t.Errorf("Expected Lat %v, got %v", busStop.Lat, result.Lat)
		}
	})

	t.Run("Get non-existent bus stop by name", func(t *testing.T) {
		_, err := repo.GetByName("Unknown Stop")
		if err == nil || err.Error() != "Bus stop not found" {
			t.Errorf("Expected 'Bus stop not found' error, got %v", err)
		}
	})
}

func TestSqliteBusStopRepository_Add(t *testing.T) {
	repo, cleanup := setupTestDBBusStop(t)
	defer cleanup()

	busStop := &models.BusStop{
		RouteId: "route1",
		Lat:     55.7558,
		Long:    37.6173,
		Order:   1,
		Name:    "Stop A",
	}

	t.Run("Add new bus stop", func(t *testing.T) {
		err := repo.Add(busStop)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if busStop.ID == "" {
			t.Errorf("Expected bus stop ID to be set")
		}

		result, err := repo.GetByName(busStop.Name)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.Name != busStop.Name {
			t.Errorf("Expected bus stop with name %s, got %v", busStop.Name, result)
		}
		if result != nil && result.Lat != busStop.Lat {
			t.Errorf("Expected Lat %v, got %v", busStop.Lat, result.Lat)
		}
	})

	t.Run("Add duplicate bus stop", func(t *testing.T) {
		err := repo.Add(busStop)
		if err == nil || err.Error() != "Bus stop already exists" {
			t.Errorf("Expected 'Bus stop already exists' error, got %v", err)
		}
	})
}

func TestSqliteBusStopRepository_GetAll(t *testing.T) {
	repo, cleanup := setupTestDBBusStop(t)
	defer cleanup()

	busStop1 := &models.BusStop{
		ID:      uuid.New().String(),
		RouteId: "route1",
		Lat:     55.7558,
		Long:    37.6173,
		Order:   1,
		Name:    "Stop A",
	}
	busStop2 := &models.BusStop{
		ID:      uuid.New().String(),
		RouteId: "route1",
		Lat:     55.7522,
		Long:    37.6156,
		Order:   2,
		Name:    "Stop B",
	}

	_, err := repo.db.Exec(`
		INSERT INTO bus_stops (id, route_id, lat, long, "order", name)
		VALUES (?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?)`,
		busStop1.ID, busStop1.RouteId, busStop1.Lat, busStop1.Long, busStop1.Order, busStop1.Name,
		busStop2.ID, busStop2.RouteId, busStop2.Lat, busStop2.Long, busStop2.Order, busStop2.Name)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get all bus stops", func(t *testing.T) {
		busStops, err := repo.GetAll()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(busStops) != 2 {
			t.Errorf("Expected 2 bus stops, got %d", len(busStops))
		}
	})

	t.Run("Get all from empty DB", func(t *testing.T) {
		_, err := repo.db.Exec("DELETE FROM bus_stops")
		if err != nil {
			t.Fatalf("Failed to clear database: %v", err)
		}

		busStops, err := repo.GetAll()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(busStops) != 0 {
			t.Errorf("Expected 0 bus stops, got %d", len(busStops))
		}
	})
}

func TestSqliteBusStopRepository_GetAllByRouteId(t *testing.T) {
	repo, cleanup := setupTestDBBusStop(t)
	defer cleanup()

	busStop1 := &models.BusStop{
		ID:      uuid.New().String(),
		RouteId: "route1",
		Lat:     55.7558,
		Long:    37.6173,
		Order:   1,
		Name:    "Stop A",
	}
	busStop2 := &models.BusStop{
		ID:      uuid.New().String(),
		RouteId: "route1",
		Lat:     55.7522,
		Long:    37.6156,
		Order:   2,
		Name:    "Stop B",
	}

	_, err := repo.db.Exec(`
		INSERT INTO bus_stops (id, route_id, lat, long, "order", name)
		VALUES (?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?)`,
		busStop1.ID, busStop1.RouteId, busStop1.Lat, busStop1.Long, busStop1.Order, busStop1.Name,
		busStop2.ID, busStop2.RouteId, busStop2.Lat, busStop2.Long, busStop2.Order, busStop2.Name)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get all bus stops by route ID", func(t *testing.T) {
		busStops, err := repo.GetAllByRouteId("route1")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(busStops) != 2 {
			t.Errorf("Expected 2 bus stops, got %d", len(busStops))
		}
	})

	t.Run("Get all by non-existent route ID", func(t *testing.T) {
		busStops, err := repo.GetAllByRouteId("route2")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(busStops) != 0 {
			t.Errorf("Expected 0 bus stops, got %d", len(busStops))
		}
	})
}

func TestSqliteBusStopRepository_DeleteById(t *testing.T) {
	repo, cleanup := setupTestDBBusStop(t)
	defer cleanup()

	busStop := &models.BusStop{
		ID:      uuid.New().String(),
		RouteId: "route1",
		Lat:     55.7558,
		Long:    37.6173,
		Order:   1,
		Name:    "Stop A",
	}

	_, err := repo.db.Exec(`
		INSERT INTO bus_stops (id, route_id, lat, long, "order", name)
		VALUES (?, ?, ?, ?, ?, ?)`,
		busStop.ID, busStop.RouteId, busStop.Lat, busStop.Long, busStop.Order, busStop.Name)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Delete existing bus stop", func(t *testing.T) {
		err := repo.DeleteById(busStop.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		_, err = repo.GetById(busStop.ID)
		if err == nil || err.Error() != "Bus stop not found" {
			t.Errorf("Expected 'Bus stop not found' error, got %v", err)
		}
	})

	t.Run("Delete non-existent bus stop", func(t *testing.T) {
		err := repo.DeleteById(uuid.New().String())
		if err == nil || err.Error() != "Bus stop not found" {
			t.Errorf("Expected 'Bus stop not found' error, got %v", err)
		}
	})
}

func TestSqliteBusStopRepository_UpdateById(t *testing.T) {
	repo, cleanup := setupTestDBBusStop(t)
	defer cleanup()

	busStop := &models.BusStop{
		ID:      uuid.New().String(),
		RouteId: "route1",
		Lat:     55.7558,
		Long:    37.6173,
		Order:   1,
		Name:    "Stop A",
	}

	_, err := repo.db.Exec(`
		INSERT INTO bus_stops (id, route_id, lat, long, "order", name)
		VALUES (?, ?, ?, ?, ?, ?)`,
		busStop.ID, busStop.RouteId, busStop.Lat, busStop.Long, busStop.Order, busStop.Name)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Update existing bus stop", func(t *testing.T) {
		updatedBusStop := &models.BusStop{
			ID:      busStop.ID,
			RouteId: "route2",
			Lat:     55.7522,
			Long:    37.6156,
			Order:   2,
			Name:    "Stop B",
		}

		err := repo.UpdateById(updatedBusStop)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		result, err := repo.GetById(busStop.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil {
			t.Errorf("Expected bus stop, got nil")
		} else if result.Name != updatedBusStop.Name || result.Lat != updatedBusStop.Lat {
			t.Errorf("Expected updated bus stop data, got %v", result)
		}
	})

	t.Run("Update non-existent bus stop", func(t *testing.T) {
		nonExistentBusStop := &models.BusStop{
			ID:      uuid.New().String(),
			RouteId: "route2",
			Lat:     55.7522,
			Long:    37.6156,
			Order:   2,
			Name:    "Stop B",
		}

		err := repo.UpdateById(nonExistentBusStop)
		if err == nil || err.Error() != "Bus stop not found" {
			t.Errorf("Expected 'Bus stop not found' error, got %v", err)
		}
	})
}
