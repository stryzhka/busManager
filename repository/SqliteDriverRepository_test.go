package repository

import (
	"busManager/models"
	"database/sql"
	"github.com/google/uuid"
	"testing"
	"time"
)

func setupTestDBDriver(t *testing.T) (*SqliteDriverRepository, func()) {
	// Use in-memory database for testing
	db, err := sql.Open("sqlite3", ":memory:?parseTime=true")
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	// Create drivers table
	_, err = db.Exec(`
		CREATE TABLE drivers (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			surname TEXT NOT NULL,
			patronymic TEXT NOT NULL,
			birth_date TIMESTAMP NOT NULL,
			passport_series TEXT NOT NULL UNIQUE,
			snils TEXT NOT NULL UNIQUE,
			license_series TEXT NOT NULL UNIQUE
		)
	`)
	if err != nil {
		t.Fatalf("Failed to create drivers table: %v", err)
	}

	repo := &SqliteDriverRepository{db: db}
	return repo, func() { db.Close() }
}

func TestSqliteDriverRepository_GetById(t *testing.T) {
	repo, cleanup := setupTestDBDriver(t)
	defer cleanup()

	fixedTime, _ := time.Parse(time.RFC3339, "2022-11-11T11:11:11Z")

	driver := &models.Driver{
		ID:             uuid.New().String(),
		Name:           "John",
		Surname:        "Doe",
		Patronymic:     "Ivanovich",
		BirthDate:      fixedTime,
		PassportSeries: "AB123456",
		Snils:          "123-456-789 00",
		LicenseSeries:  "CD789012",
	}

	_, err := repo.db.Exec(`
		INSERT INTO drivers (id, name, surname, patronymic, birth_date, passport_series, snils, license_series)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		driver.ID, driver.Name, driver.Surname, driver.Patronymic, driver.BirthDate, driver.PassportSeries, driver.Snils, driver.LicenseSeries)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get existing driver by ID", func(t *testing.T) {
		result, err := repo.GetById(driver.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.ID != driver.ID {
			t.Errorf("Expected driver with ID %s, got %v", driver.ID, result)
		}
		if result != nil && result.BirthDate != driver.BirthDate {
			t.Errorf("Expected BirthDate %v, got %v", driver.BirthDate, result.BirthDate)
		}
	})

	t.Run("Get non-existent driver by ID", func(t *testing.T) {
		_, err := repo.GetById(uuid.New().String())
		if err == nil || err.Error() != "Driver not found" {
			t.Errorf("Expected 'Driver not found' error, got %v", err)
		}
	})
}

func TestSqliteDriverRepository_GetByPassportSeries(t *testing.T) {
	repo, cleanup := setupTestDBDriver(t)
	defer cleanup()

	fixedTime, _ := time.Parse(time.RFC3339, "2022-11-11T11:11:11Z")

	driver := &models.Driver{
		ID:             uuid.New().String(),
		Name:           "John",
		Surname:        "Doe",
		Patronymic:     "Ivanovich",
		BirthDate:      fixedTime,
		PassportSeries: "AB123456",
		Snils:          "123-456-789 00",
		LicenseSeries:  "CD789012",
	}

	_, err := repo.db.Exec(`
		INSERT INTO drivers (id, name, surname, patronymic, birth_date, passport_series, snils, license_series)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		driver.ID, driver.Name, driver.Surname, driver.Patronymic, driver.BirthDate, driver.PassportSeries, driver.Snils, driver.LicenseSeries)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get existing driver by passport series", func(t *testing.T) {
		result, err := repo.GetByPassportSeries(driver.PassportSeries)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.PassportSeries != driver.PassportSeries {
			t.Errorf("Expected driver with passport series %s, got %v", driver.PassportSeries, result)
		}
		if result != nil && result.BirthDate != driver.BirthDate {
			t.Errorf("Expected BirthDate %v, got %v", driver.BirthDate, result.BirthDate)
		}
	})

	t.Run("Get non-existent driver by passport series", func(t *testing.T) {
		_, err := repo.GetByPassportSeries("XY789012")
		if err == nil || err.Error() != "Driver not found" {
			t.Errorf("Expected 'Driver not found' error, got %v", err)
		}
	})
}

