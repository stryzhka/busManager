package repository

import (
	"busManager/models"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"strings"
)

type SqliteBusStopRepository struct {
	db *sql.DB
}

func NewSqliteBusStopRepository(dbPath string) (*SqliteBusStopRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	repo := &SqliteBusStopRepository{db: db}
	return repo, nil
}

func (r *SqliteBusStopRepository) GetById(id string) (*models.BusStop, error) {
	stop := &models.BusStop{}
	err := r.db.QueryRow(`
		SELECT id, route_id, lat, long, "order", name 
		FROM bus_stops 
		WHERE id = $1`, id).Scan(
		&stop.ID,
		&stop.RouteId,
		&stop.Lat,
		&stop.Long,
		&stop.Order,
		&stop.Name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Bus stop not found")
		}
		return nil, err
	}

	return stop, nil
}

func (r *SqliteBusStopRepository) GetByName(name string) (*models.BusStop, error) {
	stop := &models.BusStop{}
	err := r.db.QueryRow(`
		SELECT id, route_id, lat, long, "order", name 
		FROM bus_stops 
		WHERE name = $6`, name).Scan(
		&stop.ID,
		&stop.RouteId,
		&stop.Lat,
		&stop.Long,
		&stop.Order,
		&stop.Name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Bus stop not found")
		}
		return nil, err
	}

	return stop, nil
}

func (r *SqliteBusStopRepository) Add(busStop *models.BusStop) error {
	exist, err := r.GetByName(busStop.Name)
	if exist != nil {
		return errors.New("Bus stop already exists")
	}
	if strings.TrimSpace(busStop.ID) == "" {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		busStop.ID = id.String()
	}
	_, err = r.db.Exec(`INSERT into bus_stops
    (id, route_id, lat, long, "order", name ) 
VALUES ($1, $2, $3, $4, $5, $6)`,
		&busStop.ID,
		&busStop.RouteId,
		&busStop.Lat,
		&busStop.Long,
		&busStop.Order,
		&busStop.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *SqliteBusStopRepository) GetAll() ([]models.BusStop, error) {
	var busStops []models.BusStop
	rows, err := r.db.Query(`
		SELECT id, route_id, lat, long, "order", name
		FROM bus_stops 
		`)
	if err != nil {
		//if err == sql.ErrNoRows {
		//	return nil, errors.New("Empty DB")
		//}
		return nil, err
	}
	for rows.Next() {
		busStop := &models.BusStop{}
		err := rows.Scan(
			&busStop.ID,
			&busStop.RouteId,
			&busStop.Lat,
			&busStop.Long,
			&busStop.Order,
			&busStop.Name,
		)
		if err != nil {
			return nil, err
		}
		busStops = append(busStops, *busStop)

	}
	return busStops, nil
}

func (r *SqliteBusStopRepository) GetAllByRouteId(routeId string) ([]models.BusStop, error) {
	var busStops []models.BusStop
	rows, err := r.db.Query(`
		SELECT id, route_id, lat, long, "order", name
		FROM bus_stops 
		WHERE route_id = $1
		`, routeId)
	if err != nil {
		//if err == sql.ErrNoRows {
		//	return nil, errors.New("Empty DB")
		//}
		return nil, err
	}
	for rows.Next() {
		busStop := &models.BusStop{}
		err := rows.Scan(
			&busStop.ID,
			&busStop.RouteId,
			&busStop.Lat,
			&busStop.Long,
			&busStop.Order,
			&busStop.Name,
		)
		if err != nil {
			return nil, err
		}
		busStops = append(busStops, *busStop)

	}
	return busStops, nil
}

func (r *SqliteBusStopRepository) DeleteById(id string) error {
	exist, err := r.GetById(id)
	if exist == nil {
		return errors.New("Bus stop not found")
	}
	if err != nil {
		return err
	}
	_, err = r.db.Exec("DELETE FROM bus_stops WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *SqliteBusStopRepository) UpdateById(busStop *models.BusStop) error {
	exist, err := r.GetById(busStop.ID)
	if exist == nil {
		return errors.New("Bus stop not found")
	}
	if err != nil {
		return err
	}
	_, err = r.db.Exec(`UPDATE bus_stops SET route_id = $1, lat = $2, long = $3, "order" = $4, name = $5 WHERE id = $6`,
		busStop.RouteId,
		busStop.Lat,
		busStop.Long,
		busStop.Order,
		busStop.Name,
		busStop.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
