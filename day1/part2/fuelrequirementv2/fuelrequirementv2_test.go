package fuelrequirementv2_test

import (
	"testing"

	"github.com/danane/adventofcode2k19/day1/part2/fuelrequirementv2"
)

func TestFuelRequiredUpdated(t *testing.T) {
	cases := []struct {
		description string
		mass        int
		fuel        int
	}{
		{"1st test", 14, 2},
		{"2nd test", 1969, 966},
		{"3rd test", 100756, 50346},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got := fuelrequirementv2.FuelRequiredUpdated(c.mass)

			if got != c.fuel {
				t.Errorf("got %d but expected %d", got, c.fuel)
			}
		})
	}
}
