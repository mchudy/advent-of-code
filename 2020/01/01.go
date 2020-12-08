package main

import (
	"bufio"
	"log"
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

func main() {
	lines := readLines("01.txt")

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			a, _ := strconv.Atoi(lines[i])
			b, _ := strconv.Atoi(lines[j])
			if a+b == 2020 {
				println("2 numbers")
				println(a * b)
			}
			for k := j + 1; k < len(lines); k++ {
				c, _ := strconv.Atoi(lines[k])
				if a+b+c == 2020 {
					println("3 numbers")
					println(a * b * c)
				}
			}
		}
	}
}
