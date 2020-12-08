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

func parseLine(line string) (int, int, string, string) {
	data := strings.Split(line, " ")

	frequency := strings.Split(data[0], "-")
	startPosition, _ := strconv.Atoi(frequency[0])
	endPosition, _ := strconv.Atoi(frequency[1])
	letter := strings.TrimSuffix(data[1], ":")
	password := data[2]

	return startPosition, endPosition, letter, password
}

func validatePassword1(min int, max int, letter string, password string) bool {
	letterOccurences := 0
	for _, currentLetter := range password {
		if letter == string(currentLetter) {
			letterOccurences++
		}
	}
	return letterOccurences >= min && letterOccurences <= max
}

func validatePassword2(position1 int, position2 int, letter byte, password string) bool {
	return (password[position1] == letter || password[position2] == letter) &&
		password[position1] != password[position2]
}

func main() {
	lines := readLines("02.txt")
	validPasswords1 := 0
	validPasswords2 := 0

	for _, line := range lines {
		startPosition, endPosition, letter, password := parseLine(line)

		if validatePassword1(startPosition, endPosition, letter, password) {
			validPasswords1++
		}

		if validatePassword2(startPosition-1, endPosition-1, letter[0], password) {
			validPasswords2++
		}
	}
	println(validPasswords1)
	println(validPasswords2)
}
