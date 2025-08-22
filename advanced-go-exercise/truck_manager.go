package main

import "errors"

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
}

func (tManager *truckManager) AddTruck(id string, cargo int) error {
	tManager.trucks[id] = &Truck{ID: id, Cargo: cargo}

	return nil
}

func (tManager *truckManager) GetTruck(id string) (*Truck, error) {
	truck, exists := tManager.trucks[id]

	if !exists {
		return nil, ErrTruckNotFound
	}

	return truck, nil
}

func (tManager *truckManager) RemoveTruck(id string) error {
	if _, exists := tManager.trucks[id]; !exists {
		return ErrTruckNotFound
	}

	delete(tManager.trucks, id)
	return nil
}

func (tManager *truckManager) UpdateTruckCargo(id string, cargo int) error {
	if _, exists := tManager.trucks[id]; !exists {
		return ErrTruckNotFound
	}

	tManager.trucks[id].Cargo = cargo
	return nil
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}
