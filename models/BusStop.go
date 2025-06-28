package models

type BusStop struct {
	ID      string
	RouteId string
	Lat     float64
	Long    float64
	Order   int
	Name    string
}
