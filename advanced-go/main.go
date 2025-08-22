package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type ContextKeyType string

var UserIDKey ContextKeyType = "userID"

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
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
	id      string
	cargo   int
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
func processTruck(ctx context.Context, truck Truck) error {
	fmt.Printf("Started processing truck %+v \n", truck)

	// accessing the user ID
	userID := ctx.Value(UserIDKey)
	fmt.Println(map[string]interface{}{"userID": userID})

	// Timeouts
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	// simulate a long running process
	delay := time.Second * 1
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	fmt.Printf("Finished processing truck %+v \n", truck)

	return ErrTruckNotFound
}

// processFleet demonstrates concurrent processing of multiple trucks
func processFleet(ctx context.Context, trucks []Truck) error {
	var wg sync.WaitGroup
	errorsChan := make(chan error, len(trucks))
	errors := []error{}

	for _, truck := range trucks {
		wg.Add(1)

		go func(t Truck) {
			if err := processTruck(ctx, t); err != nil {
				log.Println(err)
				errorsChan <- err
			}
			wg.Done()
		}(truck) // IIFE
	}

	wg.Wait()
	close(errorsChan)

	// select {
	// case err := <- errorsChan:
	// 	return err
	// default:
	// 	return nil
	// }

	for err := range errorsChan {
		log.Printf("Error processing truck: %v\n", err)
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return fmt.Errorf("fleet processing had %d errors", len(errors))
	}

	return nil
}

func main() {

	// contexts are immutable
	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIDKey, 42)

	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
	}

	// process all trucks concurrently
	if err := processFleet(ctx, fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Println("All trucks processed successfully!")
}
