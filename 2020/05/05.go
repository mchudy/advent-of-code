package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
)

const rowsCount = 128
const columnsCount = 8

func binarySearchStep(min int, max int) int {
	return int(math.Ceil(float64((max - min)) / 2))
}

func getSeatPosition(code string) (int, int) {
	min := 0
	max := 127
	for i := 0; i < 8; i++ {
		step := binarySearchStep(min, max)
		if code[i] == 'F' {
			max = max - step
		} else if code[i] == 'B' {
			min = min + step
		}
	}
	row := min

	min = 0
	max = 7
	for i := 0; i < 3; i++ {
		step := binarySearchStep(min, max)
		if code[i+7] == 'L' {
			max = max - step
		} else if code[i+7] == 'R' {
			min = min + step
		}
	}
	column := min

	return row, column
}

func getSeatID(row int, column int) int {
	return row*8 + column
}

func findEmptySeat(allIDs []int) int {
	sort.Ints(allIDs)
	for i := 0; i < len(allIDs)-1; i++ {
		if allIDs[i+1]-allIDs[i] > 1 {
			return allIDs[i+1] - 1
		}
	}
	return -1
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

func main() {
	lines := readLines("05.txt")
	maxSeatID := 0
	var allIDs []int
	for _, line := range lines {
		row, column := getSeatPosition(line)
		seatID := getSeatID(row, column)
		allIDs = append(allIDs, seatID)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	println("max seat ID", maxSeatID)

	emptySeat := findEmptySeat(allIDs)
	println("my seat", emptySeat)
}
