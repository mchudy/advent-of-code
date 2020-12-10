package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func readLines(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, number)
	}

	return lines
}

const maxDiff = 3

func getJoltDifferences(numbers []int) map[int]int {
	currentJoltage := 0

	differences := map[int]int{}
	for i := 1; i < maxDiff; i++ {
		differences[i] = 0
	}
	differences[maxDiff] = 1

	for i := 0; i < len(numbers); i++ {
		diff := numbers[i] - currentJoltage
		differences[diff]++
		currentJoltage = numbers[i]
	}

	return differences
}

func countAllArrangements(numbers []int, index int, lookup map[int]int) int {
	if index == len(numbers)-1 {
		return 1
	}

	count := 0
	for i := index + 1; i < len(numbers) && numbers[i]-numbers[index] <= maxDiff; i++ {
		cached, ok := lookup[i]
		if ok {
			count += cached
		} else {
			newCount := countAllArrangements(numbers, i, lookup)
			lookup[i] = newCount
			count += newCount
		}
	}

	return count
}

func main() {
	numbers := readLines("10.txt")
	sort.Ints(numbers)
	differences := getJoltDifferences(numbers)
	println(differences[1] * differences[3])
	numbers = append([]int{0}, numbers...)
	println(countAllArrangements(numbers, 0, map[int]int{}))
}
