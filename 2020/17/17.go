package main

import (
	"bufio"
	"log"
	"os"
)

// -
const (
	Active   = '#'
	Inactive = '.'
)

// Grid -
type Grid struct {
	fields             []bool
	w, h, d, t, offset int
}

func createGrid(w, h, d, t, offset int) *Grid {
	return &Grid{
		fields: make([]bool, w*h*d*t),
		w:      w,
		h:      h,
		d:      d,
		t:      t,
		offset: offset,
	}
}

func (g *Grid) isIndexValid(x, y, z, t int) bool {
	if x+g.offset < 0 || x+g.offset > g.w-1 || y+g.offset < 0 || y+g.offset > g.h-1 || z+g.offset < 0 || z+g.offset > g.d-1 || t+g.offset < 0 || t+g.offset > g.t-1 {
		return false
	}
	return true
}

func (g *Grid) at(x, y, z, t int) bool {
	if !g.isIndexValid(x, y, z, t) {
		return false
	}

	idx := g.h*g.w*g.d*(t+g.offset) + g.h*g.w*(z+g.offset) + g.w*(y+g.offset) + x + g.offset
	if idx < 0 || idx > g.t*g.w*g.h*g.d-1 {
		return false
	}
	return g.fields[idx]
}

func (g *Grid) set(x, y, z, t int, val bool) {
	if !g.isIndexValid(x, y, z, t) {
		print("Something wrong")
	}
	idx := g.h*g.w*g.d*(t+g.offset) + g.h*g.w*(z+g.offset) + g.w*(y+g.offset) + x + g.offset
	g.fields[idx] = val
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func parseInput(lines []string) *Grid {
	grid := createGrid(len(lines[0]), len(lines), 1, 1, 0)
	for y, line := range lines {
		for x, field := range line {
			grid.set(x, y, 0, 0, field == Active)
		}
	}
	return grid
}

// func printGrid(grid *Grid) {
// 	for t := 0; t < grid.t; t++ {
// 		for z := 0; z < grid.d; z++ {
// 			for y := 0; y < grid.h; y++ {
// 				for x := 0; x < grid.w; x++ {
// 					if grid.at(x-grid.offset, y-grid.offset, z-grid.offset, t-grid.offset) {
// 						print(string(Active))
// 					} else {
// 						print(string(Inactive))
// 					}
// 				}
// 				println()
// 			}
// 			println()
// 		}
// 	}
// }

func getNewCubeState(grid *Grid, x, y, z, t int) bool {
	activeNeighbors := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				for l := t - 1; l <= t+1; l++ {
					if x != i || y != j || z != k || l != t {
						if grid.at(i, j, k, l) {
							activeNeighbors++
						}
					}
				}
			}
		}
	}
	currentlyActive := grid.at(x, y, z, t)
	if currentlyActive && (activeNeighbors == 2 || activeNeighbors == 3) {
		return true
	}
	if !currentlyActive && activeNeighbors == 3 {
		return true
	}
	return false
}

func expandGrid(grid *Grid, offset int) (*Grid, int) {
	newGrid := createGrid(grid.w+2, grid.h+2, grid.d+2, grid.t+2, offset)

	activeSum := 0
	for t := 0; t < newGrid.t; t++ {
		for z := 0; z < newGrid.d; z++ {
			for y := 0; y < newGrid.h; y++ {
				for x := 0; x < newGrid.w; x++ {
					active := getNewCubeState(grid, x-offset, y-offset, z-offset, t-offset)
					newGrid.set(x-newGrid.offset, y-newGrid.offset, z-newGrid.offset, t-newGrid.offset, active)
					if active {
						activeSum++
					}
				}
			}
		}
	}

	return newGrid, activeSum
}

func part1(lines []string) int {
	grid := parseInput(lines)
	// printGrid(grid)
	activeSum := 0
	for i := 0; i < 6; i++ {
		grid, activeSum = expandGrid(grid, (i + 1))
		// printGrid(grid)
	}
	return activeSum
}

func part2(lines []string) int {
	return 0
}

func main() {
	lines := readLines("17.txt")
	println(part1(lines))
}
