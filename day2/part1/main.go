package main

import (
	"fmt"
	"os"

	"github.com/danane/adventofcode2k19/day2/part1/alarm"
)

func main() {
	file, err := os.Open("data/input.txt")
	if err != nil {
		panic("Error while opening the input file provided.")
	}

	intcodes := alarm.IntcodeToSlice(alarm.ParseIntcode(file))

	intcodes[1] = 12
	intcodes[2] = 2

	fmt.Println(alarm.ComputeOutput(intcodes))

}
