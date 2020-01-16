package fuelrequirementv2

import (
	"testing"
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
			got := FuelRequiredUpdated(c.mass)

			if got != c.fuel {
				t.Errorf("got %d but expected %d", got, c.fuel)
			}
		})
	}
}
