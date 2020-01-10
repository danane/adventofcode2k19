package main

import (
	"fmt"
	fuelmasses "github.com/danane/adventofcode2k19/day1/part1"
	fuelmassesv2 "github.com/danane/adventofcode2k19/day1/part2"
)

func main() {
	masses := fuelmasses.ParseFileToMasses("../data/input.txt")
	fmt.Printf("Total fuel required: %d\n", fuelmasses.TotalFuelRequirement(masses, fuelmassesv2.FuelRequiredUpdated))
}
