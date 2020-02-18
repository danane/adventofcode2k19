package wiresv2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	up    = "U"
	right = "R"
	down  = "D"
	left  = "L"
)

type Coords struct {
	X int
	Y int
}

func NewCoords(a, b int) *Coords {
	return &Coords{a, b}
}

type Grid struct{}

func NewGrid() *Grid {
	return &Grid{}
}

func ParseFile(filename string) (wires []string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(fmt.Sprintf("error while opening the file %v", err))
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wires = append(wires, scanner.Text())
	}
	return wires
}

func (g *Grid) DrawLine(cmdLine string) map[Coords]int {
	currSteps := 0
	coordsMap := make(map[Coords]int)
	currCoords := NewCoords(0, 0)

	commands := strings.Split(cmdLine, ",")
	for _, c := range commands {
		g.Execute(c, currCoords, coordsMap, &currSteps)
	}
	return coordsMap
}

func (g *Grid) Execute(cmd string, currentPoint *Coords, coordsMap map[Coords]int, currSteps *int) {
	direction, steps := splitCmd(cmd)
	switch direction {
	case up:
		move(&currentPoint.X, decrement, steps, currentPoint, coordsMap, currSteps)
	case right:
		move(&currentPoint.Y, increment, steps, currentPoint, coordsMap, currSteps)
	case down:
		move(&currentPoint.X, increment, steps, currentPoint, coordsMap, currSteps)
	case left:
		move(&currentPoint.Y, decrement, steps, currentPoint, coordsMap, currSteps)
	default:
		panic(fmt.Sprintf("Unexpected command: %s", direction))
	}
}

func move(variable *int, operation func(varaddr *int), steps int, currentPoint *Coords, coordsMap map[Coords]int, currSteps *int) {
	for i := 1; i <= steps; i++ {
		operation(variable)
		*currSteps++
		coordsMap[*currentPoint] = *currSteps
	}
}

func increment(varaddr *int) {
	*varaddr++
}

func decrement(varaddr *int) {
	*varaddr--
}

func splitCmd(cmd string) (direction string, steps int) {
	direction = string(cmd[0])
	steps, err := strconv.Atoi(cmd[1:])
	if err != nil {
		panic("error while converting to int")
	}
	return
}

func (g *Grid) MinStepsIntersection(map1, map2 map[Coords]int) int {
	var ts int
	minSteps := math.MaxInt64
	for k1, v1 := range map1 {
		v2, ok := map2[k1]
		if ok {
			ts = v1 + v2
			if ts < minSteps {
				minSteps = ts
			}
		}
	}
	return int(minSteps)
}

func (c *Coords) String() string {
	return fmt.Sprintf("x,y = (%d,%d)", c.X, c.Y)
}
