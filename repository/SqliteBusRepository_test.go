package repository

import (
	"busManager/models"
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"time"
)

func setupTestDB(t *testing.T) (*SqliteBusRepository, func()) {
	// Use in-memory database for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	// Create buses table
	_, err = db.Exec(`
		CREATE TABLE buses (
			id TEXT PRIMARY KEY,
			brand TEXT,
			bus_model TEXT,
			register_number TEXT UNIQUE,
			assembly_date TIMESTAMP,
			last_repair_date TIMESTAMP
		)
	`)
	if err != nil {
		t.Fatalf("Failed to create buses table: %v", err)
	}

	repo := &SqliteBusRepository{db: db}
	return repo, func() { db.Close() }
}

func TestSqliteBusRepository_GetById(t *testing.T) {
	repo, cleanup := setupTestDB(t)
	defer cleanup()

	// Parse fixed timestamp
	fixedTime, _ := time.Parse(time.RFC3339, "2022-11-11T11:11:11Z")

	// Test bus
	bus := &models.Bus{
		ID:             uuid.New().String(),
		Brand:          "Volvo",
		BusModel:       "B7R",
		RegisterNumber: "ABC123",
		AssemblyDate:   fixedTime,
		LastRepairDate: fixedTime, // Explicitly set to avoid NULL
	}

	// Insert test data
	_, err := repo.db.Exec(`
		INSERT INTO buses (id, brand, bus_model, register_number, assembly_date, last_repair_date)
		VALUES (?, ?, ?, ?, ?, ?)`,
		bus.ID, bus.Brand, bus.BusModel, bus.RegisterNumber, bus.AssemblyDate, bus.LastRepairDate)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get existing bus by ID", func(t *testing.T) {
		result, err := repo.GetById(bus.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.ID != bus.ID {
			t.Errorf("Expected bus with ID %s, got %v", bus.ID, result)
		}
		if result != nil && result.LastRepairDate != bus.LastRepairDate {
			t.Errorf("Expected LastRepairDate %v, got %v", bus.LastRepairDate, result.LastRepairDate)
		}
	})

	t.Run("Get non-existent bus by ID", func(t *testing.T) {
		_, err := repo.GetById(uuid.New().String())
		if err == nil || err.Error() != "Bus not found" {
			t.Errorf("Expected 'Bus not found' error, got %v", err)
		}
	})
}

func TestSqliteBusRepository_GetByNumber(t *testing.T) {
	repo, cleanup := setupTestDB(t)
	defer cleanup()

	fixedTime, _ := time.Parse(time.RFC3339, "2022-11-11T11:11:11Z")

	bus := &models.Bus{
		ID:             uuid.New().String(),
		Brand:          "Volvo",
		BusModel:       "B7R",
		RegisterNumber: "ABC123",
		AssemblyDate:   fixedTime,
		LastRepairDate: fixedTime,
	}

	_, err := repo.db.Exec(`
		INSERT INTO buses (id, brand, bus_model, register_number, assembly_date, last_repair_date)
		VALUES (?, ?, ?, ?, ?, ?)`,
		bus.ID, bus.Brand, bus.BusModel, bus.RegisterNumber, bus.AssemblyDate, bus.LastRepairDate)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get existing bus by number", func(t *testing.T) {
		result, err := repo.GetByNumber(bus.RegisterNumber)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.RegisterNumber != bus.RegisterNumber {
			t.Errorf("Expected bus with number %s, got %v", bus.RegisterNumber, result)
		}
		if result != nil && result.LastRepairDate != bus.LastRepairDate {
			t.Errorf("Expected LastRepairDate %v, got %v", bus.LastRepairDate, result.LastRepairDate)
		}
	})

	t.Run("Get non-existent bus by number", func(t *testing.T) {
		_, err := repo.GetByNumber("XYZ789")
		if err == nil || err.Error() != "Bus not found" {
			t.Errorf("Expected 'Bus not found' error, got %v", err)
		}
	})
}

func TestSqliteBusRepository_Add(t *testing.T) {
	repo, cleanup := setupTestDB(t)
	defer cleanup()

	fixedTime, _ := time.Parse(time.RFC3339, "2022-11-11T11:11:11Z")

	bus := &models.Bus{
		Brand:          "Volvo",
		BusModel:       "B7R",
		RegisterNumber: "ABC123",
		AssemblyDate:   fixedTime,
		LastRepairDate: fixedTime,
	}

	t.Run("Add new bus", func(t *testing.T) {
		err := repo.Add(bus)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if bus.ID == "" {
			t.Errorf("Expected bus ID to be set")
		}

		// Verify bus was added
		result, err := repo.GetByNumber(bus.RegisterNumber)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.RegisterNumber != bus.RegisterNumber {
			t.Errorf("Expected bus with number %s, got %v", bus.RegisterNumber, result)
		}
		if result != nil && result.LastRepairDate != bus.LastRepairDate {
			t.Errorf("Expected LastRepairDate %v, got %v", bus.LastRepairDate, result.LastRepairDate)
		}
	})

	t.Run("Add duplicate bus", func(t *testing.T) {
		duplicateBus := &models.Bus{
			Brand:          "Mercedes",
			BusModel:       "Citaro",
			RegisterNumber: "ABC123",
			AssemblyDate:   fixedTime,
			LastRepairDate: fixedTime,
		}
		err := repo.Add(duplicateBus)
		if err == nil || err.Error() != "Bus already exists" {
			t.Errorf("Expected 'Bus already exists' error, got %v", err)
		}
	})
}

