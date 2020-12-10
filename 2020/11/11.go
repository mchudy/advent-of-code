package main

import (
	"bufio"
	"log"
	"os"
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

func main() {
	numbers := readLines("11.txt")
	println(numbers)
}
