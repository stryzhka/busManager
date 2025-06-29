package repository

import (
	"busManager/models"
	"database/sql"
	"github.com/google/uuid"
	"testing"
)

func setupTestDBRoute(t *testing.T) (*SqliteRouteRepository, func()) {
	db, err := sql.Open("sqlite3", ":memory:?parseTime=true")
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE routes (
			id TEXT PRIMARY KEY,
			number TEXT NOT NULL
		)
	`)
	if err != nil {
		t.Fatalf("Failed to create routes table: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE routes_drivers (
			route_id TEXT NOT NULL,
			driver_id TEXT NOT NULL,
			PRIMARY KEY (route_id, driver_id),
			FOREIGN KEY (route_id) REFERENCES routes(id)
		)
	`)
	if err != nil {
		t.Fatalf("Failed to create routes_drivers table: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE routes_bus_stops (
			route_id TEXT NOT NULL,
			bus_stop_id TEXT NOT NULL,
			PRIMARY KEY (route_id, bus_stop_id),
			FOREIGN KEY (route_id) REFERENCES routes(id)
		)
	`)
	if err != nil {
		t.Fatalf("Failed to create routes_bus_stops table: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE routes_buses (
			route_id TEXT NOT NULL,
			bus_id TEXT NOT NULL,
			PRIMARY KEY (route_id, bus_id),
			FOREIGN KEY (route_id) REFERENCES routes(id)
		)
	`)
	if err != nil {
		t.Fatalf("Failed to create routes_buses table: %v", err)
	}

	repo := &SqliteRouteRepository{db: db}
	return repo, func() { db.Close() }
}

func TestSqliteRouteRepository_GetById(t *testing.T) {
	repo, cleanup := setupTestDBRoute(t)
	defer cleanup()

	route := &models.Route{
		ID:     uuid.New().String(),
		Number: "101",
	}

	_, err := repo.db.Exec(`
		INSERT INTO routes (id, number)
		VALUES (?, ?)`,
		route.ID, route.Number)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get existing route by ID", func(t *testing.T) {
		result, err := repo.GetById(route.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.ID != route.ID {
			t.Errorf("Expected route with ID %s, got %v", route.ID, result)
		}
		if result != nil && result.Number != route.Number {
			t.Errorf("Expected Number %s, got %s", route.Number, result.Number)
		}
	})

	t.Run("Get non-existent route by ID", func(t *testing.T) {
		_, err := repo.GetById(uuid.New().String())
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})
}

func TestSqliteRouteRepository_GetByNumber(t *testing.T) {
	repo, cleanup := setupTestDBRoute(t)
	defer cleanup()

	route := &models.Route{
		ID:     uuid.New().String(),
		Number: "101",
	}

	_, err := repo.db.Exec(`
		INSERT INTO routes (id, number)
		VALUES (?, ?)`,
		route.ID, route.Number)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get existing route by number", func(t *testing.T) {
		result, err := repo.GetByNumber(route.Number)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.Number != route.Number {
			t.Errorf("Expected route with number %s, got %v", route.Number, result)
		}
		if result != nil && result.ID != route.ID {
			t.Errorf("Expected ID %s, got %s", route.ID, result.ID)
		}
	})

	t.Run("Get non-existent route by number", func(t *testing.T) {
		_, err := repo.GetByNumber("999")
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})
}

func TestSqliteRouteRepository_Add(t *testing.T) {
	repo, cleanup := setupTestDBRoute(t)
	defer cleanup()

	route := &models.Route{
		Number: "101",
	}

	t.Run("Add new route", func(t *testing.T) {
		err := repo.Add(route)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if route.ID == "" {
			t.Errorf("Expected route ID to be set")
		}

		result, err := repo.GetByNumber(route.Number)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.Number != route.Number {
			t.Errorf("Expected route with number %s, got %v", route.Number, result)
		}
	})

	t.Run("Add duplicate route", func(t *testing.T) {
		err := repo.Add(route)
		if err == nil || err.Error() != "Route already exists" {
			t.Errorf("Expected 'Route already exists' error, got %v", err)
		}
	})
}

func TestSqliteRouteRepository_GetAll(t *testing.T) {
	repo, cleanup := setupTestDBRoute(t)
	defer cleanup()

	route1 := &models.Route{
		ID:     uuid.New().String(),
		Number: "101",
	}
	route2 := &models.Route{
		ID:     uuid.New().String(),
		Number: "102",
	}

	_, err := repo.db.Exec(`
		INSERT INTO routes (id, number)
		VALUES (?, ?), (?, ?)`,
		route1.ID, route1.Number, route2.ID, route2.Number)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get all routes", func(t *testing.T) {
		routes, err := repo.GetAll()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(routes) != 2 {
			t.Errorf("Expected 2 routes, got %d", len(routes))
		}
	})

	t.Run("Get all from empty DB", func(t *testing.T) {
		_, err := repo.db.Exec("DELETE FROM routes")
		if err != nil {
			t.Fatalf("Failed to clear database: %v", err)
		}

		routes, err := repo.GetAll()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(routes) != 0 {
			t.Errorf("Expected 0 routes, got %d", len(routes))
		}
	})
}