func TestSqliteBusRepository_GetAll(t *testing.T) {
	repo, cleanup := setupTestDB(t)
	defer cleanup()

	fixedTime, _ := time.Parse(time.RFC3339, "2022-11-11T11:11:11Z")

	bus1 := &models.Bus{
		ID:             uuid.New().String(),
		Brand:          "Volvo",
		BusModel:       "B7R",
		RegisterNumber: "ABC123",
		AssemblyDate:   fixedTime,
		LastRepairDate: fixedTime,
	}
	bus2 := &models.Bus{
		ID:             uuid.New().String(),
		Brand:          "Mercedes",
		BusModel:       "Citaro",
		RegisterNumber: "XYZ789",
		AssemblyDate:   fixedTime,
		LastRepairDate: fixedTime,
	}

	_, err := repo.db.Exec(`
		INSERT INTO buses (id, brand, bus_model, register_number, assembly_date, last_repair_date)
		VALUES (?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?)`,
		bus1.ID, bus1.Brand, bus1.BusModel, bus1.RegisterNumber, bus1.AssemblyDate, bus1.LastRepairDate,
		bus2.ID, bus2.Brand, bus2.BusModel, bus2.RegisterNumber, bus2.AssemblyDate, bus2.LastRepairDate)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get all buses", func(t *testing.T) {
		buses, err := repo.GetAll()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(buses) != 2 {
			t.Errorf("Expected 2 buses, got %d", len(buses))
		}
		for _, bus := range buses {
			if bus.LastRepairDate != fixedTime {
				t.Errorf("Expected LastRepairDate %v, got %v", fixedTime, bus.LastRepairDate)
			}
		}
	})

	t.Run("Get all from empty DB", func(t *testing.T) {
		// Clear the database
		_, err := repo.db.Exec("DELETE FROM buses")
		if err != nil {
			t.Fatalf("Failed to clear database: %v", err)
		}

		buses, err := repo.GetAll()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(buses) != 0 {
			t.Errorf("Expected 0 buses, got %d", len(buses))
		}
	})
}

func TestSqliteBusRepository_DeleteById(t *testing.T) {
	repo, cleanup := setupTestDB(t)
	defer cleanup()

	fixedTime, _ := time.Parse(time.RFC3339, "2022-11-11T11:11:11Z")

	bus := &models.Bus{
		ID:             uuid.New().String(),
		Brand:          "Volvo",
		BusModel:       "B7R",
		RegisterNumber: "ABC123",
		AssemblyDate:   fixedTime,
		LastRepairDate: fixedTime,
	}

	_, err := repo.db.Exec(`
		INSERT INTO buses (id, brand, bus_model, register_number, assembly_date, last_repair_date)
		VALUES (?, ?, ?, ?, ?, ?)`,
		bus.ID, bus.Brand, bus.BusModel, bus.RegisterNumber, bus.AssemblyDate, bus.LastRepairDate)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Delete existing bus", func(t *testing.T) {
		err := repo.DeleteById(bus.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Verify bus was deleted
		_, err = repo.GetById(bus.ID)
		if err == nil || err.Error() != "Bus not found" {
			t.Errorf("Expected 'Bus not found' error, got %v", err)
		}
	})

	t.Run("Delete non-existent bus", func(t *testing.T) {
		err := repo.DeleteById(uuid.New().String())
		if err == nil || err.Error() != "Bus not found" {
			t.Errorf("Expected 'Bus not found' error, got %v", err)
		}
	})
}

func TestSqliteBusRepository_UpdateById(t *testing.T) {
	repo, cleanup := setupTestDB(t)
	defer cleanup()

	fixedTime, _ := time.Parse(time.RFC3339, "2022-11-11T11:11:11Z")

	bus := &models.Bus{
		ID:             uuid.New().String(),
		Brand:          "Volvo",
		BusModel:       "B7R",
		RegisterNumber: "ABC123",
		AssemblyDate:   fixedTime,
		LastRepairDate: fixedTime,
	}

	_, err := repo.db.Exec(`
		INSERT INTO buses (id, brand, bus_model, register_number, assembly_date, last_repair_date)
		VALUES (?, ?, ?, ?, ?, ?)`,
		bus.ID, bus.Brand, bus.BusModel, bus.RegisterNumber, bus.AssemblyDate, bus.LastRepairDate)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Update existing bus", func(t *testing.T) {
		updatedBus := &models.Bus{
			ID:             bus.ID,
			Brand:          "Mercedes",
			BusModel:       "Citaro",
			RegisterNumber: "XYZ789",
			AssemblyDate:   fixedTime,
			LastRepairDate: fixedTime,
		}

		err := repo.UpdateById(updatedBus)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Verify update
		result, err := repo.GetById(bus.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
			return
		}
		if result == nil {
			t.Errorf("Expected bus, got nil")
			return
		}
		if result.Brand != updatedBus.Brand || result.BusModel != updatedBus.BusModel || result.RegisterNumber != updatedBus.RegisterNumber {
			t.Errorf("Expected updated bus data, got %v", result)
		}
		if result.LastRepairDate != updatedBus.LastRepairDate {
			t.Errorf("Expected LastRepairDate %v, got %v", updatedBus.LastRepairDate, result.LastRepairDate)
		}
	})

	t.Run("Update non-existent bus", func(t *testing.T) {
		nonExistentBus := &models.Bus{
			ID:             uuid.New().String(),
			Brand:          "Mercedes",
			BusModel:       "Citaro",
			RegisterNumber: "XYZ789",
			AssemblyDate:   fixedTime,
			LastRepairDate: fixedTime,
		}

		err := repo.UpdateById(nonExistentBus)
		if err == nil || err.Error() != "Bus not found" {
			t.Errorf("Expected 'Bus not found' error, got %v", err)
		}
	})
}
