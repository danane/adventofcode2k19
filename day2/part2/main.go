package main

import (
	"fmt"
	"os"

	"github.com/danane/adventofcode2k19/day2/part1/alarm"
	"github.com/danane/adventofcode2k19/day2/part2/alarmv2"
)

const outputToProduce = 19690720

func main() {
	file, err := os.Open("data/input.txt")
	if err != nil {
		panic("Error while opening the input file provided")
	}

	intcodes := alarm.IntcodeToSlice(alarm.ParseIntcode(file))

	noun, verb := alarmv2.CalculateNounVerb(intcodes, outputToProduce)

	fmt.Printf("noun=%d, verb=%d\n", noun, verb)
	fmt.Printf("What is 100 * noun + verb? %d\n", 100*noun+verb)
}