func TestSqliteDriverRepository_Add(t *testing.T) {
	repo, cleanup := setupTestDBDriver(t)
	defer cleanup()

	fixedTime, _ := time.Parse(time.RFC3339, "2022-11-11T11:11:11Z")

	driver := &models.Driver{
		Name:           "John",
		Surname:        "Doe",
		Patronymic:     "Ivanovich",
		BirthDate:      fixedTime,
		PassportSeries: "AB123456",
		Snils:          "123-456-789 00",
		LicenseSeries:  "CD789012",
	}

	t.Run("Add new driver", func(t *testing.T) {
		err := repo.Add(driver)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if driver.ID == "" {
			t.Errorf("Expected driver ID to be set")
		}

		// Verify driver was added
		result, err := repo.GetByPassportSeries(driver.PassportSeries)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result == nil || result.PassportSeries != driver.PassportSeries {
			t.Errorf("Expected driver with passport series %s, got %v", driver.PassportSeries, result)
		}
		if result != nil && result.BirthDate != driver.BirthDate {
			t.Errorf("Expected BirthDate %v, got %v", driver.BirthDate, result.BirthDate)
		}
	})

	t.Run("Add duplicate driver", func(t *testing.T) {
		duplicateDriver := &models.Driver{
			Name:           "Jane",
			Surname:        "Doe",
			Patronymic:     "Ivanovna",
			BirthDate:      fixedTime,
			PassportSeries: "AB123456",
			Snils:          "987-654-321 00",
			LicenseSeries:  "EF345678",
		}
		err := repo.Add(duplicateDriver)
		if err == nil || err.Error() != "Driver already exists" {
			t.Errorf("Expected 'Driver already exists' error, got %v", err)
		}
	})
}

func TestSqliteDriverRepository_GetAll(t *testing.T) {
	repo, cleanup := setupTestDBDriver(t)
	defer cleanup()

	fixedTime, _ := time.Parse(time.RFC3339, "2022-11-11T11:11:11Z")

	driver1 := &models.Driver{
		ID:             uuid.New().String(),
		Name:           "John",
		Surname:        "Doe",
		Patronymic:     "Ivanovich",
		BirthDate:      fixedTime,
		PassportSeries: "AB123456",
		Snils:          "123-456-789 00",
		LicenseSeries:  "CD789012",
	}
	driver2 := &models.Driver{
		ID:             uuid.New().String(),
		Name:           "Jane",
		Surname:        "Doe",
		Patronymic:     "Ivanovna",
		BirthDate:      fixedTime,
		PassportSeries: "XY789012",
		Snils:          "987-654-321 00",
		LicenseSeries:  "EF345678",
	}

	_, err := repo.db.Exec(`
		INSERT INTO drivers (id, name, surname, patronymic, birth_date, passport_series, snils, license_series)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?, ?, ?)`,
		driver1.ID, driver1.Name, driver1.Surname, driver1.Patronymic, driver1.BirthDate, driver1.PassportSeries, driver1.Snils, driver1.LicenseSeries,
		driver2.ID, driver2.Name, driver2.Surname, driver2.Patronymic, driver2.BirthDate, driver2.PassportSeries, driver2.Snils, driver2.LicenseSeries)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Get all drivers", func(t *testing.T) {
		drivers, err := repo.GetAll()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(drivers) != 2 {
			t.Errorf("Expected 2 drivers, got %d", len(drivers))
		}
		for _, driver := range drivers {
			if driver.BirthDate != fixedTime {
				t.Errorf("Expected BirthDate %v, got %v", fixedTime, driver.BirthDate)
			}
		}
	})

	t.Run("Get all from empty DB", func(t *testing.T) {
		// Clear the database
		_, err := repo.db.Exec("DELETE FROM drivers")
		if err != nil {
			t.Fatalf("Failed to clear database: %v", err)
		}

		drivers, err := repo.GetAll()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(drivers) != 0 {
			t.Errorf("Expected 0 drivers, got %d", len(drivers))
		}
	})
}

