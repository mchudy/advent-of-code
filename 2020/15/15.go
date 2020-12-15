package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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

func part1(numbers []int) int {
	lastSpoken := map[int]int{}

	for i := 0; i < len(numbers)-1; i++ {
		lastSpoken[numbers[i]] = i
	}

	previousNumber := numbers[len(numbers)-1]
	currentNumber := numbers[len(numbers)-1]

	for i := len(numbers); i < 30000000; i++ {
		previousSpoken, ok := lastSpoken[currentNumber]
		if !ok {
			previousNumber = currentNumber

			currentNumber = 0
			lastSpoken[previousNumber] = i - 1
		} else {
			previousNumber = currentNumber

			currentNumber = i - previousSpoken - 1
			lastSpoken[previousNumber] = i - 1
		}
	}

	return currentNumber
}

func part2(numbers []int) int {
	return 0
}

func main() {
	lines := readLines("15.txt")
	numberText := strings.Split(lines[0], ",")
	numbers := make([]int, len(numberText))
	for i, n := range numberText {
		numbers[i], _ = strconv.Atoi(n)
	}
	println(part1(numbers))
	println(part2(numbers))
}
