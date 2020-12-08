package main

import (
	"bufio"
	"log"
	"os"
)

// SquareType : type of square
type SquareType bool

// TreeMap : map
type TreeMap [][]SquareType

const (
	// Tree : tree
	Tree SquareType = true
	// OpenSquare : open square
	OpenSquare = false
)

const treeChar = '#'
const openSquareChar = '.'

// Vector : simple 2D vector type
type Vector struct {
	x int
	y int
}

func nextPosition(position Vector, slope Vector) Vector {
	return Vector{x: position.x + slope.x, y: position.y + slope.y}
}

func buildTreeMap(lines []string) TreeMap {
	height := len(lines)
	width := len(lines[0])

	treeMap := make(TreeMap, 0)

	for i := 0; i < height; i++ {
		line := lines[i]
		tmp := make([]SquareType, 0)
		for j := 0; j < width; j++ {
			var value SquareType
			if line[j] == treeChar {
				value = Tree
			} else {
				value = OpenSquare
			}

			tmp = append(tmp, value)
		}
		treeMap = append(treeMap, tmp)
	}

	return treeMap
}

func countTrees(treeMap TreeMap, slope Vector) int {
	encounteredTrees := 0
	height := len(treeMap)
	width := len(treeMap[0])
	position := Vector{x: 0, y: 0}

	for position.y < height {
		if treeMap[position.y%height][position.x%width] == Tree {
			encounteredTrees++
		}
		position = nextPosition(position, slope)
	}

	return encounteredTrees
}

func main() {
	lines := readLines("03.txt")
	treeMap := buildTreeMap(lines)

	slopes := []Vector{
		Vector{x: 1, y: 1},
		Vector{x: 3, y: 1},
		Vector{x: 5, y: 1},
		Vector{x: 7, y: 1},
		Vector{x: 1, y: 2},
	}

	result := 1
	for _, slope := range slopes {
		treesCount := countTrees(treeMap, slope)
		println(slope.x, slope.y, result)
		result *= treesCount
	}
	println("product", result)
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