func TestSqliteDriverRepository_DeleteById(t *testing.T) {
	repo, cleanup := setupTestDBDriver(t)
	defer cleanup()

	fixedTime, _ := time.Parse(time.RFC3339, "2022-11-11T11:11:11Z")

	driver := &models.Driver{
		ID:             uuid.New().String(),
		Name:           "John",
		Surname:        "Doe",
		Patronymic:     "Ivanovich",
		BirthDate:      fixedTime,
		PassportSeries: "AB123456",
		Snils:          "123-456-789 00",
		LicenseSeries:  "CD789012",
	}

	_, err := repo.db.Exec(`
		INSERT INTO drivers (id, name, surname, patronymic, birth_date, passport_series, snils, license_series)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		driver.ID, driver.Name, driver.Surname, driver.Patronymic, driver.BirthDate, driver.PassportSeries, driver.Snils, driver.LicenseSeries)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Delete existing driver", func(t *testing.T) {
		err := repo.DeleteById(driver.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Verify driver was deleted
		_, err = repo.GetById(driver.ID)
		if err == nil || err.Error() != "Driver not found" {
			t.Errorf("Expected 'Driver not found' error, got %v", err)
		}
	})

	t.Run("Delete non-existent driver", func(t *testing.T) {
		err := repo.DeleteById(uuid.New().String())
		if err == nil || err.Error() != "Driver not found" {
			t.Errorf("Expected 'Driver not found' error, got %v", err)
		}
	})
}

func TestSqliteDriverRepository_UpdateById(t *testing.T) {
	repo, cleanup := setupTestDBDriver(t)
	defer cleanup()

	fixedTime, _ := time.Parse(time.RFC3339, "2022-11-11T11:11:11Z")

	driver := &models.Driver{
		ID:             uuid.New().String(),
		Name:           "John",
		Surname:        "Doe",
		Patronymic:     "Ivanovich",
		BirthDate:      fixedTime,
		PassportSeries: "AB123456",
		Snils:          "123-456-789 00",
		LicenseSeries:  "CD789012",
	}

	_, err := repo.db.Exec(`
		INSERT INTO drivers (id, name, surname, patronymic, birth_date, passport_series, snils, license_series)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		driver.ID, driver.Name, driver.Surname, driver.Patronymic, driver.BirthDate, driver.PassportSeries, driver.Snils, driver.LicenseSeries)
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	t.Run("Update existing driver", func(t *testing.T) {
		updatedDriver := &models.Driver{
			ID:             driver.ID,
			Name:           "Jane",
			Surname:        "Doe",
			Patronymic:     "Ivanovna",
			BirthDate:      fixedTime,
			PassportSeries: "XY789012",
			Snils:          "987-654-321 00",
			LicenseSeries:  "EF345678",
		}

		err := repo.UpdateById(updatedDriver)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		// Verify update
		result, err := repo.GetById(driver.ID)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
			return
		}
		if result == nil {
			t.Errorf("Expected driver, got nil")
			return
		}
		if result.Name != updatedDriver.Name || result.Surname != updatedDriver.Surname || result.PassportSeries != updatedDriver.PassportSeries {
			t.Errorf("Expected updated driver data, got %v", result)
		}
		if result.BirthDate != updatedDriver.BirthDate {
			t.Errorf("Expected BirthDate %v, got %v", updatedDriver.BirthDate, result.BirthDate)
		}
	})

	t.Run("Update non-existent driver", func(t *testing.T) {
		nonExistentDriver := &models.Driver{
			ID:             uuid.New().String(),
			Name:           "Jane",
			Surname:        "Doe",
			Patronymic:     "Ivanovna",
			BirthDate:      fixedTime,
			PassportSeries: "XY789012",
			Snils:          "987-654-321 00",
			LicenseSeries:  "EF345678",
		}

		err := repo.UpdateById(nonExistentDriver)
		if err == nil || err.Error() != "Driver not found" {
			t.Errorf("Expected 'Driver not found' error, got %v", err)
		}
	})
}
