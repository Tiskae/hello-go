package main

import (
	"errors"
	"sync"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck

	sync.RWMutex
}

func (tManager *truckManager) AddTruck(id string, cargo int) error {
	tManager.Lock()
	defer tManager.Unlock()

	tManager.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (tManager *truckManager) GetTruck(id string) (Truck, error) {
	tManager.RLock()
	defer tManager.RUnlock()

	truck, exists := tManager.trucks[id]

	if !exists {
		return Truck{}, ErrTruckNotFound
	}

	return *truck, nil
}

func (tManager *truckManager) RemoveTruck(id string) error {
	tManager.Lock()
	defer tManager.Unlock()

	if _, exists := tManager.trucks[id]; !exists {
		return ErrTruckNotFound
	}

	delete(tManager.trucks, id)
	return nil
}

func (tManager *truckManager) UpdateTruckCargo(id string, cargo int) error {
	tManager.Lock()
	defer tManager.Unlock()

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
