package models

import (
	"time"
)

type Bus struct {
	ID             string
	Brand          string
	BusModel       string
	RegisterNumber string
	AssemblyDate   time.Time
	LastRepairDate time.Time
}

//func NewBus(id, brand, model, registerNumber string, assemblyDate, lastRepairDate)
