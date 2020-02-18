package main

import (
	"fmt"

	"github.com/danane/adventofcode2k19/day3/part2/wiresv2"
)

const filePath = "data/input.txt"

func main() {
	lines := wiresv2.ParseFile(filePath)
	g := wiresv2.NewGrid()

	map1 := g.DrawLine(lines[0])
	map2 := g.DrawLine(lines[1])

	fmt.Printf("Fewest combined steps the wires must take to reach an intersection = %d\n", g.MinStepsIntersection(map1, map2))
}
