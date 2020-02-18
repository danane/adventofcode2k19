package wires

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

type Grid struct {
	Intersections []Coords
}

func NewGrid() *Grid {
	return &Grid{}
}

func ParseFile(filename string) (wires []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("error while opening the file %v", err))
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wires = append(wires, scanner.Text())
	}
	return wires
}

func (g *Grid) DrawLine(cmdLine string) map[Coords]bool {
	coordsMap := make(map[Coords]bool)
	currCoords := &Coords{0, 0}

	commands := strings.Split(cmdLine, ",")
	for _, c := range commands {
		g.Execute(c, currCoords, coordsMap)
	}
	return coordsMap
}

func (g *Grid) Execute(cmd string, currentPoint *Coords, coordsMap map[Coords]bool) {
	direction, steps := splitCmd(cmd)
	switch direction {
	case up:
		move(&currentPoint.X, decrement, steps, currentPoint, coordsMap)
	case right:
		move(&currentPoint.Y, increment, steps, currentPoint, coordsMap)
	case down:
		move(&currentPoint.X, increment, steps, currentPoint, coordsMap)
	case left:
		move(&currentPoint.Y, decrement, steps, currentPoint, coordsMap)
	default:
		panic(fmt.Sprintf("Unexpected command: %s", direction))
	}
}

func move(variable *int, operation func(varaddr *int), steps int, currentPoint *Coords, coordsMap map[Coords]bool) {
	for i := 0; i < steps; i++ {
		operation(variable)
		coordsMap[*currentPoint] = true
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
	stringSteps, _ := strconv.Atoi(cmd[1:])
	steps = int(stringSteps)
	return
}

func (g *Grid) CalculateIntersection(map1, map2 map[Coords]bool) {
	for k := range map1 {
		_, ok := map2[k]
		if ok {
			g.Intersections = append(g.Intersections, k)
		}
	}
}

func ManhattanDistance(origin, p Coords) float64 {
	return math.Abs(float64(origin.X-p.X)) + math.Abs(float64(origin.Y-p.Y))
}

func (g *Grid) GetDistanceCloserIntersection() (minDistance float64) {
	var mDistance float64
	origin := Coords{0, 0}

	minDistance = math.Inf(0)
	for _, coords := range g.Intersections {
		mDistance = ManhattanDistance(origin, coords)
		if minDistance > mDistance {
			minDistance = mDistance
		}
	}
	return
}

func (c *Coords) String() string {
	return fmt.Sprintf("x,y = (%d,%d)", c.X, c.Y)
}
