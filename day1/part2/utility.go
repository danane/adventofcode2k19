package fuelmassesv2

import (
	fuelmasses "github.com/danane/adventofcode2k19/day1/part1"
)

func FuelRequiredUpdated(mass int) (total int) {
	for {
		fr := fuelmasses.FuelRequired(mass)
		if fr <= 0 {
			return
		}
		total += fr
		mass = fr
	}
}
