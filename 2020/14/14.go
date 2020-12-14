package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Instruction -
type Instruction interface {
	isInstruction()
}

// MaskInstruction -
type MaskInstruction struct {
	mask string
}

// WriteInstruction -
type WriteInstruction struct {
	address uint64
	value   uint64
}

func (i MaskInstruction) isInstruction()  {}
func (i WriteInstruction) isInstruction() {}

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
	arrayAccessRegexp := regexp.MustCompile(`\[(.*)\]`)

	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		if strings.HasPrefix(line, "mask") {
			split := strings.Split(line, "=")
			mask := strings.ReplaceAll(split[1], " ", "")
			instructions[i] = MaskInstruction{mask: mask}
		} else if strings.HasPrefix(line, "mem") {
			split := strings.Split(line, "=")
			indexMatch := arrayAccessRegexp.FindStringSubmatch(split[0])
			index, _ := strconv.ParseInt(indexMatch[1], 10, 64)
			value, _ := strconv.ParseInt(strings.ReplaceAll(split[1], " ", ""), 10, 64)
			instructions[i] = WriteInstruction{address: uint64(index), value: uint64(value)}
		}
	}
	return instructions
}

func setBit(value uint64, bit int) uint64 {
	return value | uint64(1)<<(35-bit)
}

func clearBit(value uint64, bit int) uint64 {
	return value & uint64(^(uint64(1) << (35 - bit)))
}

func applyMask(value uint64, mask string) uint64 {
	if mask == "" {
		return value
	}

	for i, maskBit := range mask {
		if maskBit == '0' {
			value = clearBit(value, i)
		} else if maskBit == '1' {
			value = setBit(value, i)
		}
	}

	return value
}

func applyFloatingMask(value uint64, mask string) []uint64 {
	if mask == "" {
		return []uint64{value}
	}

	floatingBits := make([]int, 0)

	for i, maskBit := range mask {
		if maskBit == '1' {
			value = setBit(value, i)
		} else if maskBit == 'X' {
			floatingBits = append(floatingBits, i)
		}
	}

	values := []uint64{value}

	for _, floatingBit := range floatingBits {
		newValues := make([]uint64, 0)
		for _, val := range values {
			newValues = append(newValues, setBit(val, floatingBit))
			newValues = append(newValues, clearBit(val, floatingBit))
		}
		values = newValues
	}

	return values
}

func getSum(memory map[uint64]uint64) uint64 {
	var sum uint64 = 0
	for _, val := range memory {
		sum += val
	}
	return sum
}

func part1(program []Instruction) uint64 {
	memory := map[uint64]uint64{}
	currentMask := ""

	for _, line := range program {
		switch instruction := line.(type) {
		case WriteInstruction:
			memory[instruction.address] = applyMask(instruction.value, currentMask)
		case MaskInstruction:
			currentMask = instruction.mask
		}
	}

	return getSum(memory)
}

func part2(program []Instruction) uint64 {
	memory := map[uint64]uint64{}
	currentMask := ""

	for _, line := range program {
		switch instruction := line.(type) {
		case WriteInstruction:
			addresses := applyFloatingMask(instruction.address, currentMask)
			for _, address := range addresses {
				memory[address] = instruction.value
			}
		case MaskInstruction:
			currentMask = instruction.mask
		}
	}

	return getSum(memory)
}

func main() {
	lines := readLines("14.txt")
	program := parseProgram(lines)
	println(part1(program))
	println(part2(program))
}
