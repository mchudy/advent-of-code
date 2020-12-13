package main

import (
	"bufio"
	"log"
	"math"
	"math/big"
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

func lcm(a, b, gcd *big.Int) *big.Int {
	result := new(big.Int)
	result.Mul(a, b)
	result.Div(result, gcd)
	return result
}

func offsetLcm(a, b, offsetA, offsetB *big.Int) (*big.Int, *big.Int) {
	s := new(big.Int)
	t := new(big.Int)
	gcd := new(big.Int)
	period := new(big.Int)
	newOffset := new(big.Int)

	gcd.GCD(s, t, a, b)
	period.Set(lcm(a, b, gcd))

	// s*a + t*b = 1 Bezout identity, a and b coprime
	// m*a - offsetA = n*b - offsetB
	// m*a - n*b = offsetA - offsetB = z
	// s*a*z + t*b*z = z
	// m = z*s
	// newOffset = (offsetA - m*a) mod lcm(a,b)

	z := new(big.Int).Sub(offsetA, offsetB)
	newOffset.Mul(z, s)
	newOffset.Mul(newOffset, a)
	newOffset.Sub(offsetA, newOffset)
	newOffset.Mod(newOffset, period)

	return period, newOffset
}

func part1(lines []string) int {
	startTime, _ := strconv.Atoi(lines[0])
	busesText := strings.Split(lines[1], ",")
	buses := make([]int, 0)

	earliestBus := math.MaxInt32
	earliestWait := math.MaxInt32

	for _, bus := range busesText {
		if bus != "x" {
			busID, _ := strconv.Atoi(bus)
			buses = append(buses, busID)

			factor := int(math.Ceil(float64(startTime) / float64(busID)))
			wait := factor*busID - startTime
			if wait < earliestWait {
				earliestWait = wait
				earliestBus = busID
			}
		}
	}
	return earliestBus * earliestWait
}

func part2(lines []string) *big.Int {
	busesText := strings.Split(lines[1], ",")

	period, _ := new(big.Int).SetString(busesText[0], 10)
	offset := big.NewInt(int64(0))

	for i := 1; i < len(busesText); i++ {
		if busesText[i] != "x" {
			current, _ := new(big.Int).SetString(busesText[i], 10)
			newPeriod, newOffset := offsetLcm(period, current, offset, big.NewInt(int64(-i)))
			period.Set(newPeriod)
			offset.Set(newOffset)
		}
	}

	return offset
}

func main() {
	lines := readLines("13.txt")
	println(part1(lines))
	println(part2(lines).String())
}
