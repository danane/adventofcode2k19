package alarmv2

import "github.com/danane/adventofcode2k19/day2/part1/alarm"

func CalculateNounVerb(memory []int, outputToProduce int) (int, int) {
	memcopy := make([]int, len(memory))
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			copy(memcopy, memory)
			memcopy[1] = i
			memcopy[2] = j
			output := alarm.ComputeOutput(memcopy)
			if output == outputToProduce {
				return i, j
			}
		}
	}
	return -1, -1
}
