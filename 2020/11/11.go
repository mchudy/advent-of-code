package main

import (
	"bufio"
	"log"
	"os"
)

// -
const (
	EmptySeat    = 'L'
	OccupiedSeat = '#'
	Floor        = '.'
)

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

func copyLines(lines []string) []string {
	newLayout := make([]string, len(lines))
	for i := 0; i < len(lines); i++ {
		newLayout[i] = lines[i]
	}
	return newLayout
}

func countOccupied(lines []string) int {
	count := 0
	for _, line := range lines {
		for _, seat := range line {
			if seat == OccupiedSeat {
				count++
			}
		}
	}
	return count
}

func safeGetSeat(lines []string, x int, y int) (byte, bool) {
	if (x < 0 || x > len(lines)-1) || y < 0 || y > len(lines[0])-1 {
		return ' ', false
	}
	return lines[x][y], true
}

func checkSeatType(lines []string, x int, y int, seatType byte) bool {
	seat, ok := safeGetSeat(lines, x, y)
	if !ok {
		return false
	}
	return seat == seatType
}

func countAdjacentOcccupied(lines []string, x int, y int, directions [][]int) int {
	count := 0
	for _, direction := range directions {
		if checkSeatType(lines, x+direction[0], y+direction[1], OccupiedSeat) {
			count++
		}
	}
	return count
}

func countOccupiedInDirection(lines []string, x int, y int, directions [][]int) int {
	count := 0
	for _, direction := range directions {
		currentX := x
		currentY := y

		for true {
			currentX += direction[0]
			currentY += direction[1]

			seat, ok := safeGetSeat(lines, currentX, currentY)
			if ok && seat == OccupiedSeat {
				count++
				break
			}
			if seat == EmptySeat || !ok {
				break
			}
		}
	}
	return count
}

func part1(lines []string, directions [][]int) int {
	for i := 0; ; i++ {
		changed := false
		newLayout := copyLines(lines)

		for j := 0; j < len(lines); j++ {
			for k := 0; k < len(lines[j]); k++ {
				seat := lines[j][k]

				if seat == EmptySeat && countAdjacentOcccupied(lines, j, k, directions) == 0 {
					newLayout[j] = newLayout[j][:k] + string(OccupiedSeat) + newLayout[j][k+1:]
					changed = true
				} else if seat == OccupiedSeat && countAdjacentOcccupied(lines, j, k, directions) >= 4 {
					newLayout[j] = newLayout[j][:k] + string(EmptySeat) + newLayout[j][k+1:]
					changed = true
				} else {
					newLayout[j] = newLayout[j][:k] + string(lines[j][k]) + newLayout[j][k+1:]
				}
			}
		}

		lines = newLayout
		if !changed {
			return countOccupied(lines)
		}
	}
}

func part2(lines []string, directions [][]int) int {
	for i := 0; ; i++ {
		changed := false
		newLayout := copyLines(lines)

		for j := 0; j < len(lines); j++ {
			for k := 0; k < len(lines[j]); k++ {
				seat := lines[j][k]

				if seat == EmptySeat && countOccupiedInDirection(lines, j, k, directions) == 0 {
					newLayout[j] = newLayout[j][:k] + string(OccupiedSeat) + newLayout[j][k+1:]
					changed = true
				} else if seat == OccupiedSeat && countOccupiedInDirection(lines, j, k, directions) >= 5 {
					newLayout[j] = newLayout[j][:k] + string(EmptySeat) + newLayout[j][k+1:]
					changed = true
				} else {
					newLayout[j] = newLayout[j][:k] + string(lines[j][k]) + newLayout[j][k+1:]
				}
			}
		}

		lines = newLayout
		if !changed {
			return countOccupied(lines)
		}
	}
}

func main() {
	lines := readLines("11.txt")

	directions := [][]int{
		[]int{-1, -1},
		[]int{-1, 1},
		[]int{1, -1},
		[]int{-1, 0},
		[]int{1, 0},
		[]int{0, -1},
		[]int{0, 1},
		[]int{1, 1},
	}

	println(part1(lines, directions))
	println(part2(lines, directions))
}
