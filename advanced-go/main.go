package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id string
	cargo int
}

// In struct methods, instruct methods (pun intended) to take pointers when
// you intend to change a property of the original struct on the fly  ...
func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

type ElectricTruck struct {
	id string
	cargo int
	battery float64
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery -= 2
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery -= 1
	return nil
}

// processTruck handles the loading and unloading of a truck.
func processTruck(truck Truck) error {
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	return nil
}

func main() {
	nt := NormalTruck{id: "Normal-truck-1", cargo: 25}
	et := ElectricTruck{id: "Electric-truck-1", cargo: 11, battery: 85}

	err := processTruck(&nt)
	if err != nil {
		log.Fatal("Error processing normal truck", err)
	}

	err = processTruck(&et)
	if err != nil {
		log.Fatal("Error processing electric truck", err)
	}

	fmt.Printf("%+v \n", nt)
	fmt.Printf("%+v \n", et)

	truckID := 42
	truckIdReference := &truckID

	truckID = 50

	log.Println(truckID)
	log.Println(*truckIdReference)

	nt2 := NormalTruck{id: "20", cargo: 12}

	fillTruckCargo(nil)

	log.Println(nt2.cargo)
}

func fillTruckCargo(t *NormalTruck) {
	t.cargo = 50
}