func TestSqliteRouteRepository_DeleteById(t *testing.T) {
	repo, cleanup := setupTestDBRoute(t)
	defer cleanup()

	route := &models.Route{
		ID:     uuid.New().String(),
		Number: "101",
	}

	_, err := repo.db.Exec(`
		INSERT INTO routes (id, number)
		VALUES (?, ?)`,
		route.ID, route.Number)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Delete existing route", func(t *testing.T) {
		err := repo.DeleteById(route.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		_, err = repo.GetById(route.ID)
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})

	t.Run("Delete non-existent route", func(t *testing.T) {
		err := repo.DeleteById(uuid.New().String())
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})
}

func TestSqliteRouteRepository_UpdateById(t *testing.T) {
	repo, cleanup := setupTestDBRoute(t)
	defer cleanup()

	route := &models.Route{
		ID:     uuid.New().String(),
		Number: "101",
	}

	_, err := repo.db.Exec(`
		INSERT INTO routes (id, number)
		VALUES (?, ?)`,
		route.ID, route.Number)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Update existing route", func(t *testing.T) {
		updatedRoute := &models.Route{
			ID:     route.ID,
			Number: "102",
		}

		err := repo.UpdateById(updatedRoute)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		result, err := repo.GetById(route.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil {
			t.Errorf("Expected route, got nil")
		} else if result.Number != updatedRoute.Number {
			t.Errorf("Expected updated number %s, got %s", updatedRoute.Number, result.Number)
		}
	})

	t.Run("Update non-existent route", func(t *testing.T) {
		nonExistentRoute := &models.Route{
			ID:     uuid.New().String(),
			Number: "103",
		}

		err := repo.UpdateById(nonExistentRoute)
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})
}

func TestSqliteRouteRepository_AssignDriver(t *testing.T) {
	repo, cleanup := setupTestDBRoute(t)
	defer cleanup()

	routeID := uuid.New().String()
	driverID := uuid.New().String()

	_, err := repo.db.Exec(`
		INSERT INTO routes (id, number)
		VALUES (?, ?)`,
		routeID, "101")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Assign driver to route", func(t *testing.T) {
		err := repo.AssignDriver(routeID, driverID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		var count int
		err = repo.db.QueryRow(`SELECT COUNT(*) FROM routes_drivers WHERE route_id = ? AND driver_id = ?`, routeID, driverID).Scan(&count)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if count != 1 {
			t.Errorf("Expected 1 record in routes_drivers, got %d", count)
		}
	})

	t.Run("Assign driver to non-existent route", func(t *testing.T) {
		err := repo.AssignDriver(uuid.New().String(), driverID)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		// Проверяем, что запись не добавлена
		var count int
		err = repo.db.QueryRow(`SELECT COUNT(*) FROM routes_drivers WHERE route_id = ? AND driver_id = ?`, uuid.New().String(), driverID).Scan(&count)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if count != 0 {
			t.Errorf("Expected 0 records, got %d", count)
		}
	})
}

func TestSqliteRouteRepository_AssignBusStop(t *testing.T) {
	repo, cleanup := setupTestDBRoute(t)
	defer cleanup()

	routeID := uuid.New().String()
	busStopID := uuid.New().String()

	_, err := repo.db.Exec(`
		INSERT INTO routes (id, number)
		VALUES (?, ?)`,
		routeID, "101")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Assign bus stop to route", func(t *testing.T) {
		err := repo.AssignBusStop(routeID, busStopID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		var count int
		err = repo.db.QueryRow(`SELECT COUNT(*) FROM routes_bus_stops WHERE route_id = ? AND bus_stop_id = ?`, routeID, busStopID).Scan(&count)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if count != 1 {
			t.Errorf("Expected 1 record in routes_bus_stops, got %d", count)
		}
	})

	t.Run("Assign bus stop to non-existent route", func(t *testing.T) {
		err := repo.AssignBusStop(uuid.New().String(), busStopID)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		var count int
		err = repo.db.QueryRow(`SELECT COUNT(*) FROM routes_bus_stops WHERE route_id = ? AND bus_stop_id = ?`, uuid.New().String(), busStopID).Scan(&count)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if count != 0 {
			t.Errorf("Expected 0 records, got %d", count)
		}
	})
}

func TestSqliteRouteRepository_AssignBus(t *testing.T) {
	repo, cleanup := setupTestDBRoute(t)
	defer cleanup()

	routeID := uuid.New().String()
	busID := uuid.New().String()

	_, err := repo.db.Exec(`
		INSERT INTO routes (id, number)
		VALUES (?, ?)`,
		routeID, "101")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Assign bus to route", func(t *testing.T) {
		err := repo.AssignBus(routeID, busID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		var count int
		err = repo.db.QueryRow(`SELECT COUNT(*) FROM routes_buses WHERE route_id = ? AND bus_id = ?`, routeID, busID).Scan(&count)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if count != 1 {
			t.Errorf("Expected 1 record in routes_buses, got %d", count)
		}
	})

	t.Run("Assign bus to non-existent route", func(t *testing.T) {
		err := repo.AssignBus(uuid.New().String(), busID)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
		var count int
		err = repo.db.QueryRow(`SELECT COUNT(*) FROM routes_buses WHERE route_id = ? AND bus_id = ?`, uuid.New().String(), busID).Scan(&count)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if count != 0 {
			t.Errorf("Expected 0 records, got %d", count)
		}
	})
}
