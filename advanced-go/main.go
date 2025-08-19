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

type Truck struct {
	id string
}

func (t Truck) loadCargo() error {
	return nil
}

func (t Truck) unloadCargo() error {
	return nil
}

// processTruck handles the loading and unloading of a truck.
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.id)

	if err := truck.loadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.unloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	return nil
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)

		if err := processTruck(truck); err != nil {
			if errors.Is(err, ErrTruckNotFound) {
				log.Fatal("TRUE")
			}
			log.Fatalf("Error processing truck: %s", err)

		}
	}
}
