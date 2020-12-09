package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func readLines(path string) []int64 {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		number, _ := strconv.ParseInt(text, 10, 64)
		lines = append(lines, number)
	}

	return lines
}

func findError(numbers []int64, preambleSize int) int64 {
	for i := preambleSize; i < len(numbers); i++ {
		valid := false
		for j := i - preambleSize; j < i; j++ {
			for k := j; k < i; k++ {
				if numbers[j]+numbers[k] == numbers[i] && numbers[j] != numbers[k] {
					valid = true
					continue
				}
			}
			if valid {
				continue
			}
		}
		if !valid {
			return numbers[i]
		}
	}
	return 0
}

func findContiguousSum(numbers []int64, n int64) (int, int) {
	start := 0
	sum := int64(0)

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == n {
			start = i + 1
			continue
		}

		sum += numbers[i]
		for sum > n && start < i {
			sum -= numbers[start]
			start++
		}

		if sum == n {
			return start, i
		}
	}
	return 0, 0
}

func main() {
	numbers := readLines("09.txt")
	errorNumber := findError(numbers, 25)
	println(errorNumber)
	start, end := findContiguousSum(numbers, errorNumber)

	min := int64(math.MaxInt64)
	max := int64(0)
	for i := start; i <= end; i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	println(start, end)
	println(min + max)
}
