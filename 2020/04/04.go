package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// PassportDataKey -
type PassportDataKey string

// Bullshit comment to satisfy the annoying linter
const (
	BirthYear      PassportDataKey = "byr"
	IssueYear      PassportDataKey = "iyr"
	ExpirationYear PassportDataKey = "eyr"
	Height         PassportDataKey = "hgt"
	HairColor      PassportDataKey = "hcl"
	EyeColor       PassportDataKey = "ecl"
	PassportID     PassportDataKey = "pid"
	CountryID      PassportDataKey = "cid"
)

// Passport - shut up
type Passport struct {
	birthYear      string
	issueYear      string
	expirationYear string
	height         string
	hairColor      string
	eyeColor       string
	passportID     string
	countryID      string
}

func readLineBlocks(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lineBlocks := make([][]string, 0)
	currentLineBlock := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			lineBlocks = append(lineBlocks, currentLineBlock)
			currentLineBlock = make([]string, 0)
		} else {
			currentLineBlock = append(currentLineBlock, line)
		}
	}
	lineBlocks = append(lineBlocks, currentLineBlock)

	return lineBlocks
}

func parsePassport(lines []string) Passport {
	data := make(map[PassportDataKey]string)
	for _, line := range lines {
		pairs := strings.Split(line, " ")
		for _, pair := range pairs {
			pairData := strings.Split(pair, ":")
			data[PassportDataKey(pairData[0])] = pairData[1]
		}
	}

	return Passport{
		birthYear:      data[BirthYear],
		issueYear:      data[IssueYear],
		countryID:      data[CountryID],
		expirationYear: data[ExpirationYear],
		eyeColor:       data[EyeColor],
		hairColor:      data[HairColor],
		height:         data[Height],
		passportID:     data[PassportID],
	}
}

func contains(array []string, element string) bool {
	for _, arrayElement := range array {
		if arrayElement == element {
			return true
		}
	}
	return false
}

func validateNumberRange(value string, min int, max int) bool {
	number, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return number >= min && number <= max
}

// Unit -
type Unit struct {
	suffix string
	min    int
	max    int
}

func validateUnit(value string, allowedUnits []Unit) bool {
	for _, unit := range allowedUnits {
		if strings.HasSuffix(value, unit.suffix) {
			return validateNumberRange(
				strings.TrimSuffix(value, unit.suffix),
				unit.min,
				unit.max)
		}
	}
	return false
}

func validatePassport(passport Passport) bool {
	validEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	passportIDValid, _ := regexp.MatchString(`^[0-9]{9}$`, passport.passportID)
	hairColorValid, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, passport.hairColor)

	return validateNumberRange(passport.birthYear, 1920, 2002) &&
		validateNumberRange(passport.issueYear, 2010, 2020) &&
		validateNumberRange(passport.expirationYear, 2020, 2030) &&
		contains(validEyeColors, passport.eyeColor) &&
		hairColorValid &&
		validateUnit(passport.height, []Unit{
			Unit{suffix: "cm", min: 150, max: 193},
			Unit{suffix: "in", min: 59, max: 76},
		}) &&
		passportIDValid
}

func main() {
	lineBlocks := readLineBlocks("04.txt")
	validPassports := 0

	for _, lineBlock := range lineBlocks {
		passport := parsePassport(lineBlock)
		if validatePassport(passport) {
			validPassports++
		}
	}

	println(validPassports)
}
