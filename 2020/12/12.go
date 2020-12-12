package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
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

// Direction -
type Direction byte

// -
const (
	East  Direction = 'E'
	West  Direction = 'W'
	North Direction = 'N'
	South Direction = 'S'
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func rotate(currentDirection Direction, degrees int) Direction {
	clockwiseDirections := []Direction{East, South, West, North}
	step := degrees / 90

	index := 0
	for i, direction := range clockwiseDirections {
		if direction == currentDirection {
			index = i
			break
		}
	}

	newIndex := index + step
	if newIndex < 0 {
		newIndex = len(clockwiseDirections) + newIndex
	}

	return clockwiseDirections[newIndex%len(clockwiseDirections)]

}

func rotateWaypoint(x int, y int, degrees int) (int, int) {
	radians := float64(degrees) * math.Pi / 180

	newX := float64(x)*math.Cos(radians) + float64(y)*math.Sin(radians)
	newY := -float64(x)*math.Sin(radians) + float64(y)*math.Cos(radians)

	return int(math.Round(newX)), int(math.Round(newY))
}

func part1(lines []string) int {
	direction := East
	northPosition := 0
	eastPosition := 0

	for _, line := range lines {
		command := line[0]
		value, _ := strconv.Atoi(line[1:])

		switch command {
		case 'N':
			northPosition += value
		case 'S':
			northPosition -= value
		case 'E':
			eastPosition += value
		case 'W':
			eastPosition -= value
		case 'L':
			direction = rotate(direction, -value)
		case 'R':
			direction = rotate(direction, value)
		case 'F':
			switch direction {
			case East:
				eastPosition += value
			case West:
				eastPosition -= value
			case North:
				northPosition += value
			case South:
				northPosition -= value
			}
		}

	}

	return abs(northPosition) + abs(eastPosition)
}

func part2(lines []string) int {
	shipNorth := 0
	shipEast := 0

	waypointNorth := 1
	waypointEast := 10

	for _, line := range lines {
		command := line[0]
		value, _ := strconv.Atoi(line[1:])

		switch command {
		case 'N':
			waypointNorth += value
		case 'S':
			waypointNorth -= value
		case 'E':
			waypointEast += value
		case 'W':
			waypointEast -= value
		case 'L':
			waypointEast, waypointNorth = rotateWaypoint(waypointEast, waypointNorth, -value)
		case 'R':
			waypointEast, waypointNorth = rotateWaypoint(waypointEast, waypointNorth, value)
		case 'F':
			shipEast += waypointEast * value
			shipNorth += waypointNorth * value
		}
	}

	return abs(shipNorth) + abs(shipEast)
}

func main() {
	lines := readLines("12.txt")
	println(part1(lines))
	println(part2(lines))
}
