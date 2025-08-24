package main

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload a truck cargo", func(t *testing.T) {
			nt := NormalTruck{id: "Normal-truck-1", cargo: 25}
			et := ElectricTruck{id: "Electric-truck-1", cargo: 11, battery: 85}

			err := processTruck(nil, &nt)
			if err != nil {
				log.Fatal("Error processing normal truck", err)
			}

			err = processTruck(nil, &et)
			if err != nil {
				t.Fatal("Error processing electric truck", err)
			}

			// asserting
			if nt.cargo != 0 {
				t.Fatal("Normal truck cargo should be 0")
			}

			if et.battery != 82 {
				t.Fatal("Electric truck battery should be 82")
			}
		})
	})
}
