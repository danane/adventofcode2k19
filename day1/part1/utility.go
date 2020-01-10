package fuelmasses

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func ParseFileToMasses(filename string) (masses []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("couldn't open %s", filename))
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("error while scanning the values from the file")
		}
		masses = append(masses, value)
	}
	return
}

func FuelRequired(mass int) int {
	massF := float64(mass)
	return int(math.Floor(massF/3) - 2)
}

func TotalFuelRequirement(masses []int, fr func(mass int) int) (total int) {
	for _, mass := range masses {
		total += fr(mass)
	}
	return
}
