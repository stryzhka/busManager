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

	// Добавляем таблицы для новых методов
	_, err = db.Exec(`
        CREATE TABLE drivers (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            surname TEXT,
            patronymic TEXT,
            birth_date DATETIME,
            passport_series TEXT,
            snils TEXT,
            license_series TEXT
        )
    `)
	if err != nil {
		t.Fatalf("Failed to create drivers table: %v", err)
	}

	_, err = db.Exec(`
        CREATE TABLE bus_stops (
            id TEXT PRIMARY KEY,
            lat REAL,
            long REAL,
            name TEXT
        )
    `)
	if err != nil {
		t.Fatalf("Failed to create bus_stops table: %v", err)
	}

	_, err = db.Exec(`
        CREATE TABLE buses (
            id TEXT PRIMARY KEY,
            brand TEXT,
            bus_model TEXT,
            register_number TEXT,
            assembly_date DATETIME,
            last_repair_date DATETIME
        )
    `)
	if err != nil {
		t.Fatalf("Failed to create buses table: %v", err)
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

func TestSqliteRouteRepository_UnassignBusStop(t *testing.T) {
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

	// Предварительно добавляем связь
	_, err = repo.db.Exec(`
		INSERT INTO routes_bus_stops (route_id, bus_stop_id)
		VALUES (?, ?)`,
		routeID, busStopID)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Unassign bus stop from route", func(t *testing.T) {
		err := repo.UnassignBusStop(routeID, busStopID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		var count int
		err = repo.db.QueryRow(`SELECT COUNT(*) FROM routes_bus_stops WHERE route_id = ? AND bus_stop_id = ?`, routeID, busStopID).Scan(&count)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if count != 0 {
			t.Errorf("Expected 0 records in routes_bus_stops, got %d", count)
		}
	})

	t.Run("Unassign bus stop from non-existent route", func(t *testing.T) {
		err := repo.UnassignBusStop(uuid.New().String(), busStopID)
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})

	t.Run("Unassign bus stop with no existing relationship", func(t *testing.T) {
		// Проверяем случай, когда связи нет
		err := repo.UnassignBusStop(routeID, uuid.New().String())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		// Ожидаем, что ничего не сломается
	})
}

