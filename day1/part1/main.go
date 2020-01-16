package main

import (
	"fmt"
	"github.com/danane/adventofcode2k19/day1/part1/fuelrequirement"
)

func main() {
	masses := fuelrequirement.ParseFileToMasses("data/input.txt")
	fmt.Printf("Total fuel required: %d\n", fuelrequirement.TotalFuelRequirement(masses, fuelrequirement.FuelRequired))
}
