package repository

import (
	"busManager/models"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"strings"
)

type SqliteRouteRepository struct {
	db *sql.DB
}

func NewSqliteRouteRepository(dbPath string) (*SqliteRouteRepository, error) {
	db, err := sql.Open("sqlite3", dbPath+"?parseTime=true")
	//defer db.Close()
	if err != nil {
		return nil, err
	}
	repo := &SqliteRouteRepository{db: db}
	return repo, nil
}

func (r *SqliteRouteRepository) GetById(id string) (*models.Route, error) {
	route := &models.Route{}
	err := r.db.QueryRow(`
		SELECT id, number 
		FROM routes 
		WHERE id = $1`, id).Scan(
		&route.ID,
		&route.Number,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Route not found")
		}
		return nil, err
	}

	return route, nil
}

func (r *SqliteRouteRepository) GetByNumber(number string) (*models.Route, error) {
	route := &models.Route{}
	err := r.db.QueryRow(`
		SELECT id, number
		FROM routes 
		WHERE number = $1`, number).Scan(
		&route.ID,
		&route.Number,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Route not found")
		}
		return nil, err
	}

	return route, nil
}

func (r *SqliteRouteRepository) Add(route *models.Route) error {
	exist, err := r.GetByNumber(route.Number)
	if exist != nil {
		return errors.New("Route already exists")
	}
	if strings.TrimSpace(route.ID) == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		route.ID = id.String()
	}
	_, err = r.db.Exec(`INSERT into routes (id, number) 
VALUES ($1, $2)`, &route.ID,
		&route.Number,
	)
	if err != nil {
		return err
	}
	return nil

}

func (r *SqliteRouteRepository) GetAll() ([]models.Route, error) {
	var routes []models.Route
	rows, err := r.db.Query(`
		SELECT id, number
		FROM routes 
		`)
	if err != nil {
		//if err == sql.ErrNoRows {
		//	return nil, errors.New("Empty DB")
		//}
		return nil, err
	}
	for rows.Next() {
		route := &models.Route{}
		err := rows.Scan(
			&route.ID,
			&route.Number,
		)
		if err != nil {
			return nil, err
		}
		routes = append(routes, *route)

	}
	return routes, nil
}

func (r *SqliteRouteRepository) DeleteById(id string) error {
	exist, err := r.GetById(id)
	if exist == nil {
		return errors.New("Route not found")
	}
	if err != nil {
		return err
	}
	_, err = r.db.Exec("DELETE FROM routes WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *SqliteRouteRepository) UpdateById(route *models.Route) error {

	exist, err := r.GetById(route.ID)
	if exist == nil {
		return errors.New("Route not found")
	}
	if err != nil {
		return err
	}
	_, err = r.db.Exec("UPDATE routes SET number = $1 WHERE id = $2",
		route.Number, route.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *SqliteRouteRepository) AssignDriver(routeId, driverId string) error {
	exist, err := r.GetById(routeId)
	if exist == nil {
		return errors.New("Route not found")
	}
	_, err = r.db.Exec(`INSERT into routes_drivers (route_id, driver_id) 
VALUES ($1, $2)`, routeId,
		driverId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *SqliteRouteRepository) AssignBusStop(routeId, busStopId string) error {
	exist, err := r.GetById(routeId)
	if exist == nil {
		return errors.New("Route not found")
	}
	_, err = r.db.Exec(`INSERT into routes_bus_stops (route_id, bus_stop_id) 
VALUES ($1, $2)`, routeId,
		busStopId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *SqliteRouteRepository) AssignBus(routeId, busId string) error {
	exist, err := r.GetById(routeId)
	if exist == nil {
		return errors.New("Route not found")
	}
	_, err = r.db.Exec(`INSERT into routes_buses (route_id, bus_id) 
VALUES ($1, $2)`, routeId,
		busId,
	)
	if err != nil {
		return err
	}
	return nil
}
