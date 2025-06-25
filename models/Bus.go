package models

import "time"

type Bus struct {
	Id             string
	Brand          string
	Model          string
	RegisterNumber string
	AssemblyDate   time.Time
	LastRepairDate time.Time
}

//func NewBus(id, brand, model, registerNumber string, assemblyDate, lastRepairDate)
