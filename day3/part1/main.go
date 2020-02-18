package main

import (
	"fmt"

	"github.com/danane/adventofcode2k19/day3/part1/wires"
)

const filePath = "data/input.txt"

func main() {
	lines := wires.ParseFile(filePath)
	g := wires.NewGrid()

	map1 := g.DrawLine(lines[0])
	map2 := g.DrawLine(lines[1])

	g.CalculateIntersection(map1, map2)
	fmt.Printf("Manhattan distance from central port to closest intersection = %.0f\n", g.GetDistanceCloserIntersection())
}
