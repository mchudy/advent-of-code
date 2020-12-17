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
	fields          []bool
	w, h, d, offset int
}

func createGrid(w, h, d, offset int) *Grid {
	return &Grid{
		fields: make([]bool, w*h*d),
		w:      w,
		h:      h,
		d:      d,
		offset: offset,
	}
}

func (t *Grid) at(x, y, z int) bool {
	if x+t.offset < 0 || x+t.offset > t.w-1 || y+t.offset < 0 || y+t.offset > t.h-1 || z+t.offset < 0 || z+t.offset > t.d-1 {
		return false
	}

	idx := t.h*t.w*(z+t.offset) + t.w*(y+t.offset) + x + t.offset
	if idx < 0 || idx > t.w*t.h*t.d-1 {
		return false
	}
	return t.fields[idx]
}

func (t *Grid) set(x, y, z int, val bool) {
	if x+t.offset < 0 || x+t.offset > t.w-1 || y+t.offset < 0 || y+t.offset > t.h-1 || z+t.offset < 0 || z+t.offset > t.d-1 {
		println("Something wrong")
	}
	idx := t.h*t.w*(z+t.offset) + t.w*(y+t.offset) + x + t.offset
	t.fields[idx] = val
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
	grid := createGrid(len(lines[0]), len(lines), 1, 0)
	for y, line := range lines {
		for x, field := range line {
			grid.set(x, y, 0, field == Active)
		}
	}
	return grid
}

func printGrid(grid *Grid) {
	for z := 0; z < grid.d; z++ {
		println("z =", z)
		for y := 0; y < grid.h; y++ {
			for x := 0; x < grid.w; x++ {
				if grid.at(x-grid.offset, y-grid.offset, z-grid.offset) {
					print(string(Active))
				} else {
					print(string(Inactive))
				}
			}
			println()
		}
		println()
	}
}

func getNewCubeState(grid *Grid, x, y, z int) bool {
	activeNeighbors := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				if x != i || y != j || z != k {
					if grid.at(i, j, k) {
						activeNeighbors++
					}
				}
			}
		}
	}
	currentlyActive := grid.at(x, y, z)
	if currentlyActive && (activeNeighbors == 2 || activeNeighbors == 3) {
		return true
	}
	if !currentlyActive && activeNeighbors == 3 {
		return true
	}
	return false
}

func expandGrid(grid *Grid, offset int) (*Grid, int) {
	newGrid := createGrid(grid.w+2, grid.h+2, grid.d+2, offset)

	activeSum := 0
	for z := 0; z < newGrid.d; z++ {
		for y := 0; y < newGrid.h; y++ {
			for x := 0; x < newGrid.w; x++ {
				active := getNewCubeState(grid, x-offset, y-offset, z-offset)
				newGrid.set(x-newGrid.offset, y-newGrid.offset, z-newGrid.offset, active)
				if active {
					activeSum++
				}
			}
		}
	}

	return newGrid, activeSum
}

func part1(lines []string) int {
	grid := parseInput(lines)
	printGrid(grid)
	activeSum := 0
	for i := 0; i < 6; i++ {
		grid, activeSum = expandGrid(grid, (i + 1))
		printGrid(grid)
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