func TestSqliteRouteRepository_UnassignBus(t *testing.T) {
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

	// Предварительно добавляем связь
	_, err = repo.db.Exec(`
		INSERT INTO routes_buses (route_id, bus_id)
		VALUES (?, ?)`,
		routeID, busID)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Unassign bus from route", func(t *testing.T) {
		err := repo.UnassignBus(routeID, busID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		var count int
		err = repo.db.QueryRow(`SELECT COUNT(*) FROM routes_buses WHERE route_id = ? AND bus_id = ?`, routeID, busID).Scan(&count)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if count != 0 {
			t.Errorf("Expected 0 records in routes_buses, got %d", count)
		}
	})

	t.Run("Unassign bus from non-existent route", func(t *testing.T) {
		err := repo.UnassignBus(uuid.New().String(), busID)
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})

	t.Run("Unassign bus with no existing relationship", func(t *testing.T) {
		err := repo.UnassignBus(routeID, uuid.New().String())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

func TestSqliteRouteRepository_UnassignDriver(t *testing.T) {
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

	// Предварительно добавляем связь
	_, err = repo.db.Exec(`
		INSERT INTO routes_drivers (route_id, driver_id)
		VALUES (?, ?)`,
		routeID, driverID)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Unassign driver from route", func(t *testing.T) {
		err := repo.UnassignDriver(routeID, driverID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		var count int
		err = repo.db.QueryRow(`SELECT COUNT(*) FROM routes_drivers WHERE route_id = ? AND driver_id = ?`, routeID, driverID).Scan(&count)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if count != 0 {
			t.Errorf("Expected 0 records in routes_drivers, got %d", count)
		}
	})

	t.Run("Unassign driver from non-existent route", func(t *testing.T) {
		err := repo.UnassignDriver(uuid.New().String(), driverID)
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})

	t.Run("Unassign driver with no existing relationship", func(t *testing.T) {
		err := repo.UnassignDriver(routeID, uuid.New().String())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

func TestSqliteRouteRepository_GetAllDriversById(t *testing.T) {
	repo, cleanup := setupTestDBRoute(t)
	defer cleanup()

	routeID := uuid.New().String()
	driverID1 := uuid.New().String()
	driverID2 := uuid.New().String()

	// Вставляем маршрут
	_, err := repo.db.Exec(`
        INSERT INTO routes (id, number)
        VALUES (?, ?)`,
		routeID, "101")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Вставляем водителей
	_, err = repo.db.Exec(`
        INSERT INTO drivers (id, name, surname, patronymic, birth_date, passport_series, snils, license_series)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?, ?, ?)`,
		driverID1, "John", "Doe", "Smith", "1990-01-01", "1234", "123-456-789 01", "AB123",
		driverID2, "Jane", "Smith", "Johnson", "1992-02-02", "5678", "987-654-321 02", "CD456")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Вставляем связи
	_, err = repo.db.Exec(`
        INSERT INTO routes_drivers (route_id, driver_id)
        VALUES (?, ?), (?, ?)`,
		routeID, driverID1, routeID, driverID2)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get all drivers by route", func(t *testing.T) {
		drivers, err := repo.GetAllDriversById(routeID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(drivers) != 2 {
			t.Errorf("Expected 2 drivers, got %d", len(drivers))
		}
	})

	t.Run("Get drivers from non-existent route", func(t *testing.T) {
		_, err := repo.GetAllDriversById(uuid.New().String())
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})
}

func TestSqliteRouteRepository_GetAllBusStopsById(t *testing.T) {
	repo, cleanup := setupTestDBRoute(t)
	defer cleanup()

	routeID := uuid.New().String()
	busStopID1 := uuid.New().String()
	busStopID2 := uuid.New().String()

	// Вставляем маршрут
	_, err := repo.db.Exec(`
        INSERT INTO routes (id, number)
        VALUES (?, ?)`,
		routeID, "101")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Вставляем автобусные остановки
	_, err = repo.db.Exec(`
        INSERT INTO bus_stops (id, lat, long, name)
        VALUES (?, ?, ?, ?), (?, ?, ?, ?)`,
		busStopID1, 55.7558, 37.6173, "Red Square",
		busStopID2, 55.7539, 37.6208, "Kremlin")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Вставляем связи
	_, err = repo.db.Exec(`
        INSERT INTO routes_bus_stops (route_id, bus_stop_id)
        VALUES (?, ?), (?, ?)`,
		routeID, busStopID1, routeID, busStopID2)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get all bus stops by route", func(t *testing.T) {
		busStops, err := repo.GetAllBusStopsById(routeID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(busStops) != 2 {
			t.Errorf("Expected 2 bus stops, got %d", len(busStops))
		}
	})

	t.Run("Get bus stops from non-existent route", func(t *testing.T) {
		_, err := repo.GetAllBusStopsById(uuid.New().String())
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})
}

func TestSqliteRouteRepository_GetAllBusesById(t *testing.T) {
	repo, cleanup := setupTestDBRoute(t)
	defer cleanup()

	routeID := uuid.New().String()
	busID1 := uuid.New().String()
	busID2 := uuid.New().String()

	// Вставляем маршрут
	_, err := repo.db.Exec(`
        INSERT INTO routes (id, number)
        VALUES (?, ?)`,
		routeID, "101")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Вставляем автобусы
	_, err = repo.db.Exec(`
        INSERT INTO buses (id, brand, bus_model, register_number, assembly_date, last_repair_date)
        VALUES (?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?)`,
		busID1, "Mercedes", "Citaro", "X123YZ", "2020-01-01", "2023-06-01",
		busID2, "Volvo", "B8RLE", "Y456AB", "2019-02-01", "2023-06-15")
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Вставляем связи
	_, err = repo.db.Exec(`
        INSERT INTO routes_buses (route_id, bus_id)
        VALUES (?, ?), (?, ?)`,
		routeID, busID1, routeID, busID2)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get all buses by route", func(t *testing.T) {
		buses, err := repo.GetAllBusesById(routeID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(buses) != 2 {
			t.Errorf("Expected 2 buses, got %d", len(buses))
		}
	})

	t.Run("Get buses from non-existent route", func(t *testing.T) {
		_, err := repo.GetAllBusesById(uuid.New().String())
		if err == nil || err.Error() != "Route not found" {
			t.Errorf("Expected 'Route not found' error, got %v", err)
		}
	})
}
