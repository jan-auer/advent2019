package main

import (
	"fmt"
)

func requiredFuel(mass int) int {
	return mass/3 - 2
}

func requiredFuelIncluding(mass int) (fuel int) {
	for mass > 0 {
		mass = requiredFuel(mass)

		if mass < 0 {
			mass = 0
		}

		fuel += mass
	}

	return
}

func main() {
	total := 0

	for {
		var mass int

		_, err := fmt.Scanln(&mass)
		if err != nil {
			break
		}

		// total += requiredFuel(mass)
		total += requiredFuelIncluding(mass)
	}

	fmt.Println("Required fuel:", total)
}
