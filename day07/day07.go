package day07

import (
	_ "embed"
	"log"
	"strings"
)

//go:embed input0.txt
var input0 string

//go:embed input1.txt
var input1 string

type XY struct {
	X, Y int
}

func Part1(input string) int {
	grid := make(map[XY]byte)
	limX, limY := 0, 0
	for y, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		limX = len(line)
		limY = y + 1
		for x := range line {
			grid[XY{x, y}] = line[x]
		}
	}
	var beamSplits int
	// start on second row, because we'll be looking a cell above (x,y)
	for y := 1; y < limY; y++ {
		for x := 0; x < limX; x++ {
			aboveCh := grid[XY{x, y - 1}]

			beamFromAbove := aboveCh == 'S' || aboveCh == '|'
			if !beamFromAbove {
				continue
			}
			// aboveCh is 'S' or '|'

			currCh := grid[XY{x, y}]
			if currCh == '.' {
				grid[XY{x, y}] = '|'
				continue
			}
			if currCh == '|' {
				continue
			}
			if currCh == '^' {
				beamSplits++
				// assume here that there are no two adjacent '^' cells
				grid[XY{x - 1, y}] = '|'
				grid[XY{x + 1, y}] = '|'
				continue
			}
			log.Panicf("invalid cell %c", grid[XY{x, y}])
		}
	}
	return beamSplits
}

func Part2(input string) int {
	grid := make(map[XY]byte)
	limX, limY := 0, 0
	for y, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}
		limX = len(line)
		limY = y + 1
		for x := range line {
			grid[XY{x, y}] = line[x]
		}
	}
	// start on second row, because we'll be looking a cell above (x,y)
	for y := 1; y < limY; y++ {
		for x := 0; x < limX; x++ {
			aboveCh := grid[XY{x, y - 1}]

			beamFromAbove := aboveCh == 'S' || aboveCh == '|'
			if !beamFromAbove {
				continue
			}
			// aboveCh is 'S' or '|'

			currCh := grid[XY{x, y}]
			if currCh == '.' {
				grid[XY{x, y}] = '|'
				continue
			}
			if currCh == '|' {
				continue
			}
			if currCh == '^' {
				// assume here that there are no two adjacent '^' cells
				grid[XY{x - 1, y}] = '|'
				grid[XY{x + 1, y}] = '|'
				continue
			}
			log.Panicf("invalid cell %c", grid[XY{x, y}])
		}
	}

	var routeCounts int
	for x := 0; x < limX; x++ {
		routeCounts += routesToPosition(XY{x, limY - 1}, grid)
	}
	return routeCounts
}

var cachedRoutes = make(map[XY]int)

func routesToPosition(xy XY, grid map[XY]byte) int {
	if cached, exists := cachedRoutes[xy]; exists {
		return cached
	}
	//log.Println("considering", xy)
	if grid[xy] == '^' || grid[xy] == '.' {
		return 0
	}
	if grid[xy] == 'S' {
		return 1
	}
	if grid[xy] != '|' {
		log.Panicf("invalid cell %c", grid[xy])
	}
	// where did we come from? could've been any of the
	// three cells above us
	var routeCount int
	aboveMid := XY{xy.X, xy.Y - 1}
	if grid[aboveMid] == '|' || grid[aboveMid] == 'S' {
		routeCount += routesToPosition(aboveMid, grid)
	}
	// check left
	if grid[XY{xy.X - 1, xy.Y}] == '^' {
		routeCount += routesToPosition(XY{xy.X - 1, xy.Y - 1}, grid)
	}
	// check right
	if grid[XY{xy.X + 1, xy.Y}] == '^' {
		routeCount += routesToPosition(XY{xy.X + 1, xy.Y - 1}, grid)
	}
	cachedRoutes[xy] = routeCount
	//log.Println("cached", xy, routeCount)
	return routeCount
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
