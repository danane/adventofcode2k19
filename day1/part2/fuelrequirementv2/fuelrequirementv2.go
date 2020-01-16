package fuelrequirementv2

import (
	"github.com/danane/adventofcode2k19/day1/part1/fuelrequirement"
)

func FuelRequiredUpdated(mass int) (total int) {
	for {
		fr := fuelrequirement.FuelRequired(mass)
		if fr <= 0 {
			return
		}
		total += fr
		mass = fr
	}
}
