package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Instruction -
type Instruction struct {
	command string
	value   int
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

func parseProgram(lines []string) []Instruction {
	program := make([]Instruction, len(lines))
	for i := 0; i < len(lines); i++ {
		split := strings.Split(lines[i], " ")
		value, _ := strconv.Atoi(split[1])
		program[i] = Instruction{command: split[0], value: value}
	}
	return program
}

func executeInstruction(instruction Instruction, line int, acc int) (int, int) {
	switch instruction.command {
	case "jmp":
		return line + instruction.value, acc
	case "acc":
		return line + 1, acc + instruction.value
	case "nop":
		return line + 1, acc
	}
	return line, acc
}

func executeProgram(instructions []Instruction) (int, bool) {
	acc := 0
	line := 0
	executedLines := make([]bool, len(instructions))

	for line >= 0 && line < len(instructions) && executedLines[line] == false {
		executedLines[line] = true
		line, acc = executeInstruction(instructions[line], line, acc)
	}

	return acc, line == len(instructions)
}

func fixProgram(instructions []Instruction) int {
	for i := 0; i < len(instructions); i++ {
		originalCommand := instructions[i].command
		if originalCommand == "jmp" {
			instructions[i].command = "nop"
		} else if originalCommand == "nop" {
			instructions[i].command = "jmp"
		} else {
			continue
		}

		acc, success := executeProgram(instructions)
		if success {
			return acc
		}

		instructions[i].command = originalCommand

	}
	return 0
}

func main() {
	lines := readLines("08.txt")
	program := parseProgram(lines)
	println(executeProgram(program))
	println(fixProgram(program))
}
